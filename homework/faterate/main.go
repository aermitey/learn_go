package main

import "fmt"

func main() {
	for {
		var name [3]string
		fmt.Print("姓名:")
		fmt.Scanln(&name[0])
		fmt.Scanln(&name[1])
		fmt.Scanln(&name[2])
		var weight [3]float64
		fmt.Print("体重(kg):")
		fmt.Scanln(&weight[0])
		fmt.Scanln(&weight[1])
		fmt.Scanln(&weight[2])
		var tall [3]float64
		fmt.Print("身高(m):")
		fmt.Scanln(&tall[0])
		fmt.Scanln(&tall[1])
		fmt.Scanln(&tall[2])
		var bmi [3]float64
		for i := 0; i < 3; i++ {
			bmi[i] = weight[i] / (tall[i] * tall[i])
		}
		var age [3]int
		fmt.Print("年龄:")
		fmt.Scanln(&age[0])
		fmt.Scanln(&age[1])
		fmt.Scanln(&age[2])
		var sex [3]string
		fmt.Print("性别(男/女):")
		fmt.Scanln(&sex[0])
		fmt.Scanln(&sex[1])
		fmt.Scanln(&sex[2])
		var sexWeight int
		for i := 0; i < 3; i++ {
			if sex[i] == "男" {
				sexWeight = 1
			} else if sex[i] == "女" {
				sexWeight = 0
			}
		}
		var fateRate [3]float64
		for i := 0; i < 3; i++ {
			fateRate[i] = (1.2*bmi[i] + 0.23*float64(age[i]) - 5.4 - 10.8*float64(sexWeight)) / 100
		}

		fmt.Println("体脂率是", fateRate)
		for i := 0; i < 3; i++ {
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
		var whetherContinue string
		fmt.Print("是否录入下一个(y/n):")
		fmt.Scanln(&whetherContinue)
		if whetherContinue != "y" {
			break
		}
	}
}
