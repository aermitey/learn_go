package main

import "fmt"

var _ Door = &GlassDoor{}

type GlassDoor struct {
}

func (g *GlassDoor) Unlock() {
	fmt.Println("GlassDoor UnLock")
}

func (g *GlassDoor) Lock() {
	fmt.Println("GlassDoor Lock")
}

func (g *GlassDoor) Open() {
	fmt.Println("GlassDoor Open")
}

func (g *GlassDoor) Close() {
	fmt.Println("GlassDoor Close")
}
