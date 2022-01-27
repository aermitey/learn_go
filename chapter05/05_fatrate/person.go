package main

import (
	gobmi "github.com/armstrongli/go-bmi"
	"learngo/chapter02/15.fatrate_refactor/calc"
	"log"
)

type Person struct {
	name    string
	sex     string
	tall    float64
	weight  float64
	age     int
	aaa     func(int) int
	bmi     float64
	fatRate float64
}

func (p *Person) calcBmi() error {
	bmi, err := gobmi.BMI(p.weight, p.tall)
	if err != nil {
		log.Printf("error when calculating bmi for Person[%s]: %v", p.name, err)
		return err
	}
	p.bmi = bmi
	return nil
}

func (p *Person) calcFatRate() error {
	fatRate, err := calc.CalcFatRate(p.bmi, p.age, p.sex)
	if err != nil {
		log.Printf("error when calculating fatRate for Person[%s]: %v", p.name, err)
	}
	p.fatRate = fatRate
	return nil
}
