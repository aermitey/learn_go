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

func (f *FatRateRank) InputRecord(name string, fatRate ...float64) {
	minFatRate := math.MaxFloat64
	for _, item := range fatRate {
		if minFatRate > item {
			minFatRate = item
		}
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

func (f *FatRateRank) GetRank(name string) (rank int, fatRate float64) {
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

func (f *FatRateRank) GetRankWithBubble(name string) (rank int, fatRate float64) {
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
	bubble(&rankArr)
	for i, frItem := range rankArr {
		if frItem == fatRate {
			rank = i + 1
			break
		}
	}
	return
}

func (f *FatRateRank) GetRankWithQuickSort(name string) (rank int, fatRate float64) {
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
	quickSort(&rankArr, 0, len(rankArr)-1)
	for i, frItem := range rankArr {
		if frItem == fatRate {
			rank = i + 1
			break
		}
	}
	return
}

func bubble(arr *[]float64) {
	for i := 0; i < len(*arr); i++ {
		for j := 0; j < len(*arr)-i-1; j++ {
			if (*arr)[j] > (*arr)[j+1] {
				(*arr)[j], (*arr)[j+1] = (*arr)[j+1], (*arr)[j]
			}
		}
	}
}

func quickSort(arr *[]float64, start, end int) {
	pivotIdx := (start + end) / 2
	pivotValue := (*arr)[pivotIdx]
	l, r := start, end
	for l < r {
		for (*arr)[l] < pivotValue {
			l++
		}
		for (*arr)[r] > pivotValue {
			r--
		}
		if l >= r {
			break
		}
		(*arr)[l], (*arr)[r] = (*arr)[r], (*arr)[l]
		l++
		r--
	}
	if l == r {
		l++
		r--
	}
	if r > start {
		quickSort(arr, start, r)
	}
	if l < end {
		quickSort(arr, l, end)
	}
}
