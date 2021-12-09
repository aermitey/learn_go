package main

import "fmt"

func main() {
	var money = 30
	var busy bool = false
	//分支表达式可以完全替代ifelse
	switch {
	case money >= 0 && money < 200:
		fmt.Println("点个外卖")
		fallthrough
	case money >= 200:
		fmt.Println("下个馆子")
		if busy {
			break
		} else {
			fmt.Printf("再吃点零食")
		}
		//default:
		//	fmt.Println("容我想想")
	}
	fmt.Printf("end")
}
