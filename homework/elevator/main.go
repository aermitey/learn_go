package main

import (
	"fmt"
	"time"
)

func main() {
	e := &elevator{
		floor:  10,
		status: 0,
	}
	s := service{
		ele: e,
	}
	s.ele.add(10, 15, 16, 7, 6, 3, 1)
	go func() {
		time.Sleep(5 * time.Second)
		s.ele.add(5, 13, 2)
	}()
	fmt.Println(s.ElevatorOperation(e))
}
