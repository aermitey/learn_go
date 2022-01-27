package main

import (
	gobmi "github.com/armstrongli/go-bmi"
	"learngo/chapter02/15.fatrate_refactor/calc"
	"log"
)

type Calc struct {
	a int
}

func (c *Calc) BMI(person *Person) error {
	bmi, err := gobmi.BMI(person.weight, person.tall)
	if err != nil {
		log.Println("error when calculating bmi", err)
		return err
	}
	c.a = 2
	person.bmi = bmi
	return nil
}

func (c *Calc) FatRate(person *Person) error {
	fatRate, err := calc.CalcFatRate(person.bmi, person.age, person.sex)
	if err != nil {
		log.Println("error when calculating fatRate", err)
		return err
	}
	person.fatRate = fatRate
	return nil
}
