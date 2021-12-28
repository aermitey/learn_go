package main

import (
	"math"
	"sort"
)

var (
	personFatRate = map[string]float64{}
)

func inputRecord(name string, fatRate ...float64) {
	minFatRate := math.MaxFloat64
	for _, item := range fatRate {
		minFatRate = item
	}
	personFatRate[name] = minFatRate
}

func getRank(name string) (rank int, fatRate float64) {
	fatRatePerson := map[float64][]string{}
	rankArr := make([]float64, 0, len(personFatRate))
	for nameItem, frItem := range personFatRate {
		fatRatePerson[frItem] = append(fatRatePerson[frItem], nameItem)
		rankArr = append(rankArr, frItem)
	}
	sort.Float64s(rankArr)
	for i, f := range rankArr {
		_names := fatRatePerson[f]
		for _, _name := range _names {
			if _name == name {
				rank = i + 1
				fatRate = f
				return
			}
		}
	}
	return
}
