package main

import (
	"fmt"
	"sort"
)

type Button struct {
	Floor int
}

type Elevator struct {
	button   Buttons
	position int
}

type Buttons []Button

func (b Buttons) Len() int {
	return len(b)
}

func (b Buttons) Less(i, j int) bool {
	return b[i].Floor < b[j].Floor
}

func (b Buttons) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func main() {
	ele := &Elevator{
		position: 2,
		button: Buttons{
			{Floor: 3},
			{Floor: 1},
			{Floor: 5},
			{Floor: 2},
			{Floor: 4},
		},
	}
	sort.Sort(ele.button)
	sort.Sort(sort.Reverse(ele.button)) //大于
	fmt.Println(ele.button)
}
