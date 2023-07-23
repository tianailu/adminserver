package config

import (
	"flag"
	"fmt"
	"github.com/beego/beego/v2/adapter/config"
	"github.com/mitchellh/mapstructure"
	"os"
	"strings"
)

const (
	defaultFilePath   = ""
	sep               = "_"
	DefaultConfigPath = "./config.conf"
	DefaultConfigType = "ini"
)

type Args struct {
	ConfigPath          string
	SyncLicenseInterval int
}

func (a *Args) Parse() {
	c := flag.String("c", DefaultConfigPath, "config file path")

	sLic := flag.Int("sync_license_interval", 10, "sync newest license interval")

	flag.Parse()
	a.ConfigPath = *c
	a.SyncLicenseInterval = *sLic
}

var (
	AuthConf    = make(map[string]string)
	LoggerConf  = Logger{}
	AuthConf2   = Auth{}
	MysqlConf   = Mysql{}
	MongoDBConf = Mongodb{}
	RedisConf   = Redis{}
)

type (
	SettingsConfig struct {
		FilePath string
		FileType string
		ConfigEr config.Configer
	}

	Logger struct {
		Redirect     bool   `mapstructure:"redirect"`
		FilePath     string `mapstructure:"filepath"`
		MaxSaveDay   int64  `mapstructure:"max_save_day"`
		MaxCleanSize string `mapstructure:"max_clean_size"`
		MaxCheckTime string `mapstructure:"max_check_time"`
	}

	Auth struct {
		AdminSecretKey    string `mapstructure:"admin_secret_key"`
		AdminPasswordSalt string `mapstructure:"admin_password_salt"`
		AppId             string `mapstructure:"app_id"`
		AppSecret         string `mapstructure:"app_secret"`
		MsgAppId          string `mapstructure:"msg_app_id"`
		MsgAppSecret      string `mapstructure:"msg_app_secret"`
		SecretId          string `mapstructure:"secret_id"`
		SecretKey         string `mapstructure:"secret_key"`
		MsgAppSign        string `mapstructure:"msg_app_sign"`
		MsgAppTid         string `mapstructure:"msg_app_tid"`
	}

	Mysql struct {
		Ip              string `mapstructure:"ip"`
		Port            int64  `mapstructure:"port"`
		Username        string `mapstructure:"username"`
		Password        string `mapstructure:"password"`
		DBName          string `mapstructure:"db_name"`
		ConnMaxLifetime int64  `mapstructure:"conn_max_lifetime"`
		ConnMaxIdle     int64  `mapstructure:"conn_max_idle"`
	}

	Mongodb struct {
		Ip            string `mapstructure:"ip"`
		Port          int64  `mapstructure:"port"`
		Username      string `mapstructure:"username"`
		Password      string `mapstructure:"password"`
		MaxPoolLimit  int64  `mapstructure:"max_pool_limit"`
		SocketTimeout int64  `mapstructure:"socket_timeout"`
	}

	Redis struct {
		Ip       string `mapstructure:"ip"`
		Port     int64  `mapstructure:"port"`
		Username string `mapstructure:"username"`
		Password string `mapstructure:"password"`
		DB       int64  `mapstructure:"db"`
	}
)

func Parse(fileType string, filepath string) (*SettingsConfig, error) {
	if fileType == "" {
		fileType = DefaultConfigType
	}
	if filepath == "" {
		filepath = defaultFilePath
	}

	conf, err := config.NewConfig(fileType, filepath)
	if err != nil {
		return nil, err
	}

	newConfig := new(SettingsConfig)
	newConfig.ConfigEr = conf
	newConfig.FileType = fileType
	newConfig.FilePath = filepath
	// 解析auth key
	AuthConf = newConfig.GetConfig("auth")
	if err := mapstructure.WeakDecode(newConfig.GetConfig("stdout_logger"), &LoggerConf); err != nil {
		return nil, err
	}
	if err := mapstructure.WeakDecode(newConfig.GetConfig("auth"), &AuthConf2); err != nil {
		return nil, err
	}
	if err := mapstructure.WeakDecode(newConfig.GetConfig("mysql"), &MysqlConf); err != nil {
		return nil, err
	}
	if err := mapstructure.WeakDecode(newConfig.GetConfig("mongodb"), &MongoDBConf); err != nil {
		return nil, err
	}
	if err := mapstructure.WeakDecode(newConfig.GetConfig("redis"), &RedisConf); err != nil {
		return nil, err
	}

	return newConfig, nil
}

// 各插件默认配置(modify 20210707)
func getDefaultConf(section string) map[string]string {
	switch section {
	case "mongodb":
		return map[string]string{
			"ip":             "127.0.0.1",
			"port":           "27017",
			"max_pool_limit": "20",
			"socket_timeout": "5",
		}
	case "mysql":
		return map[string]string{
			"ip":       "127.0.0.1",
			"port":     "8086",
			"username": "root",
			"password": "root",
		}
	case "redis":
		return map[string]string{
			"ip":   "127.0.0.1",
			"port": "6379",
		}
	default:
		return map[string]string{}

	}
}

// 从环境变量中获取参数，忽略大小写
func getFromEnv(s string) (v string) {
	if v = os.Getenv(strings.ToUpper(s)); v != "" {
		return
	}
	if v = os.Getenv(s); v != "" {
		return
	}
	return
}

// 如跳过解析配置文件，该函数会出异常
func getConf(owner *SettingsConfig, section string, df map[string]string) map[string]string {
	c, _ := owner.ConfigEr.GetSection(section)
	// 检查环境变量、配置文件覆盖默认配置
	for k, _ := range df {
		s := fmt.Sprintf("%s%s%s", section, sep, k)
		if v := getFromEnv(s); v != "" {
			df[k] = v
			continue
		}
		if v := c[k]; v != "" {
			df[k] = v
		}
	}
	// 检查默认配置中没有的字段
	for k, v := range c {
		if _, ok := df[k]; ok {
			continue
		}
		s := fmt.Sprintf("%s%s%s", section, sep, k)
		if j := getFromEnv(s); j != "" {
			df[k] = j
			continue
		}
		df[k] = v
	}
	return df
}

func (s *SettingsConfig) GetConfig(section string) map[string]string {
	return getConf(s, section, getDefaultConf(section))
}
