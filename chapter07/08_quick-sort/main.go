package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	arr := make([]int, 0)
	size := 10
	for i := 0; i < size; i++ {
		arr = append(arr, rand.Intn(100))
	}
	fmt.Println("初始数组：", arr)

	start := time.Now()
	quickSort(&arr, 0, size-1)
	finish := time.Now()
	fmt.Println(finish.Sub(start))
	fmt.Println("最终数组：", arr)
}

func quickSort(arr *[]int, start, end int) {
	//确认终止条件，否则将无限递归下去
	pivotIdx := (start + end) / 2
	pivotValue := (*arr)[pivotIdx]
	l, r := start, end
	for l < r {
		for (*arr)[l] < pivotValue {
			l++
		}
		for (*arr)[r] > pivotValue {
			r--
		}
		if l >= r {
			break
		}
		(*arr)[l], (*arr)[r] = (*arr)[r], (*arr)[l]
		l++
		r--
	}
	//fmt.Println("l:", l, "r:", r)
	//fmt.Println(*arr)
	if l == r {
		l++
		r--
	}
	if r > start {
		quickSort(arr, start, r)
	}
	if l < end {
		quickSort(arr, l, end)
	}
}
