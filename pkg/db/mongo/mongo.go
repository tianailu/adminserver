package mongo

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"net/url"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"time"

	"encoding/json"
	"github.com/casbin/casbin"
	"github.com/pkg/errors"
	"io/ioutil"
	"sync"
)

const (
	rbacModelFile   = "/var/config-map/rbac_model.conf"
	permissionsFile = "/var/config-map/permissions.conf"
)

type SessionCaller struct {
	Name      string `json:"name,omitempty"`
	File      string `json:"file"`
	Line      int    `json:"line"`
	Time      string `json:"time"`
	SessionId string `json:"session_id"`
}

var (
	globalSession      mgo.Session
	globalPoolLimit    int
	GlobalSessionCount = new(sync.Map)
	DBAddr             string
	Enforcer           *casbin.Enforcer
	dbLocker           = sync.Mutex{}
)

const (
	ModeSlave   = "Slave"   //主从复制集 Slave
	ModeMaster  = "Master"  //主从复制集 Master
	ModeArbiter = "Arbiter" //主从复制集 Arbiter 仲裁者
	ModeNormal  = "Normal"  //Normal
	KeyFilePath = "/var/CA/client/mongodb-keyfile"
	ReplicaPath = "/var/CA/client/mongodb-replSet"
)

var (
	DBKeyFile string //mongodb keyfile
	DBMode    = ModeNormal
)

func InitDB(username, password, ip, port, maxPoolLimitString, socketTime string) {
	addr := fmt.Sprintf("mongodb://%s:%s@%s:%s/admin", url.QueryEscape(username), url.QueryEscape(password), ip, port)
	maxPoolLimit, _ := strconv.Atoi(maxPoolLimitString)
	gg, _ := strconv.Atoi(socketTime)
	st := time.Duration(gg)
	timeout := st * time.Second
	defer CloseDB()

	//初始化模块权限
	//err := InitPermissions()
	//if err != nil {
	//	log.Println(err.Error())
	//	return
	//}
	s, err := mgo.Dial(addr)
	if err != nil {
		panic(err)
	}
	s.SetMode(mgo.Primary, true)
	s.SetPoolLimit(maxPoolLimit)
	s.SetSocketTimeout(timeout * time.Second)

	// 确保表索引存在
	_ = s.DB("user_manage").C("user").EnsureIndex(DbIndexKey("username"))

}

func CloseDB() {
	defer globalSession.Close()
}

func GetDB() *DB {
	dbLocker.Lock()
	defer dbLocker.Unlock()
	id := bson.NewObjectId().Hex()
	if pc, _, _, ok := runtime.Caller(1); ok {
		if f := runtime.FuncForPC(pc); f != nil {
			var count int
			var caller SessionCaller
			caller.Name = f.Name()
			caller.SessionId = id
			caller.Time = time.Now().Format("2006-01-02 15:04:05.999999999")
			caller.File, caller.Line = f.FileLine(pc)
			GlobalSessionCount.Store(id, caller)
			GlobalSessionCount.Range(func(key, value interface{}) bool {
				count++
				return true
			})
			if count >= globalPoolLimit {
				var dbDebugLog = make(map[string][]SessionCaller)
				GlobalSessionCount.Range(func(key, value interface{}) bool {
					if callInfo, ok := value.(SessionCaller); ok {
						pushInfo := callInfo
						pushInfo.Name = ""
						if oneLog, exist := dbDebugLog[callInfo.Name]; exist {
							dbDebugLog[callInfo.Name] = append(oneLog, pushInfo)
						} else {
							dbDebugLog[callInfo.Name] = []SessionCaller{pushInfo}
						}
					}
					return true
				})

				if data, err := json.MarshalIndent(dbDebugLog, "", "\t"); err == nil {
					_ = ioutil.WriteFile(filepath.Join("/var/log", "edr-mongo.log"), data, os.ModePerm)
				}
			}
		}

	}
	return &DB{session: globalSession.Copy(), sessionId: id}
}

type DB struct {
	sessionId string
	session   *mgo.Session
	dname     string // db name
	cname     string // collection name
	D         *mgo.Database
	C         *mgo.Collection
}

func (db *DB) Ping() error {
	return db.session.Ping()
}

func (db *DB) SwitchTo(dname string, cname string) {
	if dname != "" && db.dname != dname {
		db.D = db.session.DB(dname)
		db.dname = dname
	}

	if cname != "" {
		db.C = db.D.C(cname)
		db.cname = cname
	}
}

func (db *DB) Close() {
	db.session.Close()
	GlobalSessionCount.Delete(db.sessionId)
}

func DbIndexKey(key ...string) mgo.Index {
	return mgo.Index{
		Key:        key,
		Unique:     true,
		DropDups:   true,
		Background: false, // See notes.
		Sparse:     true,
	}
}

