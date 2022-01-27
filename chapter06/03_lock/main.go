package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//for i := 0; i < 10; i++ {
	//	countDictGoroutineSafeRW()
	//}
	countDictGoroutineSafeRW()
}

func countDict() {
	fmt.Println("开始数")
	var totalCount int64 = 0
	wg := sync.WaitGroup{}
	wg.Add(5000)
	for i := 0; i < 5000; i++ {
		go func() {
			defer wg.Done()
			fmt.Printf("正在统计第%d页\n", i)
			totalCount += 100 //在这里有重复覆盖问题，需要引入锁来解决
		}()
	}
	wg.Wait()
	fmt.Println("预计有500000字")
	fmt.Printf("共有%d字", totalCount)
}

func countDictGoroutineSafe() {
	fmt.Println("开始数")
	var totalCount int64 = 0
	totalCountLock := sync.Mutex{}
	wg := sync.WaitGroup{}
	wg.Add(5000)
	for i := 0; i < 5000; i++ {
		go func(i int) {
			defer wg.Done()
			//defer totalCountLock.Unlock()//当不需要使用锁的时候，最好及时释放，如果一直等到最后的话很容易造成性能损失
			//fmt.Printf("正在统计第%d页\n", i)
			totalCountLock.Lock()
			totalCount += 100 //在这里有重复覆盖问题，需要引入锁来解决
			totalCountLock.Unlock()
			if i == 0 {
				time.Sleep(3 * time.Second)
			}
		}(i)
	}
	wg.Wait()
	fmt.Println("预计有500000字")
	fmt.Printf("共有%d字", totalCount)
}

func countDictForgetUnlock() {
	fmt.Println("开始数")
	var totalCount int64 = 0
	totalCountLock := sync.Mutex{}
	wg := sync.WaitGroup{}
	wg.Add(5000)
	for i := 0; i < 5000; i++ {
		go func(pageName int) {
			defer wg.Done()
			//defer totalCountLock.Unlock()
			//fmt.Printf("正在统计第%d页\n", i)
			totalCountLock.Lock()
			totalCount += 100 //在这里有重复覆盖问题，需要引入锁来解决
			if pageName%2 == 0 {
				totalCountLock.Unlock()
			}
		}(i)
	}
	wg.Wait()
	fmt.Println("预计有500000字")
	fmt.Printf("共有%d字", totalCount)
}

func countDictUnlockPrice() {
	fmt.Println("开始数")
	var totalCount int64 = 0
	totalCountLock := sync.Mutex{}
	wg := sync.WaitGroup{}
	wg.Add(5000)
	for i := 0; i < 5000; i++ {
		go func(pageName int) {
			defer wg.Done()
			defer totalCountLock.Unlock()
			//fmt.Printf("正在统计第%d页\n", i)
			totalCountLock.Lock()
			totalCount += 100 //在这里有重复覆盖问题，需要引入锁来解决
			if pageName == 3 {
				time.Sleep(3 * time.Second)
			}
		}(i)
	}
	wg.Wait()
	fmt.Println("预计有500000字")
	fmt.Printf("共有%d字", totalCount)
}

func countDictGoroutineSafeRW() { //读写互斥，读锁互不影响，读锁拿到时，需要拿写锁的goroutine等待，当写锁拿到时，读锁也等待。
	fmt.Println("开始数")
	var totalCount int64 = 0
	totalCountLock := sync.RWMutex{} //读写锁
	wg := sync.WaitGroup{}
	workerCount := 5
	wg.Add(workerCount)
	doneCh := make(chan struct{})
	for i := 0; i < workerCount; i++ {
		go func(i int) {
			fmt.Println(i, "读锁开始时间：", time.Now())
			totalCountLock.RLock()
			fmt.Println(i, "读锁持有锁的时间：", time.Now())
			time.Sleep(1 * time.Second)
			totalCountLock.RUnlock()
		}(i)
	}
	for i := 0; i < workerCount; i++ {
		go func(i int) {
			defer wg.Done()
			fmt.Println(i, "写锁开始时间：", time.Now())
			totalCountLock.Lock()
			fmt.Println(i, "写锁持有锁的时间：", time.Now())
			defer totalCountLock.Unlock()
			totalCount += 100
		}(i)
	}
	wg.Wait()
	close(doneCh)
	time.Sleep(1 * time.Second)
	fmt.Println("预计有500000字")
	fmt.Printf("共有%d字", totalCount)
}

func countDictGoroutineSafeRWChan() {
	fmt.Println("开始数")
	var totalCount int64 = 0
	totalCountLock := sync.RWMutex{} //读写锁
	wg := sync.WaitGroup{}
	workerCount := 5
	wg.Add(workerCount)
	doneCh := make(chan struct{})
	for i := 0; i < workerCount; i++ {
		go func(i int) {
			defer wg.Done()
			result := map[int64]struct{}{}
			for {
				fmt.Println(i, "写锁开始时间：", time.Now())
				totalCountLock.RLock()
				fmt.Println(i, "写锁持有锁的时间：", time.Now())
				fmt.Println(totalCount)
				result[totalCount] = struct{}{}
				time.Sleep(1 * time.Second)
				totalCountLock.RUnlock()
				needBreak := false
				select {
				case <-doneCh:
					needBreak = true
				default:

				}
				if needBreak {
					break
				}
			}
			fmt.Println("result=", result)
		}(i)
	}
	for i := 0; i < workerCount; i++ {
		go func(i int) {
			defer wg.Done()
			fmt.Println(i, "写锁开始时间：", time.Now())
			totalCountLock.Lock()
			fmt.Println(i, "写锁持有锁的时间：", time.Now())
			totalCount += 100
			totalCountLock.Unlock()
		}(i)
	}
	wg.Wait()
	close(doneCh)
	time.Sleep(1 * time.Second)
	fmt.Println("预计有500000字")
	fmt.Printf("共有%d字", totalCount)
}
