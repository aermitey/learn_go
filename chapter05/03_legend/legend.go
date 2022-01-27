package main

import (
	"fmt"
	"reflect"
)

func legendary(legend PutEleIntoRef, r Refrigerator, e Elephant) {
	fmt.Println("如何装大象")
	if value, ok := legend.(*manLegend); ok {
		fmt.Println("不建议使用人工", reflect.TypeOf(value))
	}

	legend.OpenTheDoor(r)
	legend.PutEleInto(e, r)
	legend.CloseTheDoor(r)
	fmt.Println("完成, legendary")
}
