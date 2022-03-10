package main

import "fmt"

type InputFromStd struct {
}

func (i *InputFromStd) Register() *PersonalInformation {
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

	return &PersonalInformation{
		Name:   name,
		Sex:    sex,
		Tall:   tall,
		Weight: weight,
		Age:    age,
	}
}

func (i *InputFromStd) Update() *PersonalInformation {
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

	return &PersonalInformation{
		Name:   name,
		Tall:   tall,
		Weight: weight,
		Age:    age,
	}
}

func (i *InputFromStd) fakeRegister() *PersonalInformation {
	return &PersonalInformation{
		Name:   "辰辰",
		Sex:    "女",
		Tall:   1.65,
		Weight: 80,
		Age:    27,
	}
}

func (i *InputFromStd) fakeUpdate() *PersonalInformation {
	return &PersonalInformation{
		Name:   "陈曦",
		Tall:   1.7,
		Weight: 70,
		Age:    28,
	}
}
