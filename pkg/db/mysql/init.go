package mysql

import (
	"fmt"
	"github.com/labstack/gommon/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
	"strconv"
	"time"
)

//Db数据库连接池
var (
	conn *gorm.DB
	err  error
)

func GetDb() *gorm.DB {
	return conn
}

// InitMysqlDB InitDB 注意方法名大写，就是public
func InitMysqlDB(username, password, ip, port, dbName, connMaxLifetime, connMaxIdle string) {
	dsn := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", username, password, ip, port, dbName)

	//构建连接："用户名:密码@tcp(IP:端口)/数据库?charset=utf8"
	conn, err = gorm.Open(mysql.New(mysql.Config{
		DSN:               dsn,
		DefaultStringSize: 256,
	}))

	//打开数据库,前者是驱动名，所以要导入： _ "github.com/go-sql-driver/mysql"
	if err != nil {
		log.Errorf("日志文件打开失败：%v \n", err)
		os.Exit(1)
	}
	//字符串转int
	connMaxLifetimeInt, _ := strconv.Atoi(connMaxLifetime)
	connMaxIdleInt, _ := strconv.Atoi(connMaxIdle)
	sqlDb, _ := conn.DB()
	//可重用连接的最大时间
	sqlDb.SetMaxIdleConns(connMaxIdleInt)
	//设置数据库最大连接数
	sqlDb.SetMaxOpenConns(connMaxIdleInt)
	//设置上数据库最大闲置连接数
	sqlDb.SetConnMaxLifetime(time.Second * time.Duration(connMaxLifetimeInt))
	//验证连接
	if err := sqlDb.Ping(); err != nil {
		return
	}
	fmt.Println("connect success")
}
