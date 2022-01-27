package main

type InputService interface {
	GetInPut() Person
}

type OutputService interface {
	OutPut(Person, string)
}
