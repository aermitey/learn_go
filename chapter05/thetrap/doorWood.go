package main

import "fmt"

var _ Door = &WoodDoor{}

type WoodDoor struct {
}

func (g *WoodDoor) Unlock() {
	fmt.Println("WoodDoor UnLock")
}

func (g *WoodDoor) Lock() {
	fmt.Println("WoodDoor Lock")
}

func (g *WoodDoor) Open() {
	fmt.Println("WoodDoor Open")
}

func (g *WoodDoor) Close() {
	fmt.Println("WoodDoor Close")
}
