package main

import (
	"fmt"
	"testing"
)

func TestAssertion(t *testing.T) {
	//r := Refrigerator{size: "very big"}
	//var b Box = r
	//var c Close = b
	r := TestBox{}
	//var tb TB = r
	//var c Close = tb
	var c Close = r

	switch cDetail := c.(type) {
	case Refrigerator:
		fmt.Println(cDetail.size)
	case Box:
		fmt.Println("这是一个Box，不是冰箱")
	case TestBox:
		fmt.Println("这是一个TestBox，不是冰箱")
	}

	result, ok := checkRefrigerator(c)
	if ok {
		fmt.Println("是个冰箱", result)
	} else {
		fmt.Println("不是冰箱", result)
	}

	//r2 := c.(Refrigerator)
	//r3 := b.(c)//c未实现b，不可转
	//fmt.Println(r2.size)
}

func checkRefrigerator(c Close) (Refrigerator, bool) {
	r, ok := c.(Refrigerator)
	return r, ok
}

type TB interface {
	Close
}

type TestBox struct {
}

func (tb TestBox) Close() error {
	return nil
}
