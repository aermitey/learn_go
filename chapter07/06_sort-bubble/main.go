package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	arr := make([]int, 0)
	for i := 0; i < 100000; i++ {
		arr = append(arr, rand.Intn(100))
	}
	start := time.Now()
	bubble(&arr)
	finish := time.Now()
	fmt.Println(finish.Sub(start))
}

func bubble(arr *[]int) {
	for i := 0; i < len(*arr); i++ {
		for j := 0; j < len(*arr)-i-1; j++ {
			if (*arr)[j] > (*arr)[j+1] {
				(*arr)[j], (*arr)[j+1] = (*arr)[j+1], (*arr)[j]
			}
		}
		//fmt.Println("中间状态", *arr)
	}
	//fmt.Println("最终结果", *arr)
}
