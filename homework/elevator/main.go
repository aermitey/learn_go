package main

import (
	"fmt"
	"time"
)

var s elevator

func main() {
	e := &elevator{}
	s := service{
		ele: e,
	}
	s.ele.add(10, 15, 16, 7, 6, 3, 1)
	go func() {
		time.Sleep(5 * time.Second)
		s.ele.add(5, 13, 2)
	}()
	fmt.Println(s.ElevatorOperation(e))
	//s := []string{"1", "2"}
	//for i, v := range s {
	//	fmt.Printf("%p,%p\n", &s[i], &v)
	//}
	//os.Exit(1)
}

//type A frInterface {
//	Push() error
//}
//
//type a struct{}
