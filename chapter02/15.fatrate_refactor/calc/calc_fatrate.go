package calc

import "fmt"

func FatRateFromBMI(bmi float64, age int, sex string) (fatRate float64, err error) {
	sexWeight := 0
	if sex == "男" {
		sexWeight = 1
	} else if sex == "女" {
		sexWeight = 0
	} else {
		err = fmt.Errorf("未识别性别:%s", sex)
		return
	}
	ageWeight, err := getAgeWeight(age)
	if err != nil {
		return
	}
	fatRate = (1.2*bmi + ageWeight*float64(age) - 5.4 - 10.8*float64(sexWeight)) / 100
	return
}

func getAgeWeight(age int) (ageWeight float64, err error) {
	if age < 18 {
		err = fmt.Errorf("不计算未成年人")
		return
	} else if age >= 18 && age < 30 {
		ageWeight = 0.23
	} else if age >= 30 && age < 40 {
		ageWeight = 0.22
	} else if age >= 40 && age < 60 {
		ageWeight = 0.21
	} else if age >= 60 && age < 80 {
		ageWeight = 0.21
	} else if age >= 80 && age < 150 {
		ageWeight = 0.20
	} else {
		err = fmt.Errorf("不计算修仙者")
		return
	}
	return
}
