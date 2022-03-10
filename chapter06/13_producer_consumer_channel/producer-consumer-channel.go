package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Store struct {
	init  sync.Once
	store chan int
	Max   int
}

func (s *Store) instrument() {
	s.init.Do(func() {
		s.store = make(chan int, s.Max)
	})
}

func (s *Store) close() {

}

type Producer struct {
}

func (p Producer) Produce(s *Store) {
	fmt.Println("开始生产，+1")
	s.store <- rand.Int()
}

type Consumer struct {
}

func (c Consumer) Consume(s *Store) {
	fmt.Println("开始消费，-1")
	<-s.store
}

func main() {
	wg := sync.WaitGroup{}
	s := &Store{
		Max: 10,
	}
	s.instrument()
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
}
