package main

import "fmt"

type FatRateSuggestion struct {
}

func (FatRateSuggestion) GetSuggestion(person *Person) (suggestion string, err error) {
	if person.fatRate <= 0 {
		err = fmt.Errorf("体脂率不正确")
		return
	}
	if person.sex == "男" {
		// 编写男性体脂率与体脂状态表
		if person.age >= 18 && person.age <= 39 {
			if person.fatRate <= 0.1 {
				suggestion = "便瘦"
			} else if person.fatRate > 0.1 && person.fatRate <= 0.16 {
				suggestion = "标准"
			} else if person.fatRate > 0.16 && person.fatRate <= 0.21 {
				suggestion = "偏胖"
			} else if person.fatRate > 0.21 && person.fatRate <= 0.26 {
				suggestion = "肥胖"
			} else {
				suggestion = "非常肥胖"
			}
		} else if person.age >= 40 && person.age <= 59 {
			if person.fatRate <= 0.11 {
				suggestion = "便瘦"
			} else if person.fatRate > 0.11 && person.fatRate <= 0.17 {
				suggestion = "标准"
			} else if person.fatRate > 0.17 && person.fatRate <= 0.22 {
				suggestion = "偏胖"
			} else if person.fatRate > 0.22 && person.fatRate <= 0.27 {
				suggestion = "肥胖"
			} else {
				suggestion = "非常肥胖"
			}
		} else if person.age >= 60 {
			if person.fatRate <= 0.14 {
				suggestion = "便瘦"
			} else if person.fatRate > 0.14 && person.fatRate <= 0.19 {
				suggestion = "标准"
			} else if person.fatRate > 0.19 && person.fatRate <= 0.24 {
				suggestion = "偏胖"
			} else if person.fatRate > 0.24 && person.fatRate <= 0.29 {
				suggestion = "肥胖"
			} else {
				suggestion = "非常肥胖"
			}
		}
	} else if person.sex == "女" {
		// 编写女性体脂率与体脂状态表
		if person.age >= 18 && person.age <= 39 {
			if person.fatRate <= 0.2 {
				suggestion = "便瘦"
			} else if person.fatRate > 0.2 && person.fatRate <= 0.27 {
				suggestion = "标准"
			} else if person.fatRate > 0.27 && person.fatRate <= 0.34 {
				suggestion = "偏胖"
			} else if person.fatRate > 0.34 && person.fatRate <= 0.39 {
				suggestion = "肥胖"
			} else {
				suggestion = "非常肥胖"
			}
		} else if person.age >= 40 && person.age <= 59 {
			if person.fatRate <= 0.21 {
				suggestion = "便瘦"
			} else if person.fatRate > 0.21 && person.fatRate <= 0.28 {
				suggestion = "标准"
			} else if person.fatRate > 0.28 && person.fatRate <= 0.35 {
				suggestion = "偏胖"
			} else if person.fatRate > 0.35 && person.fatRate <= 0.40 {
				suggestion = "肥胖"
			} else {
				suggestion = "非常肥胖"
			}
		} else if person.age >= 60 {
			if person.fatRate <= 0.22 {
				suggestion = "便瘦"
			} else if person.fatRate > 0.22 && person.fatRate <= 0.29 {
				suggestion = "标准"
			} else if person.fatRate > 0.29 && person.fatRate <= 0.36 {
				suggestion = "偏胖"
			} else if person.fatRate > 0.36 && person.fatRate <= 0.41 {
				suggestion = "肥胖"
			} else {
				suggestion = "非常肥胖"
			}
		}
	}
	return
}
