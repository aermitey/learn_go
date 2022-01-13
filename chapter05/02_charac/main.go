package main

import (
	"fmt"
	"reflect"
)

func main() {
	//var ip *int//默认为nil
	//fmt.Println(ip)//Panic，空指针异常
	//m := map[string]string{} //map能初始化实例，但是接口不行
	//var putER PutEleIntoRef //默认为nil
	//putER.OpenTheDoor(ref)  //空指针异常
	var ref Refrigerator
	var ele Elephant
	var t Test
	var legend PutEleIntoRef = t
	var leg PutEleIntoRef = ref //如果某个对象的成员方法有的在对象上有的在指针上，那么给接口赋值时一定要用指针
	legend.OpenTheDoor(ref)
	legend.PutEleInto(ele, ref)
	leg.CloseTheDoor(ref)
	var open Open = Refrigerator{}
	var close Close = Refrigerator{}
	fmt.Println(open, close)
	var box Box = Refrigerator{}
	fmt.Println(box)
	close = box
	//box = close//无法将 'close' (类型 Close) 用作类型 Box 类型未实现 'Box'，因为缺少某些方法: Open() error
	var i interface{}
	i = 3
	i = 3.14
	i = Refrigerator{}
	fmt.Println(reflect.TypeOf(i), reflect.ValueOf(i))
}

type Test func()

func (t Test) OpenTheDoor(ref Refrigerator) error {
	fmt.Println("关门冰箱门")
	return nil
}

func (t Test) PutEleInto(ele Elephant, ref Refrigerator) error {
	fmt.Println("关门冰箱门")
	return nil
}

func (t Test) CloseTheDoor(ref Refrigerator) error {
	fmt.Println("关门冰箱门")
	return nil
}

type PutEleIntoRef interface {
	OpenTheDoor(ref Refrigerator) error
	PutEleInto(ele Elephant, ref Refrigerator) error
	CloseTheDoor(ref Refrigerator) error
}

type PutEleIntoRefImpl struct {
}

func (p *PutEleIntoRefImpl) OpenTheDoor(ref Refrigerator) error {
	fmt.Println("打开冰箱门")
	return nil
}

func (p PutEleIntoRefImpl) PutEleInto(ele Elephant, ref Refrigerator) error {
	fmt.Println("大象装进去")
	return nil
}

func (p *PutEleIntoRefImpl) CloseTheDoor(ref Refrigerator) error {
	fmt.Println("关门冰箱门")
	return nil
}

type Refrigerator struct {
	size string
}

func (r Refrigerator) OpenTheDoor(ref Refrigerator) error {
	fmt.Println("打开冰箱门")
	return nil
}

func (r Refrigerator) PutEleInto(ele Elephant, ref Refrigerator) error {
	fmt.Println("打开冰箱门")
	return nil
}

func (r Refrigerator) CloseTheDoor(ref Refrigerator) error {
	fmt.Println("打开冰箱门")
	return nil
}

type Elephant struct {
	name string
}

type Open interface {
	Open() error
}

type Close interface {
	Close() error
}

type Box interface {
	Open
	Close
}

func (r Refrigerator) Open() error {
	return nil
}

func (r Refrigerator) Close() error {
	return nil
}
