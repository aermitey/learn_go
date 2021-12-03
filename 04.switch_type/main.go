package main

import "fmt"

func main() {
	var money interface{} = "10"
	switch money.(type) {
	case int:
		fmt.Println("int")
	case int64:
		fmt.Println("int64")
	default:
		fmt.Println("未知类型")
	}
	fmt.Println("结束")


}
