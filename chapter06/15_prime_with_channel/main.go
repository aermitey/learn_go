package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	startTime := time.Now()
	workNum := 8
	baseNumber := make(chan int, 200000)
	result := make(chan int, 200000)
	for i := 2; i <= 200000; i++ {
		baseNumber <- i
	}
	close(baseNumber)
	wg := sync.WaitGroup{}
	wg.Add(workNum)
	for i := 0; i < workNum; i++ {
		go func() {
			defer wg.Done()
			for num := range baseNumber {
				if isPrime(num) {
					result <- num
				}
			}

		}()
	}
	wg.Wait()
	finishTime := time.Now()
	fmt.Println(len(result))
	fmt.Println("共耗时：", finishTime.Sub(startTime))
}

func isPrime(num int) (isPrime bool) {
	for i := 2; i < num; i++ {
		if num%i == 0 {
			isPrime = false
			return
		}
	}
	return true
}

func collectPrime(startNum int, endNum int) (result []int) {
	for i := startNum; i <= endNum; i++ {
		if isPrime(i) {
			result = append(result, i)
		}
	}
	return
}
