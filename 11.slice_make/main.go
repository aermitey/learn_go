package main

import "fmt"

func main() {
	{
		a := make([]int, 0, 0)
		fmt.Println(a, len(a), cap(a))
		a = append(a, 1)
		fmt.Println(a, len(a), cap(a))
	}

	{
		a := make([]int, 1, 1)
		fmt.Println(a, len(a), cap(a))
		a[0] = 2
		a = append(a, 1)
		fmt.Println(a, len(a), cap(a))
	}
}
