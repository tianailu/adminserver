package user

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/tianailu/adminserver/pkg/common"
	"github.com/tianailu/adminserver/pkg/db/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
)

var (
	userChan = make(chan User, 1000)
	table    = "ti_home_user"
)

// AddUser 添加用户
func AddUser(c echo.Context) error {
	var (
		resp common.ResponseNoData
		user User
	)
	err := c.Bind(&user)
	if err != nil {
		resp.Status = 1
		resp.Msg = "login failed"
		return c.JSON(http.StatusOK, resp)
	}
	userChan <- user
	resp.Status = 0
	resp.Msg = "ok"
	return c.JSON(http.StatusOK, resp)
}

// 获取数据
func GetUser(c echo.Context) error {
	var (
		resp common.Response
		user User
	)
	uId := c.QueryParam("uid")

	err := mysql.GetDb().Table(table).Where("uid = ?", uId).First(&user).Error
	if err != nil {
		fmt.Println("查询出错了", err)
		resp.Status = 1
		resp.Msg = "查询出错了， err：" + err.Error()
		return c.JSON(http.StatusOK, resp)
	}
	resp.Status = 0
	resp.Msg = "ok"
	resp.Data = user
	return c.JSON(http.StatusOK, resp)
}

func dealUserData() {
	for {
		fmt.Println("write user data")
		user := <-userChan
		err := createDatabase()
		if err != nil {
			log.Println("创建数据库失败, err:", err)
		}
		err = insertUser(user)
		if err != nil {
			log.Println("写入数据失败, err:", err)
		}
	}
}

// create database
func createDatabase() error {
	var (
		user User
	)
	/*sql := `CREATE TABLE ` + table + ` IF NOT EXISTS (
	    id INT(10) NOT NULL AUTO_INCREMENT,
	    uid VARCHAR(64) NULL DEFAULT NULL,
	    name VARCHAR(64) NULL DEFAULT NULL,
	    gender VARCHAR(16) NULL DEFAULT NULL,
	    birthday VARCHAR(32) NULL DEFAULT NULL,
	    height VARCHAR(16) NULL DEFAULT NULL,
	    weight VARCHAR(16) NULL DEFAULT NULL,
	    hometown VARCHAR(64) NULL DEFAULT NULL,
	    education VARCHAR(16) NULL DEFAULT NULL,
	    school VARCHAR(32) NULL DEFAULT NULL,
	    work VARCHAR(32) NULL DEFAULT NULL,
	    co_type VARCHAR(32) NULL DEFAULT NULL,
	    income VARCHAR(16) NULL DEFAULT NULL,
	    house_car VARCHAR(16) NULL DEFAULT NULL,
	    m_phone VARCHAR(11) NULL DEFAULT NULL,
	    wei_xin VARCHAR(32) NULL DEFAULT NULL,
	    marriage VARCHAR(32) NULL DEFAULT NULL,
	    habit VARCHAR(32) NULL DEFAULT NULL,
	    family VARCHAR(64) NULL DEFAULT NULL,
	    interest VARCHAR(128) NULL DEFAULT NULL,
	    character VARCHAR(128) NULL DEFAULT NULL,
	    future_plan VARCHAR(128) NULL DEFAULT NULL,
	    values VARCHAR(128) NULL DEFAULT NULL,
	    love_view VARCHAR(128) NULL DEFAULT NULL,
	    best_wish VARCHAR(128) NULL DEFAULT NULL,
	    best_height VARCHAR(128) NULL DEFAULT NULL,
	    is_divorce VARCHAR(128) NULL DEFAULT NULL,
	    status INT(8) NULL DEFAULT NULL,
	    img VARCHAR(1000) NULL DEFAULT NULL,
	    level INT(8)) NULL DEFAULT NULL,
	    created_at INT(10) NULL DEFAULT NULL,
	    PRIMARY KEY(id)
	)ENGINE=InnoDB DEFAULT CHARSET=utf8;`*/

	//fmt.Println("sql语句", sql)
	err := mysql.GetDb().Set("gorm:ti_home_user", "ENGINE=InnoDB").AutoMigrate(&user)
	if err != nil {
		return err
	}

	return nil
}

// insert user
func insertUser(user User) error {
	//开启事务
	err := mysql.GetDb().Transaction(func(tx *gorm.DB) error {
		return tx.Table(table).Create(user).Error
	})
	if err != nil {
		return err
	}
	return nil
}

// GetVipUser 获取vip
// 获取数据
func GetVipUser(c echo.Context) error {
	var (
		resp common.Response
		user User
	)
	mPhone := c.QueryParam("m_phone")
	err := mysql.GetDb().Where("m_phone = ?", mPhone).Where("level = ?", 2).First(&user).Error
	if err != nil {
		fmt.Println("查询出错了")
		resp.Status = 1
		resp.Msg = "login failed"
		return c.JSON(http.StatusOK, resp)
	}
	resp.Status = 0
	resp.Msg = "ok"
	resp.Data = user
	return c.JSON(http.StatusOK, resp)
}
