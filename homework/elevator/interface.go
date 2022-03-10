package main

type ElevatorOperation interface {
	add(...int)
	arrive(int) error
	up() error
	down() error
}
