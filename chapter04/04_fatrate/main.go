package main

import (
	"fmt"
)

func main() {
	frSvc := &fatRateService{s: GetFatRateSuggestions()}

	for {
		person := getPersonInfoFromInput()
		fmt.Println(frSvc.GiveSuggestionToPerson(person))
	}
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
