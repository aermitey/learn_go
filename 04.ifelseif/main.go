package main

import "fmt"

func main() {
	//if只会走符合条件的那一个分支
	var money int = 22
	if money <= 20{
		//如果有不超过20，点外卖
		fmt.Println("点个外卖")
	}else if money <= 2000 && money > 200 {
		//如果不超过2000，去米其林
		fmt.Println("去米其林")
	}else if money <= 200 {
		//如果不超过200，下馆子
		fmt.Println("下馆子")
	}else if money <= 20000{
		//如果不超过20000，去新马泰
		fmt.Println("去新马泰")
	}else {
		//如果再多，容我想想
		fmt.Println("容我想想")
	}
	fmt.Println("end")
}
