package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestWebServerV1(t *testing.T) {
	c := Config{
		Content: "a",
	}
	v1 := &WebServerV1{
		config: c,
	}
	go v1.ReloadWorker()

	start := time.Now()
	wg := sync.WaitGroup{}
	wg.Add(10000)
	for i := 0; i < 10000; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				v1.Visit()
			}
		}()
	}
	wg.Wait()
	finish := time.Now()
	fmt.Println("总时间", finish.Sub(start))
	fmt.Println(c, v1.config)
}

func TestWebServerV2(t *testing.T) {
	c := &Config{
		Content: "a",
	}
	v2 := &WebServerV2{
		c,
	}

	go v2.ReloadWorker()

	start := time.Now()
	wg := sync.WaitGroup{}
	wg.Add(10000)
	for i := 0; i < 10000; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 10000; j++ {
				v2.Visit()
			}
		}()
	}
	wg.Wait()
	finish := time.Now()
	fmt.Println("总时间", finish.Sub(start))
	fmt.Printf("c=%v, v2=%v\n", c, v2)
}
