package main

import (
	"fmt"
	"reflect"
	"time"
)

func main() {
	//arr := []int{}
	startTime := time.Now()
	defer func() {
		finishTime := time.Now()
		time := finishTime.Sub(startTime)
		fmt.Println(time)
	}()

	defer func() {
		fmt.Println(1)
	}()
	arr2 := make([]int, 5, 5)
	arr2[0] = 1
	fmt.Println(arr2[4])

	arr3 := make([]int, 5, 5)
	copy(arr3, arr2)
	fmt.Println(arr3)

	i := 333
	j := &i
	fmt.Println(reflect.TypeOf(j))
}
