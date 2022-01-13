package main

import "fmt"

func main() {
	var data string
	{
		var equipment IOInterface = &Soft{}
		data = equipment.Read()
		fmt.Println(data)
	}
	{
		var equipment IOInterface = &Mag{}
		data = equipment.Read()
		fmt.Println(data)
	}
	{
		var equipment IOInterface = &Paper{}
		data = equipment.Read()
		fmt.Println(data)
	}
	{
		var equipment IOInterface = &Sata{}
		data = equipment.Read()
		fmt.Println(data)
	}

}

type IOInterface interface {
	Read() string
}

type Soft struct {
}

func (Soft) Read() string {
	return "软盘数据"
}

type Mag struct {
}

func (Mag) Read() string {
	return "磁带数据"
}

type Paper struct {
}

func (Paper) Read() string {
	return "纸带数据"
}

type Sata struct {
}

func (Sata) Read() string {
	return "固态数据"
}
