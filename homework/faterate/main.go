package main

import "fmt"

func main() {
	var name [3]string
	var weight [3]float64
	var tall [3]float64
	var age [3]int
	var sex [3]string
	var bmi [3]float64
	var fateRate [3]float64
	var avgRate float64

	for i := 0; i < 3; i++ {
		fmt.Print("姓名:")
		fmt.Scanln(&name[i])

		fmt.Print("体重(kg):")
		fmt.Scanln(&weight[i])

		fmt.Print("身高(m):")
		fmt.Scanln(&tall[i])

		bmi[i] = weight[i] / (tall[i] * tall[i])

		fmt.Print("年龄:")
		fmt.Scanln(&age[i])

		fmt.Print("性别(男/女):")
		fmt.Scanln(&sex[i])

		var sexWeight int

		if sex[i] == "男" {
			sexWeight = 1
		} else if sex[i] == "女" {
			sexWeight = 0
		}
		fmt.Println(sexWeight)

		fateRate[i] = (1.2*bmi[i] + 0.23*float64(age[i]) - 5.4 - 10.8*float64(sexWeight)) / 100

		fmt.Println("体脂率是", fateRate[i])

	}
	//var whetherContinue string
	//fmt.Print("是否录入下一个(y/n):")
	//fmt.Scanln(&whetherContinue)
	//if whetherContinue != "y" {
	//	break
	//}

	for i := 0; i < 3; i++ {
		fmt.Print("姓名是:", name[i], "体重是:", weight[i], "身高是:", tall[i], "年龄是:", age[i], "性别是:", sex[i], "体脂率是:", fateRate[i])
		if sex[i] == "男" {
			// 编写男性体脂率与体脂状态表
			if age[i] >= 18 && age[i] <= 39 {
				if fateRate[i] <= 0.1 {
					fmt.Println("便瘦")
				} else if fateRate[i] > 0.1 && fateRate[i] <= 0.16 {
					fmt.Println("标准")
				} else if fateRate[i] > 0.16 && fateRate[i] <= 0.21 {
					fmt.Println("偏胖")
				} else if fateRate[i] > 0.21 && fateRate[i] <= 0.26 {
					fmt.Println("肥胖")
				} else {
					fmt.Println("非常肥胖")
				}
			} else if age[i] >= 40 && age[i] <= 59 {
				if fateRate[i] <= 0.11 {
					fmt.Println("便瘦")
				} else if fateRate[i] > 0.11 && fateRate[i] <= 0.17 {
					fmt.Println("标准")
				} else if fateRate[i] > 0.17 && fateRate[i] <= 0.22 {
					fmt.Println("偏胖")
				} else if fateRate[i] > 0.22 && fateRate[i] <= 0.27 {
					fmt.Println("肥胖")
				} else {
					fmt.Println("非常肥胖")
				}
			} else if age[i] >= 60 {
				if fateRate[i] <= 0.14 {
					fmt.Println("便瘦")
				} else if fateRate[i] > 0.14 && fateRate[i] <= 0.19 {
					fmt.Println("标准")
				} else if fateRate[i] > 0.19 && fateRate[i] <= 0.24 {
					fmt.Println("偏胖")
				} else if fateRate[i] > 0.24 && fateRate[i] <= 0.29 {
					fmt.Println("肥胖")
				} else {
					fmt.Println("非常肥胖")
				}
			} else {
				fmt.Println("体脂率计算不包含未成年人")
			}
		} else if sex[i] == "女" {
			// 编写女性体脂率与体脂状态表
			if age[i] >= 18 && age[i] <= 39 {
				if fateRate[i] <= 0.2 {
					fmt.Println("便瘦")
				} else if fateRate[i] > 0.2 && fateRate[i] <= 0.27 {
					fmt.Println("标准")
				} else if fateRate[i] > 0.27 && fateRate[i] <= 0.34 {
					fmt.Println("偏胖")
				} else if fateRate[i] > 0.34 && fateRate[i] <= 0.39 {
					fmt.Println("肥胖")
				} else {
					fmt.Println("非常肥胖")
				}
			} else if age[i] >= 40 && age[i] <= 59 {
				if fateRate[i] <= 0.21 {
					fmt.Println("便瘦")
				} else if fateRate[i] > 0.21 && fateRate[i] <= 0.28 {
					fmt.Println("标准")
				} else if fateRate[i] > 0.28 && fateRate[i] <= 0.35 {
					fmt.Println("偏胖")
				} else if fateRate[i] > 0.35 && fateRate[i] <= 0.40 {
					fmt.Println("肥胖")
				} else {
					fmt.Println("非常肥胖")
				}
			} else if age[i] >= 60 {
				if fateRate[i] <= 0.22 {
					fmt.Println("便瘦")
				} else if fateRate[i] > 0.22 && fateRate[i] <= 0.29 {
					fmt.Println("标准")
				} else if fateRate[i] > 0.29 && fateRate[i] <= 0.36 {
					fmt.Println("偏胖")
				} else if fateRate[i] > 0.36 && fateRate[i] <= 0.41 {
					fmt.Println("肥胖")
				} else {
					fmt.Println("非常肥胖")
				}
			} else {
				fmt.Println("体脂率计算不包含未成年人")
			}
		}
	}
	var sumRate float64
	for i := 0; i < 3; i++ {
		sumRate += fateRate[i]
	}
	avgRate = sumRate / 3
	fmt.Println(avgRate)
}
