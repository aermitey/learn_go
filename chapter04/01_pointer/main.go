package main

import "fmt"

func main() {
	a := 1
	b := 2
	c := &[]int{1, 2, 3, 4, 5}
	e := &c
	d := *c
	fmt.Println(d[1])
	fmt.Println((*c)[1])
	fmt.Println("e=", e, "*e=", *e, "**e=", **e)
	type o struct {
		m int
		n int
	}
	x := &o{
		m: 0,
		n: 0,
	}
	y := *x
	fmt.Println("y=", y.n)
	fmt.Println("x=", x.n)
	type oo struct {
		p *o
	}
	xx := &oo{
		p: x,
	}

	fmt.Println("xx=", &xx)
	adder := add(&a, &b)
	fmt.Println(adder)
	j := new(o)
	fmt.Println(j)
	f1 := add
	f1(&a, &b)
	//fmt.Println(f1(&a, &b))
	f1p1 := &f1
	(*f1p1)(&a, &b)
	fmt.Println("(*f1p1)(&a,&b) = ", (*f1p1)(&a, &b))

	z := map[string]string{}
	zp1 := &z
	fmt.Println(zp1)
	put(z)
	fmt.Println("z=", z)
	fmt.Println("zp1=", zp1)
	put(*zp1)
	fmt.Println("zp1=", zp1)
	//var nothing *int
	//*nothing = 3 //这里是没有指向任何东西的int指针，会报错
	//fmt.Println(*nothing)
	//var nothingMap map[string]string
	//nmap := nothingMap
	//nmap["a"] = "A" //崩溃，必须实例化才可以装东西
	nothingMap := map[string]string{}
	nothingMap["a"] = "A"
	var nothingSlice []int
	//nothingSlice[0] = 100//崩溃，必须实例化
	nothingSlice = append(nothingSlice, 10) //可以append
}

func add(pa, pb *int) int {
	*pa = *pa + *pb
	return *pa
}

func put(m map[string]string) {
	m["a"] = "A"
}
