package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Runner struct {
	Name string
}

func (r Runner) Run(startPointWg *sync.WaitGroup, wg *sync.WaitGroup) { //这里需要传入指针，不然wg.Wait()会报错
	defer wg.Done()
	startPointWg.Wait()
	startTime := time.Now()
	fmt.Println(r.Name, "开始跑", startTime)
	rand.Seed(time.Now().UnixNano())
	time.Sleep(time.Duration(rand.Uint64()%10) * time.Second)
	finishTime := time.Now()
	fmt.Println(r.Name, "跑到终点", finishTime.Sub(startTime))
}

func main() {
	runnerCount := 10
	var runners []Runner
	wg := sync.WaitGroup{}
	wg.Add(runnerCount)
	startPointWg := sync.WaitGroup{}
	startPointWg.Add(1)
	for i := 0; i < runnerCount; i++ {
		runners = append(runners, Runner{
			Name: fmt.Sprintf("%d", i),
		})
	}

	for _, runnerItem := range runners {
		go runnerItem.Run(&startPointWg, &wg)
	}
	fmt.Println("各就位，预备")
	time.Sleep(1 * time.Second)
	fmt.Println("跑")
	startPointWg.Done() //准备一个信号，同步做一个事情，此时Wait()在run方法内部，协程启动之后，收到了done才会真正做事情
	wg.Wait()
	fmt.Println("赛跑结束")
}
