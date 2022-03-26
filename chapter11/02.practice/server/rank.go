package main

import (
	"math"
	"sort"
)

type RankItem struct {
	Name    string
	FatRate float64
}

type FatRateRank []RankItem

func NewFatRateRank() *FatRateRank {
	return &FatRateRank{}
}

func (f *FatRateRank) inputRecord(name string, fatRate ...float64) {
	minFatRate := math.MaxFloat64
	for _, item := range fatRate {
		minFatRate = item
	}

	found := false
	for i, item := range *f {
		if item.Name == name {
			if item.FatRate >= minFatRate {
				item.FatRate = minFatRate
			}
			(*f)[i] = item
			found = true
			break
		}
	}
	if !found {
		*f = append(*f, RankItem{
			Name:    name,
			FatRate: minFatRate,
		})
	}
}

func (f *FatRateRank) getRank(name string) (rank int, fatRate float64) {
	sort.Slice(*f, func(i, j int) bool {
		return (*f)[i].FatRate < (*f)[j].FatRate
	})
	frs := map[float64]struct{}{}
	for _, item := range *f {
		frs[item.FatRate] = struct{}{}
		if item.Name == name {
			fatRate = item.FatRate
		}
	}
	rankArr := make([]float64, 0, len(frs))
	for k := range frs {
		rankArr = append(rankArr, k)
	}
	sort.Float64s(rankArr)
	for i, frItem := range rankArr {
		if frItem == fatRate {
			rank = i + 1
			break
		}
	}
	return
}
