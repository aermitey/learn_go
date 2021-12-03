package main

import "fmt"

func main() {
	var fruit string = "6, apples"
	var watermellan bool = true
	if watermellan { //if不满足条件时不执行
		fruit = "1, apple"
	}else {
		fruit = "7, apples"
	}
	fmt.Println("buy", fruit)
}
