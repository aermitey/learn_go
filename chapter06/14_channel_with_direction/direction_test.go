package main

import "testing"

func TestChannelWithDirection(t *testing.T) {
	c := make(chan int, 100)
	inOnly(c)
	outOnly(c)
}

func inOnly(c chan<- int) {
	c <- 1
	//<-c//当c是单项入channel的时候不能取数，编译错误
}

func outOnly(c <-chan int) {
	//c <- 1//当c是单项出channel的时候不能存数，编译错误
	<-c
}
