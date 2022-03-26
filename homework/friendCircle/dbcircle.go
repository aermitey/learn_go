package main

import (
	"fmt"
	"gorm.io/gorm"
	"learngo/chapter02/15.fatrate_refactor/calc"
	"learngo/pkg/apis"
	"log"
	"time"
)

var _ CircleInterface = &FriendCircleDB{}

type FriendCircleDB struct {
	conn *gorm.DB
}

func NewFriendCircleDB(conn *gorm.DB) CircleInterface {
	if conn == nil {
		log.Fatal("数据库连接为空")
	}
	return &FriendCircleDB{
		conn: conn,
	}
}

func (f *FriendCircleDB) PostingStatus(fc *apis.FriendCircle) error {
	pi := &apis.PersonalInformation{}
	resp := f.conn.Where(&apis.PersonalInformation{Id: fc.PersonId}).Find(pi)
	if err := resp.Error; err != nil {
		log.Println("该用户不存在,请先注册", err)
	}
	bmi, err := calc.CalcBMI(float64(pi.Tall), float64(pi.Weight))
	if err != nil {
		log.Println("计算bmi失败", err)
		return err
	}
	fr, err := calc.FatRateFromBMI(bmi, int(pi.Age), pi.Sex)
	if err != nil {
		log.Println("计算体脂率失败", err)
		return err
	}
	resp = f.conn.Create(&apis.FriendCircle{
		PersonId:   pi.Id,
		Name:       pi.Name,
		Sex:        pi.Sex,
		Tall:       pi.Tall,
		Weight:     pi.Weight,
		Age:        pi.Age,
		FatRate:    float32(fr),
		Content:    fc.Content,
		CreateTime: fmt.Sprintf("%v", time.Now()),
		IsDeleted:  1,
	})
	if err = resp.Error; err != nil {
		log.Println("发布状态失败", err)
		return err
	}
	return nil
}

func (f *FriendCircleDB) DeleteStatus(id int) error {
	fc := &apis.FriendCircle{Id: int64(id)}
	resp := f.conn.Model(fc).Update("isDeleted", 2)
	if err := resp.Error; err != nil {
		log.Println("删除失败", err)
		return err
	}
	log.Println("删除成功")
	return nil
}

func (f *FriendCircleDB) Browse() ([]*apis.FriendCircle, error) {
	output := make([]*apis.FriendCircle, 0)
	resp := f.conn.Where(&apis.FriendCircle{IsDeleted: 1}).Find(&output)
	if err := resp.Error; err != nil {
		log.Println("展示朋友圈失败", err)
		return nil, err
	}
	return output, nil
}
