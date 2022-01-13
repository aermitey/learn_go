package main

import "fmt"

func main() {
	e := &elevator{
		floor:  10,
		status: 0,
	}
	s := service{
		ele: e,
	}
	s.ele.add(10, 3, 7, 8, 5, 16, 15, 19)
	go func() {
		s.ele.add(12, 17, 18)
	}()
	fmt.Println(s.ElevatorOperation(e))
}
