package main

import "fmt"

var counter int
var att int = 3

func calcSum(nums ...int) (sum int) {
	for _, item := range nums {
		sum += item
	}
	counter++
	fmt.Println(sum)
	return sum
}

func showUsedTime() {
	fmt.Println(counter)
}

func genImprovementFunc(base float64) func(percentage float64) float64 {
	return func(percentage float64) float64 {
		base = base * (1 + percentage)
		return base
	}
}

func closureMain() (a float64, b float64, c float64, d float64, e float64) {
	imp := genImprovementFunc(1000)
	a = imp(0.1)
	b = imp(0.1)
	c = imp(0.1)
	d = imp(0.1)
	e = imp(0.1)
	return
}
