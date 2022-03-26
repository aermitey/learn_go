package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"learngo/pkg/apis"
	"log"
	"time"
)

func main() {
	testDB, err := getDBConnection()
	defer testDB.Close()
	err = testDBConnection(err, testDB)
	insertData(testDB)
	//queryAllData(err, testDB)
	queryAllDataWithHack(err, testDB)
}

func insertData(testDB *sql.DB) error {
	stmt, err := testDB.Prepare("insert into personInfo(name, sex, tall, weight, age) values ('%s','%s','%f','%f','%d')")
	if err != nil {
		return err
	}
	stmt.Query("小贝",
		"女",
		1.7,
		68.0,
		19)
	_, err = testDB.Exec(fmt.Sprintf("insert into personInfo(name, sex, tall, weight, age) values ('%s','%s','%f','%f','%d')",
		time.Now().String(),
		"女",
		1.7,
		68.0,
		19))
	if err != nil {
		log.Println("新增数据失败", err)
		return err
	}
	return nil
}

func queryAllData(err error, testDB *sql.DB) {
	rows, err := testDB.Query("select name, age from personInfo where name = '陈曦' and sex = '男'")
	if err != nil {
		log.Println("查询失败")
		return
	}
	list := &apis.PersonalInformationList{}
	for rows.Next() {
		var name string
		var age int
		err := rows.Scan(&name, &age) //要和查询语句的列的顺序相同
		if err != nil {
			return
		}
		list.Items = append(list.Items, &apis.PersonalInformation{
			Name: name,
			Age:  int64(age),
		})
	}
	data, _ := json.Marshal(list)
	fmt.Println(string(data))
}

func queryAllDataWithHack(err error, testDB *sql.DB) error {
	sqlQuery := fmt.Sprintf(`select name, age from personInfo where name = '%s' and sex = '%s'`, "小强23456' -- ", "男")
	rows, err := testDB.Query(sqlQuery)
	if err != nil {
		log.Println("查询失败", err)
		return err
	}
	list := &apis.PersonalInformationList{}
	for rows.Next() {
		var name string
		var age int
		err := rows.Scan(&name, &age) //要和查询语句的列的顺序相同
		if err != nil {
			return err
		}
		list.Items = append(list.Items, &apis.PersonalInformation{
			Name: name,
			Age:  int64(age),
		})
	}
	data, _ := json.Marshal(list)
	fmt.Println(string(data))
	return nil
}

func testDBConnection(err error, testDB *sql.DB) error {
	if err = testDB.Ping(); nil != err {
		log.Fatal("DB测试失败", err)
	}
	fmt.Println("数据库连接成功")
	return err
}

func getDBConnection() (*sql.DB, error) {
	testDB, err := sql.Open("mysql", "chenxi:123456@tcp(127.0.0.1:3306)/test")
	if err != nil {
		log.Fatal("连接数据库失败:", err)
	}
	return testDB, err
}
