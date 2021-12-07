package main

import "fmt"

func main() {
	for {
		var name string
		fmt.Print("姓名:")
		fmt.Scanln(&name, &name, &name)
		var weight float64
		fmt.Print("体重(kg):")
		fmt.Scanln(&weight, &weight, &weight)
		var tall float64
		fmt.Print("身高(m):")
		fmt.Scanln(&tall, tall, tall)
		var bmi float64 = weight / (tall * tall)
		var age int
		fmt.Print("年龄:")
		fmt.Scanln(&age)
		var sexWeight int
		var sex string
		fmt.Print("性别(男/女):")
		fmt.Scanln(&sex)
		var fateRate float64 = (1.2*bmi + 0.23*float64(age) - 5.4 - 10.8*float64(sexWeight)) / 100
		if sex == "男" {
			sexWeight = 1
		} else if sex == "女" {
			sexWeight = 0
		}
		fmt.Println("体脂率是", fateRate)

		if sex == "男" {
			// 编写男性体脂率与体脂状态表
			if age >= 18 && age <= 39 {
				if fateRate <= 0.1 {
					fmt.Println("便瘦")
				} else if fateRate > 0.1 && fateRate <= 0.16 {
					fmt.Println("标准")
				} else if fateRate > 0.16 && fateRate <= 0.21 {
					fmt.Println("偏胖")
				} else if fateRate > 0.21 && fateRate <= 0.26 {
					fmt.Println("肥胖")
				} else {
					fmt.Println("非常肥胖")
				}
			} else if age >= 40 && age <= 59 {
				//todo
			} else if age >= 60 {
				//todo
			} else {
				fmt.Println("体脂率计算不包含未成年人")
			}
		} else if sex == "女" {
			// 编写女性体脂率与体脂状态表
		}
		var whetherContinue string
		fmt.Print("是否录入下一个(y/n):")
		fmt.Scanln(&whetherContinue)
		if whetherContinue != "y" {
			break
		}
	}
}
