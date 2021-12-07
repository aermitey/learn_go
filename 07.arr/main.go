package main

import "fmt"

func main() {
	var age int = 26
	fmt.Println(age)
	var ages [5]int = [5]int{23, 24, 25, 26, 27}
	fmt.Println(ages)
	var age1 = [5]int{23, 25, 27, 28, 29}
	fmt.Println(age1)
	age2 := [5]int{1, 2, 3, 4, 5}
	age2 = age1
	fmt.Println(age2) //数组长度不能变,固定类型
	age3 := [...]int{1, 2, 3, 5, 7}
	fmt.Println(age3)
	age3 = age1
	age4 := [...]int{}
	fmt.Println(age4)
	age5 := [3]int{}
	age5[0] = 1
	age5[1] = -1
	fmt.Println(age5)
	fmt.Println(len(age5)) //不可越界为元素赋值

	for i := 0; i < len(age3); i++ {
		fmt.Println(age3[i])
	}

	for i, val := range age3 {
		fmt.Println(age3[i], "===>", i, "->", val)
	}
}
