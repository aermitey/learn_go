package main

type PersonalInformation struct {
	Id     int64   `json:"id,omitempty"    gorm:"primaryKey;column:id"`
	Name   string  `json:"name,omitempty"  gorm:"column:name"`
	Sex    string  `json:"sex,omitempty"   gorm:"column:sex"`
	Tall   float64 `json:"tall,omitempty"   gorm:"column:tall"`
	Weight float64 `json:"weight,omitempty"   gorm:"column:weight"`
	Age    int64   `json:"age,omitempty"   gorm:"column:age"`
}

// TableName 把结构体和数据库表关联在一起
func (p *PersonalInformation) TableName() string {
	return "personInfo"
}
