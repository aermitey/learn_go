package main

import "fmt"

func main() {
	hello("chen")
	msg, drink := constructHelloMessage("lalala")
	fmt.Println(msg, drink)
}

func hello(name string) {
	fmt.Println("Hello,", name)
	return
}

func constructHelloMessage(name string) (string, string) {
	//fmt.Println("你好,", name)     //给终端
	return "你好," + name + "再来一杯", "bear" //返回值给程序用
}

func calcBMI(tall float64, weight float64) float64 {
	return weight / (tall * tall)
}
