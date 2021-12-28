package main

import (
	"fmt"
)

func main() {
	person := getPersonInfoFromInput()
	c := Calc{}
	s := FatRateSuggestion{}
	c.BMI(person)
	c.FatRate(person)
	suggestion, err := s.GetSuggestion(person)
	if err != nil {
		return
	}
	fmt.Println(person)
	fmt.Println(suggestion)
}

func getPersonInfoFromInput() *Person {
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

	return &Person{
		name:   name,
		sex:    sex,
		tall:   tall,
		weight: weight,
		age:    age,
	}
}
