package main

import "fmt"

type StdOut struct {
}

func (st StdOut) OutPut(person Person, s string) {
	fmt.Println(s)
}
