package main

import (
	"fmt"
	"sync"
)

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	wg := sync.WaitGroup{}
	wg.Add(len(arr))
	for _, item := range arr {
		go func(item int) {
			defer wg.Done()
			for i := 0; i < 10; i++ {
				fmt.Println(item)
			}
		}(item)
	}
	wg.Wait()
}
