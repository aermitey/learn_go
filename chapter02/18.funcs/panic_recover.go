package main

import "fmt"

func panicAndRecover() {
	//defer coverPanic1()
	defer func() {
		coverPanic1()
	}()
	var nameScore map[string]int = nil
	nameScore["小强"] = 100
}

func coverPanic() { //未能抓住panic
	func() {
		if r := recover(); r != nil {
			fmt.Println("系统出现严重故障")
		}
	}()
}

func coverPanic1() { //抓住panic
	if r := recover(); r != nil {
		fmt.Println("系统出现严重故障")
	}
}
