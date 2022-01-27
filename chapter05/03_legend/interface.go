package main

type PutEleIntoRef interface {
	OpenTheDoor(ref Refrigerator) error
	PutEleInto(ele Elephant, ref Refrigerator) error
	CloseTheDoor(ref Refrigerator) error
}

type Refrigerator struct {
	size string
}

type Elephant struct {
	name string
}
