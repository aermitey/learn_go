package main

import "fmt"

func main() {
	a := "hello"
	fmt.Println(a)
	aByte := []byte(a)
	fmt.Println(aByte, len(aByte))
	aByte[0] = 'H'
	fmt.Println(aByte)
	a = string(aByte)
	fmt.Println(a)

	b := "您好"
	fmt.Println(b)
	bByte := []rune(b)
	fmt.Println(bByte, len(bByte))
	bByte[0] = '你'
	b = string(bByte)
	fmt.Println(b)
}
