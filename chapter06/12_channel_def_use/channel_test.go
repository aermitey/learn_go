package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

//不带有buffer的chan不可以在主协程接收或释放数据，会造成阻塞
func TestDefChannel(t *testing.T) {
	var sampleMap map[string]int = map[string]int{} //除了make还可以这样实例化
	fmt.Println(sampleMap)
	var intCh chan int
	fmt.Println(intCh)
	intCh = make(chan int, 10) //只能通过make实例化
	fmt.Println(intCh)
	intCh <- 3
	fmt.Println(intCh)
	<-intCh
	fmt.Println(intCh)
}

func TestChanPutGet(t *testing.T) {
	intCh := make(chan int) //创建一个不带size的channel(不带buffer)
	workerCount := 10
	for i := 0; i < workerCount; i++ {
		go func(i int) {
			intCh <- i
		}(i)
	}
	for o := 0; o < workerCount; o++ {
		go func(o int) {
			out := <-intCh
			fmt.Printf("出%d拿到%d\n", o, out)
		}(o)
	}
	time.Sleep(1 * time.Second)
}

//让out部分等待一段时间再取数据，来观察i的行为
//如果没有out，那么in部分会等待，直到有out开始
func TestChanPutGet2_oWait(t *testing.T) {
	intCh := make(chan int) //创建一个不带size的channel(不带buffer)
	workerCount := 10
	for i := 0; i < workerCount; i++ {
		go func(i int) {
			fmt.Println(i, "开始工作", time.Now())
			intCh <- i
			fmt.Println(i, "结束工作", time.Now())
		}(i)
	}
	fmt.Println(time.Now())
	time.Sleep(3 * time.Second)
	fmt.Println(time.Now())
	for o := 0; o < workerCount; o++ {
		go func(o int) {
			out := <-intCh
			fmt.Printf("%s出%d拿到%d\n", time.Now(), o, out)
		}(o)
	}
	time.Sleep(1 * time.Second)
}

//让out部分等待一段时间再取数据，来观察i的行为
//带有buffer的会允许in，不影响out
func TestChanPutGet2_oWait_withBuffer(t *testing.T) {
	intCh := make(chan int, 10)
	workerCount := 10
	for i := 0; i < workerCount; i++ {
		go func(i int) {
			fmt.Println(i, "开始工作", time.Now())
			intCh <- i
			fmt.Println(i, "结束工作", time.Now())
		}(i)
	}
	fmt.Println(time.Now())
	time.Sleep(3 * time.Second)
	fmt.Println(time.Now())
	for o := 0; o < workerCount; o++ {
		go func(o int) {
			out := <-intCh
			fmt.Printf("%s出%d拿到%d\n", time.Now(), o, out)
		}(o)
	}
	time.Sleep(1 * time.Second)
}

//让out部分等待一段时间再取数据，来观察i的行为
//buffer数量较少，chan填满之后in会进入等待，直到out出现
func TestChanPutGet2_oWait_withSmallBuffer(t *testing.T) {
	intCh := make(chan int, 1)
	workerCount := 10
	for i := 0; i < workerCount; i++ {
		go func(i int) {
			fmt.Println(i, "开始工作", time.Now())
			intCh <- i
			fmt.Println(i, "结束工作", time.Now())
		}(i)
	}
	fmt.Println(time.Now())
	time.Sleep(3 * time.Second)
	fmt.Println(time.Now())
	for o := 0; o < workerCount; o++ {
		go func(o int) {
			out := <-intCh
			fmt.Printf("%s出%d拿到%d\n", time.Now(), o, out)
		}(o)
	}
	time.Sleep(1 * time.Second)
}

//out先启动，in后启动
func TestChanPutGet2_iWait_withSmallBuffer(t *testing.T) {
	intCh := make(chan int, 10) //创建一个不带size的channel(不带buffer)
	workerCount := 10
	for o := 0; o < workerCount; o++ {
		go func(o int) {
			out := <-intCh
			fmt.Printf("%s出%d拿到%d\n", time.Now(), o, out)
		}(o)
	}
	fmt.Println(time.Now())
	time.Sleep(3 * time.Second)
	fmt.Println(time.Now())
	for i := 0; i < workerCount; i++ {
		go func(i int) {
			fmt.Println(i, "开始工作", time.Now())
			intCh <- i
			fmt.Println(i, "结束工作", time.Now())
		}(i)
	}
	time.Sleep(1 * time.Second)
}

