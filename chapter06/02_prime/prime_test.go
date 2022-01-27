package main

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

func TestRunPrime1(t *testing.T) {
	startTime := time.Now()
	var result []int
	for i := 2; i <= 200000; i++ {
		if isPrime(i) {
			result = append(result, i)
		}
	}
	finishTime := time.Now()
	fmt.Println(len(result))
	fmt.Println("共耗时：", finishTime.Sub(startTime))
}

func TestRunPrime2(t *testing.T) {
	startTime := time.Now()
	var result []int
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		fmt.Println("第一个worker开始计算", time.Now())
		result = append(result, collectPrime(2, 100000)...)
		fmt.Println("第一个worker结束计算", time.Now())
	}()
	go func() {
		defer wg.Done()
		fmt.Println("第二个worker开始计算", time.Now())
		result = append(result, collectPrime(100001, 200000)...)
		fmt.Println("第二个worker结束计算", time.Now())
	}()
	go fmt.Println("测试")
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

//func collectPrime(startNum int, endNum int) (result []int) {
//	for i := startNum; i <= endNum; i++ {
//		for i2 := 2; i2 <= endNum; i2++ {
//			if i%i2 != 0 {
//				result = append(result, i)
//			}
//		}
//	}
//	return
//}

func TestHello(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		fmt.Println("hello")
	}()
	go func() {
		defer wg.Done()
		fmt.Println("hello")
	}()
	wg.Wait()
} //主routine（main）在方法结束后退出，整个程序退出

func TestForLoop(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			fmt.Println("协程1:", i)
			time.Sleep(1 * time.Second)
		}
	}()
	go func() {
		defer wg.Done()
		for i := 0; i < 20; i++ {
			fmt.Println("协程2", i)
			time.Sleep(1 * time.Second)
		}
	}()
	wg.Wait()
	for i := 0; i < 10; i++ {
		fmt.Println("主协程:", i)
		time.Sleep(1 * time.Second)
	}
}

func TestForeverGoroutine(t *testing.T) {
	go func() {
		for {
			time.Sleep(1 * time.Second)
			go func() {
				fmt.Println("启动一个新的协程", time.Now())
				time.Sleep(1 * time.Hour)
			}()
		}
	}()
	for {
		fmt.Println(runtime.NumGoroutine())
		time.Sleep(1 * time.Second)
	}
}
