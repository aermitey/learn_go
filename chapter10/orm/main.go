package main

import (
	"encoding/json"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func main() {
	conn := connectDB()
	fmt.Println(conn)
	//if err := createNewPerson(conn); err != nil {
	//	return
	//}
	//if err := ormSelect(conn); err != nil {
	//	return
	//}
	//if err := ormSelectWithInaccurateQuery(conn); err != nil {
	//	return
	//}
	//if err := ormSelectWithInaccurateQueryHack(conn); err != nil {
	//	return
	//}
	//if err := updatePerson(conn); err != nil {
	//	return
	//}
	if err := updatePersonSelectedFields(conn); err != nil {
		return
	}
	//if err := deletePerson(conn); err != nil {
	//	return
	//}

}

func connectDB() *gorm.DB {
	conn, err := gorm.Open(mysql.Open("chenxi:123456@tcp(127.0.0.1:3306)/test"))
	if err != nil {
		log.Fatal("数据库连接失败", err)
	}
	fmt.Println("连接数据库成功")
	return conn
}

func createNewPerson(conn *gorm.DB) error {
	resp := conn.Create(&PersonalInformation{
		Name:   "小强23456",
		Sex:    "男",
		Tall:   1.7,
		Weight: 71,
		Age:    35,
	})
	if err := resp.Error; err != nil {
		fmt.Println("创建失败", err)
		return err
	}
	fmt.Println("创建成功")
	return nil
}

func ormSelect(conn *gorm.DB) error {
	output := make([]*PersonalInformation, 0)
	reps := conn.Where(&PersonalInformation{Name: "陈曦"}).Find(&output)
	if err := reps.Error; err != nil {
		log.Println("查找失败", err)
		return err
	}
	data, err := json.Marshal(output)
	if err != nil {
		return err
	}
	fmt.Printf("查询结果:%+v\n", string(data))
	return nil
}

func ormSelectWithInaccurateQuery(conn *gorm.DB) error {
	output := make([]*PersonalInformation, 0)
	reps := conn.Where("name like ?", "%陈%").Find(&output)
	if err := reps.Error; err != nil {
		log.Println("查找失败", err)
		return err
	}
	data, err := json.Marshal(output)
	if err != nil {
		return err
	}
	fmt.Printf("查询结果:%+v\n", string(data))
	return nil
}

func ormSelectWithInaccurateQueryHack(conn *gorm.DB) error {
	output := make([]*PersonalInformation, 0)
	reps := conn.Where("name = ? and sex = ?", "小强23456' -- ", "男").Offset(10).Find(&output) //分页，offset，limit
	if err := reps.Error; err != nil {
		log.Println("查找失败", err)
		return err
	}
	data, err := json.Marshal(output)
	if err != nil {
		return err
	}
	fmt.Printf("查询结果:%+v\n", string(data))
	return nil
}

//主键更新
func updatePerson(conn *gorm.DB) error {
	resp := conn.Updates(&PersonalInformation{
		Id:     1,
		Name:   "陈曦",
		Sex:    "男",
		Tall:   1.7,
		Weight: 71,
		Age:    30,
	})
	if err := resp.Error; err != nil {
		fmt.Println("更新失败", err)
		return err
	}
	fmt.Println("更新成功")
	return nil
}

func updatePersonSelectedFields(conn *gorm.DB) error {
	personInfo := &PersonalInformation{
		Id:     5,
		Name:   "小强23456",
		Sex:    "女",
		Tall:   1.3,
		Weight: 75,
		Age:    35,
	}
	resp := conn.Model(personInfo).Select("name", "tall", "sex").Updates(personInfo)
	if err := resp.Error; err != nil {
		fmt.Println("更新失败", err)
		return err
	}
	fmt.Println("更新成功")
	return nil
}

func deletePerson(conn *gorm.DB) error {
	personInfo := &PersonalInformation{ //也可以只传入id
		Id:     6,
		Name:   "小强23456",
		Sex:    "男",
		Tall:   1.9,
		Weight: 75,
		Age:    35,
	}
	resp := conn.Delete(personInfo)
	if err := resp.Error; err != nil {
		fmt.Println("删除失败", err)
		return err
	}
	fmt.Println("删除成功")
	return nil
}
