package main

import (
	gobmi "github.com/armstrongli/go-bmi"
	"log"
)

type Calc struct {
}

func (c *Calc) BMI(person *Person) error {
	bmi, err := gobmi.BMI(person.weight, person.tall)
	if err != nil {
		log.Println("error when calculating bmi", err)
		return err
	}
	person.bmi = bmi
	return nil
}

func (c *Calc) FatRate(person *Person) error {
	fatRate, err := CalcFatRate(person.bmi, person.age, person.sex)
	if err != nil {
		log.Println("error when calculating fatRate", err)
		return err
	}
	person.fatRate = fatRate
	return nil
}
