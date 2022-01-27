package main

import "fmt"

type manLegend struct {
}

func (s *manLegend) OpenTheDoor(ref Refrigerator) error {
	fmt.Println("使用man装大象")
	return nil
}

func (s *manLegend) PutEleInto(ele Elephant, ref Refrigerator) error {
	fmt.Println("使用man装大象")
	return nil
}

func (s *manLegend) CloseTheDoor(ref Refrigerator) error {
	fmt.Println("使用man装大象")
	return nil
}
