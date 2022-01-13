package main

type elevatorOperation interface {
	add(...int)
	arrive(int) error
	up() error
	down() error
}
