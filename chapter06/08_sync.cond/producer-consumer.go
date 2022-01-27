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
	pCond     *sync.Cond
	cCond     *sync.Cond
}

type Producer struct {
}

func (p Producer) Produce(s *Store) {
	s.lock.Lock()
	defer s.lock.Unlock()
	if s.DataCount == s.Max {
		fmt.Println("库存满了，生产者等待消费者清理库存")
		s.pCond.Wait()
	}
	fmt.Println("开始生产，+1")
	s.DataCount++
	s.cCond.Signal()
}

type Consumer struct {
}

func (c Consumer) Consume(s *Store) {
	s.lock.Lock()
	defer s.lock.Unlock()
	if s.DataCount == 0 {
		fmt.Println("没库存了，消费者等待生产者生产")
		s.cCond.Wait() //锁已经还回去了
	}
	fmt.Println("开始消费，-1")
	s.DataCount--
	s.pCond.Signal()
}

func main() {
	wg := sync.WaitGroup{}
	s := &Store{
		Max: 10,
	}
	s.cCond = sync.NewCond(&s.lock)
	s.pCond = sync.NewCond(&s.lock) //特殊的初始化
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