func DbIndexKeyComm(key ...string) mgo.Index {
	return mgo.Index{
		Key:        key,
		Unique:     false,
		DropDups:   true,
		Background: false, // See notes.
		Sparse:     true,
	}
}

func GetDataBaseConfigMode() string {
	type MongodbReplSet struct {
		EnReplicaSet bool     `json:"enReplicaSet"` //是否启用复制集
		ReplSetName  string   `json:"replSetName"`  //复制集名称
		ReplSetMe    string   `json:"rsMe"`         //自身 ip:port  192.168.155.5:27017
		ReplSetHost  []string `json:"rsHost"`       //集群主机["ip:port", "ip:port", "ip:port"] 包括自身
		Arbiter      []string `json:"arbiter"`      //仲裁者
		Primary      bool     `json:"primary"`      //主复制集 (同一个集合能只能设置一个主复制集)
	}

	ReplicaSet := MongodbReplSet{}

	configData, errConf := ioutil.ReadFile(ReplicaPath)
	if errConf != nil {
		return ModeNormal
	}

	errSet := json.Unmarshal(configData, &ReplicaSet)
	if errSet != nil {
		return ModeNormal
	}

	if ReplicaSet.EnReplicaSet == false {
		return ModeNormal
	}

	if ReplicaSet.ReplSetMe == "" || ReplicaSet.ReplSetName == "" || len(ReplicaSet.ReplSetHost) == 0 {
		return ModeNormal
	}

	if ReplicaSet.Primary {
		return ModeMaster
	}

	return ModeSlave
}

func GetDataBaseMode() string {

	type IsMasterResults struct {
		// The following fields hold information about the specific mongodb node.
		IsMaster  bool      `bson:"ismaster"`
		Secondary bool      `bson:"secondary"`
		Arbiter   bool      `bson:"arbiterOnly"`
		Address   string    `bson:"me"`
		LocalTime time.Time `bson:"localTime"`

		// The following fields hold information about the replica set.
		ReplicaSetName string   `bson:"setName"`
		Addresses      []string `bson:"hosts"`
		Arbiters       []string `bson:"arbiters"`
		PrimaryAddress string   `bson:"primary"`
	}

	type MongodbReplSet struct {
		EnReplicaSet bool     `json:"enReplicaSet"` //是否启用复制集
		ReplSetName  string   `json:"replSetName"`  //复制集名称
		ReplSetMe    string   `json:"rsMe"`         //自身 ip:port  192.168.155.5:27017
		ReplSetHost  []string `json:"rsHost"`       //集群主机["ip:port", "ip:port", "ip:port"] 包括自身
		Arbiter      []string `json:"arbiter"`      //仲裁者
		Primary      bool     `json:"primary"`      //主复制集 (同一个集合能只能设置一个主复制集)
	}

	session, err := mgo.Dial(DBAddr)
	if err != nil {
		log.Println(DBAddr)
		log.Println("connect db failed.. sleep 3s and exiting...", err.Error())
		return ModeNormal
	}
	defer session.Close()

	ReplicaSet := MongodbReplSet{}

	configData, errConf := ioutil.ReadFile(ReplicaPath)
	if errConf != nil {
		return ModeNormal
	}

	errSet := json.Unmarshal(configData, &ReplicaSet)
	if errSet != nil {
		return ModeNormal
	}

	if ReplicaSet.EnReplicaSet == false {
		return ModeNormal
	}

	if ReplicaSet.ReplSetMe == "" || ReplicaSet.ReplSetName == "" || len(ReplicaSet.ReplSetHost) == 0 {
		return ModeNormal
	}

	var result IsMasterResults
	if err := session.Run("isMaster", &result); err != nil {
		panic(err)
	}

	if result.PrimaryAddress == "" || result.ReplicaSetName == "" {
		return ModeNormal
	}

	if len(result.Addresses) == 0 {
		return ModeNormal
	}

	if ReplicaSet.ReplSetName != result.ReplicaSetName {
		return ModeNormal
	}

	if ReplicaSet.ReplSetMe == result.PrimaryAddress {
		return ModeMaster
	}

	for _, val := range result.Arbiters {
		if val == ReplicaSet.ReplSetMe {
			return ModeArbiter
		}
	}

	for _, val := range result.Addresses {
		if val == ReplicaSet.ReplSetMe {
			return ModeSlave
		}
	}

	return ModeNormal
}

func checkDataBaseMode() {
	var exitErr = errors.New("db mode change.")
	var maxTry = 3
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
			log.Println(exitErr.Error())
			os.Exit(0)
		}
	}()

	for {
		time.Sleep(time.Second * time.Duration(maxTry))
		if DBMode != GetDataBaseMode() {
			time.Sleep(time.Second * time.Duration(maxTry))
			if DBMode != GetDataBaseMode() {
				panic(exitErr)
			}
		}
	}
}
