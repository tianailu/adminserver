package mysql

import (
	"fmt"
	"github.com/labstack/gommon/log"
	"github.com/tianailu/adminserver/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
	"time"
)

//Db数据库连接池
var (
	conn *gorm.DB
	err  error
)

func GetDB() *gorm.DB {
	return conn
}

// InitMySQLDB 初始化 mysql 数据库
func InitMySQLDB(conf config.Mysql) {
	// 构建连接："用户名:密码@tcp(IP:端口)/数据库?charset=utf8"
	dsn := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		conf.Username, conf.Password, conf.Ip, conf.Port, conf.DBName)

	conn, err = gorm.Open(mysql.New(mysql.Config{
		DSN:               dsn,
		DefaultStringSize: 256,
	}), &gorm.Config{
		Logger: logger.New(
			log.New("gorm"), // 使用 Echo 的 Logger 实例
			logger.Config{
				SlowThreshold:             0,
				Colorful:                  false,
				IgnoreRecordNotFoundError: true,
				LogLevel:                  logger.Info,
			},
		),
	})

	//打开数据库,前者是驱动名，所以要导入： _ "github.com/go-sql-driver/mysql"
	if err != nil {
		log.Errorf("日志文件打开失败：%v \n", err)
		os.Exit(1)
	}

	sqlDb, _ := conn.DB()
	//可重用连接的最大时间
	sqlDb.SetMaxIdleConns(conf.ConnMaxIdle)
	//设置数据库最大连接数
	sqlDb.SetMaxOpenConns(conf.ConnMaxIdle)
	//设置上数据库最大闲置连接数
	sqlDb.SetConnMaxLifetime(time.Second * time.Duration(conf.ConnMaxLifetime))
	//验证连接
	if err := sqlDb.Ping(); err != nil {
		log.Errorf("MySQL Ping err: %s", err.Error())
		return
	}
	log.Infof("MySQL successfully connected")
}
