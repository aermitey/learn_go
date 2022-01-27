package main

import (
	"fmt"
	"sync"
	"time"
)

type Store struct {
	DataCount int
	Max       int
	lock      sync.Mutex
}

type Producer struct {
}

func (p Producer) Produce(s *Store) {
	s.lock.Lock()
	defer s.lock.Unlock()
	if s.DataCount == s.Max {
		fmt.Println("生产者看到库存满了，不生产")
		return
	}
	fmt.Println("开始生产，+1")
	s.DataCount++
}

type Consumer struct {
}

func (c Consumer) Consume(s *Store) {
	s.lock.Lock()
	defer s.lock.Unlock()
	if s.DataCount == 0 {
		fmt.Println("消费者看到没有库存，不消费")
		return
	}
	fmt.Println("开始消费，-1")
	s.DataCount--
}

func main() {
	wg := sync.WaitGroup{}
	s := &Store{
		Max: 10,
	}

	pCount, cCount := 50, 50
	count := pCount + cCount
	wg.Add(count)
	for i := 0; i < pCount; i++ {
		go func() {
			defer wg.Done()
			time.Sleep(100 * time.Millisecond)
			Producer{}.Produce(s)
		}()
	}
	for i := 0; i < cCount; i++ {
		go func() {
			defer wg.Done()
			time.Sleep(100 * time.Millisecond)
			Consumer{}.Consume(s)
		}()
	}
	wg.Wait()
	fmt.Println(s.DataCount)
}
