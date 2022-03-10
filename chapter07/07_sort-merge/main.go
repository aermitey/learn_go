package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	arr := make([]int, 0)
	size := 1000000
	for i := 0; i < size; i++ {
		arr = append(arr, rand.Intn(100))
	}
	//fmt.Println(arr)
	//fmt.Println(mergeSort(arr))
	start := time.Now()
	mergeSort(arr)
	finish := time.Now()
	fmt.Println(finish.Sub(start))
}

func mergeSort(arr []int) []int {
	left, right := arr[:len(arr)/2], arr[len(arr)/2:]
	if len(arr) <= 2 {
		return mergeArr(left, right)
	} else {
		leftSort := mergeSort(left)
		rightSort := mergeSort(right)
		mergedSort := mergeArr(leftSort, rightSort)
		return mergedSort
	}
}

func mergeArr(left, right []int) []int {
	out := make([]int, 0)
	leftI, rightI := 0, 0
	for {
		if leftI == len(left) || rightI == len(right) {
			break
		}
		if left[leftI] < right[rightI] {
			out = append(out, left[leftI])
			leftI++
		} else {
			out = append(out, right[rightI])
			rightI++
		}
	}
	for ; leftI < len(left); leftI++ {
		out = append(out, left[leftI])
	}
	for ; rightI < len(right); rightI++ {
		out = append(out, right[rightI])
	}
	return out
}
