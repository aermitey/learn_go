package main

import "fmt"

func main() {
	security := Assets{assets: []Asset{
		&GlassDoor{},
		&WoodDoor{},
		1,
	}}
	fmt.Println("开始上班")
	security.DoStartWork()
	fmt.Println("开始下班")
	security.DoStopWork()
	fmt.Println("DONE")
}