//range完所有数据之后报错
func TestRangeChannel(t *testing.T) {
	intCh := make(chan int, 10)
	for i := 0; i < 10; i++ {
		intCh <- i
	}
	for ch := range intCh {
		fmt.Println(ch)
	}
}

//先关闭chan，range不再报错
func TestRangeChannelClose(t *testing.T) {
	intCh := make(chan int, 10)
	for i := 0; i < 10; i++ {
		intCh <- i
	}
	close(intCh)
	for ch := range intCh {
		fmt.Println(ch)
	}
}

//先取出部分数据之后遍历chan，会将chan剩余的数据取出。先遍历chan，再取数据会取到假数据，需要加ok判断
func TestRangeChannelFakeData(t *testing.T) {
	intCh := make(chan int, 10)
	for i := 0; i < 10; i++ {
		intCh <- i
	}
	close(intCh)
	o1, ok := <-intCh
	fmt.Println(o1, ok)
	for ch := range intCh {
		fmt.Println(ch)
	}
	o2, ok := <-intCh
	fmt.Println(o2, ok)
}

//select的时候，哪个chan准备好了就选择哪个，如果由多个都准备好了就随机选一个，如果都没准备好就等待，如果都没准备好的时候还有default就执行default
func TestSelectChannel(t *testing.T) {
	ch1 := make(chan int, 2)
	ch2 := make(chan string)
	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- 1
	}()
	ch1 <- 1
	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "a"
	}()

	select {
	case o1 := <-ch1:
		fmt.Println("ch1 is ready, go", o1)
	case o2 := <-ch2:
		fmt.Println("ch2 is ready, go", o2)
	default:
		fmt.Println("no one ready, go")
	}
}

//关闭的chan总是准备好的
func TestSelectCloseChannel(t *testing.T) {
	ch1 := make(chan int)
	ch2 := make(chan string)
	//go func() {
	//	time.Sleep(1 * time.Second)
	//	ch1 <- 1
	//}()
	//ch1 <- 1
	close(ch1)
	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "a"
	}()

	select {
	case o1 := <-ch1:
		fmt.Println("ch1 is ready, go", o1)
	case o2 := <-ch2:
		fmt.Println("ch2 is ready, go", o2)
	default:
		fmt.Println("no one ready, go")
	}
}

func TestMultipleSelect(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(10)
	ch1 := make(chan int)
	for i := 0; i < 10; i++ {
		go func(i int) {
			defer wg.Done()
			select {
			case <-ch1:
				fmt.Println(time.Now(), i)
			}
		}(i)
	}
	fmt.Println("关闭channel", time.Now())
	//wg.Wait()//在这里等待会阻塞主协程导致死锁报错
	close(ch1)
	wg.Wait()
}

func TestDualCloseChannel(t *testing.T) {
	c := make(chan int, 5) //没有size的chan会被阻塞
	c <- 1
	close(c) //关闭的chan不可传入数据
	<-c
	//close(c)//关闭的chan不可再次关闭
}

//单线程谨慎使用chan
func TestNoRoutine(t *testing.T) {
	intCh := make(chan int)
	intCh <- 1 //直接报错，阻塞
	<-intCh    //直接报错，阻塞
}

func TestMultipleChannelReadySelect(t *testing.T) {
	ch1, ch2 := make(chan int), make(chan int)
	close(ch1)
	close(ch2)
	ch1Count, ch2Count := 0, 0
	for i := 0; i < 10000; i++ {
		select {
		case <-ch1:
			ch1Count++
		case <-ch2:
			ch2Count++
		}
	}
	fmt.Println("ch1", ch1Count, "ch2", ch2Count)
}

func TestMultipleChannelCaseInDefault(t *testing.T) { //优化channel无优先级问题，ch1的优先级更高，但是带来新的问题是ch1没有ready的时候会优先ch2
	ch1, ch2 := make(chan int), make(chan int)
	close(ch1)
	close(ch2)
	ch1Count, ch2Count := 0, 0
	for i := 0; i < 10000; i++ {
		select {
		case <-ch1:
			ch1Count++
		default:
			select {
			case <-ch2:
				ch2Count++
			}
		}
	}
	fmt.Println("ch1", ch1Count, "ch2", ch2Count)
}
