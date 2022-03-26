package main

import (
	"gorm.io/gorm"
	"learngo/chapter12/02.practice/frInterface"
	"learngo/pkg/apis"
	"log"
)

var _ RankInterface = &dbRank{}
var _ RankInitInterface = &dbRank{}

func NewDBRank(conn *gorm.DB, embedRank frInterface.ServeInterface) frInterface.ServeInterface {
	if conn == nil {
		log.Fatal("数据库连接为空")
	}
	if embedRank == nil {
		log.Fatal("内存排行榜为空")
	}
	return &dbRank{
		conn:      conn,
		embedRank: embedRank,
	}
}

type dbRank struct {
	conn      *gorm.DB
	embedRank frInterface.ServeInterface
}

func (d *dbRank) Init() error {
	output := make([]*apis.PersonalInformation, 0)
	resp := d.conn.Find(&output)
	if err := resp.Error; err != nil {
		log.Println("初始化时查找失败", err)
		return err
	}
	for _, item := range output {
		if _, err := d.embedRank.UpdatePersonalInformation(item); err != nil {
			log.Fatalf("初始化%s失败", item.Name)
		}
	}
	return nil
}

func (d *dbRank) RegisterPersonalInformation(pi *apis.PersonalInformation) error {
	resp := d.conn.Create(&pi)
	if err := resp.Error; err != nil {
		//不同企业对log有要求，比如：必须带上某个id，log时使用公司各自的log框架
		log.Printf("创建%s时失败:%v", pi.Name, err)
		return err
	}
	log.Printf("创建%s成功\n", pi.Name)
	err := d.embedRank.RegisterPersonalInformation(pi)
	if err != nil {
		log.Println("注册时更新内存失败")
		return err
	}
	return nil
}

func (d *dbRank) UpdatePersonalInformation(pi *apis.PersonalInformation) (*apis.PersonalInformationFatRate, error) {
	resp := d.conn.Updates(pi)
	if err := resp.Error; err != nil {
		log.Printf("更新%s时失败:%v", pi.Name, err)
		return nil, err
	}
	log.Printf("更新%s成功\n", pi.Name)
	return d.embedRank.UpdatePersonalInformation(pi)
}

func (d *dbRank) GetFatRate(name string) (*apis.PersonalRank, error) {
	return d.embedRank.GetFatRate(name)
}

func (d *dbRank) GetTop() ([]*apis.PersonalRank, error) {
	return d.embedRank.GetTop()
}
