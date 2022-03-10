package main

import (
	"fmt"
	"learngo/pkg/apis/proto"
)

type inputFromStd struct {
}

func (i *inputFromStd) GetInPut() *proto.PersonalInformation {
	//录入各项数据
	var name string
	fmt.Print("姓名：")
	fmt.Scanln(&name)

	var weight float64
	fmt.Print("体重（千克）：")
	fmt.Scanln(&weight)

	var tall float64
	fmt.Print("身高（米）：")
	fmt.Scanln(&tall)

	var age int
	fmt.Print("年龄：")
	fmt.Scanln(&age)

	var sex string
	fmt.Print("性别（男/女）：")
	fmt.Scanln(&sex)

	return &proto.PersonalInformation{
		Name:   name,
		Sex:    sex,
		Tall:   float32(tall),
		Weight: float32(weight),
		Age:    int64(age),
	}
}
