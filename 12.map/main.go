package main

import "fmt"

func main() {
	names := []string{"小强", "小周", "小陈"}
	fr := []float64{28, 18, 17}
	names = append(names, "小松")
	fr = append(fr, 15)
	for i, name := range names {
		if name == "小强" {
			fmt.Printf("%s 的体脂率是 %f\n", name, fr[i])
		}
	}
	s1 := make([]int, 3, 3)
	fmt.Println(s1)
	var m1 map[string]int
	m2 := map[string][]int{}
	m3 := map[string]int{"小强": 60, "小周": 18, "小陈": 17}
	m3["小强"] = 20
	m3["小刘"] = 19
	fmt.Println(m3["小强"])
	for s, i := range m3 {
		fmt.Println(s, i)
	}
	xqScore, ok := m3["小强"]
	fmt.Println(xqScore, ok)
	delete(m3, "小强")
	xqScore, ok = m3["小强"]
	fmt.Println(xqScore, ok)
	fmt.Println(m3["小强"])
	m4 := make(map[string]int)
	fmt.Println(m1, m2, m3, m4)
}
