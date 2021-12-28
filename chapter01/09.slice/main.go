package main

import "fmt"

func main() {
	a := []int{} //切片
	fmt.Println(a)
	a = append(a, 333) //追加
	fmt.Println(a)
	b := [0]int{} //数组
	fmt.Println(b)

	xqInfo := []string{"小强", "男", "在职"}
	fmt.Println(xqInfo)
	for i, i2 := range xqInfo {
		fmt.Println(i, i2)
	}

	//a = []int{111, 222, 333, 444, 555}
	//a = append(a[:2], a[3:]...) //删除333
	//fmt.Println(a)
	//a = []int{111, 222, 333, 444, 555}
	//a = append(a[:1], a[2:4]...)
	//fmt.Println(a)
	//a = []int{111, 222, 333, 444, 555}
	//a = append(a[:])
	//fmt.Println(a)
	//a = []int{111, 222, 333, 444, 555}
	//a = append(a, a[:]...)
	//fmt.Println(a)
	a = []int{111, 222, 333, 444, 555}
	c := a[1:]
	fmt.Println(c)
	a = append(a[:1], 999)
	fmt.Println(a)
	fmt.Println(c)
	a = append(a, c...)
	fmt.Println(a)
}
