package main

import "fmt"

func main() {
	var a, b float64 = 30, 11
	fmt.Println("a + b = ", a+b)
	fmt.Println("a + b = ", a-b)
	fmt.Println("a + b = ", a*b) //溢出
	fmt.Println("a + b = ", a/b) //panic
	//fmt.Println("a + b = ", a%b) //不能计算小数
	fmt.Println(true && false == false)
	fmt.Println("a > b = ", a > b)
}
