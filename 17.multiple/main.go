package main

import "fmt"

func main() {
	//bmis := []float64{1,2,3,4,5}
	avgBmi := calculateAvg()
	a := make([]float64, 3, 3)
	avgBmi1 := calculateAvgOnSlice(a)
	fmt.Println(avgBmi)
	fmt.Println(avgBmi1)
}

func calculateAvg(bmis ...float64) (avgBmi float64) {
	var total float64 = 0
	for _, item := range bmis {
		total += item
	}
	avgBmi = total / float64(len(bmis))
	return avgBmi
}

func calculateAvgOnSlice(bmis []float64) float64 {
	var total float64 = 0
	for _, item := range bmis {
		total += item
	}
	return total / float64(len(bmis))
}

func getScoresOfStudent(stdName string) (chiness int, math int, english int) {
	chiness = 100
	math = 120
	english = 110
	return chiness, math, english
}
