package pool

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

var DB *sql.DB

func init() {
	DB, _ = sql.Open("mysql", "root:root@tcp(localhost:3306)/test?charset=utf8&parseTime=True&loc=Local") // 使用本地时间，即东八区，北京时间
	// set pool params
	DB.SetMaxOpenConns(2000)
	DB.SetMaxIdleConns(1000)
	DB.SetConnMaxLifetime(time.Minute * 60) // mysql default conn timeout=8h, should < mysql_timeout
	err := DB.Ping()
	if err != nil {
		log.Fatalf("database init failed, err: ", err)
	}
	log.Println("mysql conn pool has initiated.")
}

func checkErr(err error) {
	if err != nil {
		log.Println(err)
		panic(err)
	}
}

func createTable() {
	db := DB
	table := `CREATE TABLE IF NOT EXISTS test.user (
	 user_id INT(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '用户编号',
	 user_name VARCHAR(45) NOT NULL COMMENT '用户名称',
	 user_age TINYINT(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户年龄',
	 user_sex TINYINT(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户性别',
	 PRIMARY KEY (user_id))
	 ENGINE = InnoDB
	 AUTO_INCREMENT = 1
	 DEFAULT CHARACTER SET = utf8
	 COLLATE = utf8_general_ci
	 COMMENT = '用户表'`
	if _, err := db.Exec(table); err != nil {
		checkErr(err)
	}
}

func insert() {
	db := DB
	stmt, err := db.Prepare(`INSERT user (user, age) values (?, ?)`)
	checkErr(err)
	res, err := stmt.Exec("Elvis", 26)
	checkErr(err)
	id, err := res.LastInsertId()
	checkErr(err)
	log.Println(id)
}

func query() {
	db := DB
	rows, err := db.Query("SELECT * FROM user")
	checkErr(err)
	for rows.Next() {
		var userId int
		var userName string
		var userAge int
		var userSex int
		rows.Columns()
		err = rows.Scan(&userId, &userName, &userAge, &userSex)
		checkErr(err)
		fmt.Println(userId)
		fmt.Println(userName)
		fmt.Println(userAge)
		fmt.Println(userSex)
	}
}

func queryToMap() {
	db := DB
	rows, err := db.Query("SELECT * FROM user")
	checkErr(err)
	//字典类型
	//构造scanArgs、values两个数组，scanArgs的每个值指向values相应值的地址
	columns, _ := rows.Columns()
	scanArgs := make([]interface{}, len(columns))
	values := make([]interface{}, len(columns))
	for i := range values {
		scanArgs[i] = &values[i]
	}
	for rows.Next() {
		//将行数据保存到record字典
		err = rows.Scan(scanArgs...)
		record := make(map[string]string)
		for i, col := range values {
			if col != nil {
				record[columns[i]] = string(col.([]byte))
			}
		}
		fmt.Println(record)
	}
}

func update() {
	db := DB
	stmt, err := db.Prepare(`UPDATE user SET user_age=?,user_sex=? WHERE user_id=?`)
	checkErr(err)
	res, err := stmt.Exec(21, 2, 1)
	checkErr(err)
	num, err := res.RowsAffected()
	checkErr(err)
	fmt.Println(num)
}

func remove() {
	db := DB
	stmt, err := db.Prepare(`DELETE FROM user WHERE user_id=?`)
	checkErr(err)
	res, err := stmt.Exec(1)
	checkErr(err)
	num, err := res.RowsAffected()
	checkErr(err)
	fmt.Println(num)
}
