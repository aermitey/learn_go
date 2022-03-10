package main

import (
	"fmt"
	"sync"
)

type Queue struct {
	sync.Mutex
	data []interface{}
}

func (q *Queue) Push(data ...interface{}) {
	q.Lock()
	defer q.Unlock()
	q.data = append(q.data, data...)
}

func (q *Queue) Pop() (interface{}, bool) {
	q.Lock()
	defer q.Unlock()
	if len(q.data) > 0 {
		o := q.data[0]
		q.data = q.data[1:]
		return o, true
	} else {
		return nil, false
	}
}

func main() {
	queue := &Queue{
		data: []interface{}{},
	}
	queue.Push(1, 2, 3)
	fmt.Println(queue)
	fmt.Println(queue.Pop())
	fmt.Println(queue.Pop())
	fmt.Println(queue.Pop())
	fmt.Println(queue.Pop())
}
