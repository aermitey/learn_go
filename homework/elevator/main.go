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
	s.ele.add(10, 10, 10, 7, 6, 3, 1)
	fmt.Println(s.ElevatorOperation(e))
}
