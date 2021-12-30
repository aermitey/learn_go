package main

import (
	"fmt"
	gobmi "github.com/armstrongli/go-bmi"
)

func main() {
	persons := &[]Person{
		{
			"小强",
			"男",
			1.7,
			70,
			35,
		},
	}
	gobmi.BMI((*persons)[0].tall, (*persons)[0].weight)
	hum := Person{
		"小强",
		"男",
		1.7,
		70,
		35,
	}
	gobmi.BMI(hum.tall, hum.weight)
	for _, person := range *persons {
		bmi, err := gobmi.BMI(person.weight, person.tall)
		if err != nil {
			fmt.Errorf("error = %v", err)
			return
		}
		fmt.Println(bmi)
	}
}

type Person struct {
	name   string
	sex    string
	tall   float64
	weight float64
	age    int
}
