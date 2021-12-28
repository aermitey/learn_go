package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	var hello string = "hello, golang"
	fmt.Println(reflect.TypeOf(hello))
	var world = "world"
	println(hello, world)
	var int3, int4 uint = 33, 44
	fmt.Println(int3 * int4)
	float := 1.234
	fmt.Println(float)

	var ho, ver = 3, 4.12
	var sc = ho * int(ver)
	var ss = float64(ho) * ver
	fmt.Println(ho * int(ver))
	fmt.Println(ss)
	fmt.Println(sc)

	var newname string
	fmt.Println("newname =", newname)

	var int6 uint = math.MaxUint64
	fmt.Println(int6)
	var int7 = int(int6)
	fmt.Println(int7)



}
