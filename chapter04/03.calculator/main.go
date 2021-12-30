package main

import "fmt"

func main() {
	a := new(Calculator)
	b := Calculator{
		2,
	}
	a.A = 1

	fmt.Println(Calculator{})
	fmt.Println(c{})
	fmt.Println(b)

}

type c struct {
	calc *Calculator
}
type Calculator struct {
	A int
}
