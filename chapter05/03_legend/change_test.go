package main

import (
	"fmt"
	"testing"
)

type Change interface {
	ChangeName(newName string)
	ChangeAge(newAge int)
}

type Student struct {
	Name string
	Age  int
}

func (s *Student) ChangeName(newName string) {
	s.Name = newName
}

func (s *Student) ChangeAge(newAge int) {
	s.Age = newAge
}

func TestVal(t *testing.T) {
	var stdChg Change
	stdChg = &Student{} //当实现接口的方法在对象指针上，只能用指针作为值赋值给接口。对象上则都可。如果既有对象又有指针则只能用指针
	fmt.Println(stdChg)
}

//func TestStdName(t *testing.T) {
//	s := Student{
//		Name: "tom",
//	}
//	fmt.Println(s)
//	s.ChangeName("jerry")
//	fmt.Println(s)
//}
