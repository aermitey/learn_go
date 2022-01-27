package main

import "fmt"

type shipLegend struct {
}

func (s *shipLegend) OpenTheDoor(ref Refrigerator) error {
	fmt.Println("使用ship装大象")
	return nil
}

func (s *shipLegend) PutEleInto(ele Elephant, ref Refrigerator) error {
	fmt.Println("使用ship装大象")
	return nil
}

func (s *shipLegend) CloseTheDoor(ref Refrigerator) error {
	fmt.Println("使用ship装大象")
	return nil
}
