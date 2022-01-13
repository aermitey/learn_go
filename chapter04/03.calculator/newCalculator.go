package main

type NewCalculator struct {
	*Calculator //如果嵌套的是指针，则在实例化的时候必须嵌套实例化的实体，否则会空指针。
}

