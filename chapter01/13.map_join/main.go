package main

import "fmt"

func main() {
	lMap, rMap := map[string]int{}, map[string]int{}

	lMap["语文"] = 80
	rMap["数学"] = 70
	for s, i := range lMap {
		rMap[s] = i
	}
	fmt.Println(rMap)
}
