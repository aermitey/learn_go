package main

import (
	"learngo/chapter02/15.fatrate_refactor/calc"
	"log"
)

type PersonalInformation struct {
	Name   string  `json:"name"`
	Sex    string  `json:"sex"`
	Tall   float64 `json:"tall"`
	Weight float64 `json:"weight"`
	Age    int     `json:"age"`
	//BMI     float64 `json:"bmi"`
	//FatRate float64 `json:"fat_rate"`
}

func (p *PersonalInformation) CalcBMI() (bmi float64, err error) {
	bmi, err = calc.CalcBMI(p.Tall, p.Weight)
	if err != nil {
		log.Println("计算BMI出错", err)
		return
	}
	return
}

func (p *PersonalInformation) CalcFatRate() (fatRate float64, err error) {
	bmi, err := p.CalcBMI()
	if err != nil {
		log.Println("计算BMI出错", err)
		return
	}
	fatRate, err = calc.FatRateFromBMI(bmi, p.Age, p.Sex)
	if err != nil {
		log.Println("计算体脂率出错", err)
		return
	}
	return
}
