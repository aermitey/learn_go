package main

//
//import (
//	"fmt"
//	"math/rand"
//	"time"
//)
//
//var rank = &FatRateRank{}
//var rankChan = make(chan string, 1000)
//
//func main() {
//	Registered()
//	for {
//		for i := 0; i < 1000; i++ {
//			go func(i int) {
//				Update(fmt.Sprintf("%d", i))
//			}(i)
//		}
//		for i := 0; i < 1000; i++ {
//			go func() {
//				Rank()
//			}()
//		}
//	}
//}
//
//func randFloat() float64 {
//	rand.Seed(time.Now().UnixNano())
//	num := rand.Float64()
//	for {
//		if num > 0.4 {
//			num = num - 0.4
//		} else {
//			return num
//		}
//	}
//	return num
//}
//
//func Registered() {
//	for i := 0; i < 1000; i++ {
//		rank.inputRecord(fmt.Sprintf("%d", i))
//	}
//}
//
//func Update(name string) {
//	rank.inputRecord(name, randFloat())
//	rankChan <- name
//}
//
//func Rank() {
//	name := <-rankChan
//	fmt.Println(rank.getRank(name))
//}
