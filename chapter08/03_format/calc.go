package main

import (
	gobmi "github.com/armstrongli/go-bmi"
	"learngo/chapter02/15.fatrate_refactor/calc"
	"learngo/pkg/apis"
	"log"
)

type Calc struct {
	continental string
}

func (c *Calc) BMI(person *apis.PersonalInformation) (bmi float64, err error) {
	bmi, err = gobmi.BMI(float64(person.Weight), float64(person.Tall))
	if err != nil {
		log.Println("error when calculating bmi", err)
		return 0, err
	}
	return bmi, nil
}

func (c *Calc) FatRate(person *apis.PersonalInformation) (fatRate float64, err error) {
	bmi, err := c.BMI(person)
	if err != nil {
		return 0, err
	}
	fatRate, err = calc.FatRateFromBMI(bmi, int(person.Age), person.Sex)
	if err != nil {
		log.Println("error when calculating fatRate", err)
		return 0, err
	}
	return
}
