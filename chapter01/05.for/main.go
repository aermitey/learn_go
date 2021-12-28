package main

import "fmt"

func main() {
	fmt.Println("round1")
	for i := 0; i < 100; i++ {
		fmt.Println("gogogo", i)
	}
	fmt.Println("round2")
	j := 0
	for ; j < 5; j++ {
		fmt.Println("hello")
	}
	fmt.Println("round3")
	//k := 0
	//for ; true; k++ {
	//	fmt.Println("lalala")
	//}
	fmt.Println("round4")
	l := 0
	for l < 5 {
		fmt.Println("ooo")
		l++
	}
	fmt.Println("round5")
	m := 0
	for {
		fmt.Println("jiejiejie")
		m++
		if m >= 3 {
			break
		}
	}
	fmt.Println("round5")
	n := 0
	for {
		fmt.Println("chaochaochao")
		n++
		if n >= 10 {
			break
		}
		if n%2 == 0 {
			fmt.Println("jump", n)
			continue
		}
		fmt.Println("ooo", n)
	}
}
