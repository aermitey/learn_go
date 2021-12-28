package main

import "fmt"

func main() {
	lmap, rmap := map[string]int{}, map[string]int{}

	lmap["语文"] = 80
	rmap["数学"] = 70
	for s, i := range lmap {
		rmap[s] = i
	}
	fmt.Println(rmap)
}
