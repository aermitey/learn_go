package main

import (
	"fmt"
	"time"
)

var tall float64
var weight float64

func main() {
	//panicAndRecover()
	deferGuess()
	//openFile()
	//calcSum(1, 2, 3)
	//calcSum(1, 2, 3)
	//calcSum(1, 2, 3)
	//showUsedTime()
	//fmt.Println(closureMain())
	//fmt.Println(att)
	//guess(1, 100)
	//calculatorAdd := func(a, b int) int {
	//	return a + b
	//}
	//fmt.Println(calculatorAdd(1, 2))
	//
	//{
	//	tall := 161
	//	weight := 90
	//	calculatorAdd(tall, weight)
	//}
	//{
	//	tall := 161
	//	weight := 90
	//	calculatorAdd(tall, weight)
	//}
	//sample()
	//sample2()
	//fmt.Println(fib(20))
	//c()
	//if name := "lalala"; true {
	//	a := 3
	//	name = "lala"
	//	fmt.Println(name, a)
	//} else {
	//	a := 4
	//	fmt.Println(a)
	//}
	//calcAdd()
	//tall, weight = 1, 2
	//calcAdd()
	//tall, weight := 2, 3
	//fmt.Println(tall, weight)
	//calcAdd()
	//tall, weight = 3, 4
	//calcAdd()

}

func calcAdd() float64 {
	fmt.Println(tall + weight)
	return 0
}

func sample() {
	name := "c"
	fmt.Println(name)
	{
		fmt.Println(name)
		name = "b"
		fmt.Println(name)
	}
	fmt.Println(name)
}

func sample2() {
	name := "c"
	fmt.Println(name)
	{
		name = "a"
		fmt.Println(name)
		name := "b"
		fmt.Println(name)
	}
	fmt.Println(name)
}

func fib(n uint) uint {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}

func c() {
	st := time.Now()
	defer fmt.Println(time.Now().Sub(st))
	time.Sleep(5 * time.Second)
	fmt.Println(st)
}
