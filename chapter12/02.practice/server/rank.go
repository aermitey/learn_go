package main

import (
	"learngo/chapter02/15.fatrate_refactor/calc"
	"learngo/chapter12/02.practice/frInterface"
	"learngo/pkg/apis"
	"log"
	"math"
	"sort"
	"sync"
)

var _ frInterface.ServeInterface = &FatRateRank{}

type RankItem struct {
	Name    string
	Sex     string
	FatRate float64
}

type FatRateRank struct {
	items     []RankItem
	itemsLock sync.Mutex
}

func (f *FatRateRank) RegisterPersonalInformation(pi *apis.PersonalInformation) error {
	bmi, err := calc.CalcBMI(float64(pi.Weight), float64(pi.Tall))
	if err != nil {
		log.Println("计算BMI失败", err)
		return err
	}
	fr, err := calc.FatRateFromBMI(bmi, int(pi.Age), pi.Sex)
	if err != nil {
		log.Fatal("计算体脂率失败", err)
		return err
	}
	f.inputRecord(pi.Name, pi.Sex, fr)
	return nil
}

func (f *FatRateRank) UpdatePersonalInformation(pi *apis.PersonalInformation) (*apis.PersonalInformationFatRate, error) {
	bmi, err := calc.CalcBMI(float64(pi.Weight), float64(pi.Tall))
	if err != nil {
		log.Println("计算BMI失败", err)
		return nil, err
	}
	fr, err := calc.FatRateFromBMI(bmi, int(pi.Age), pi.Sex)
	if err != nil {
		log.Fatal("计算体脂率失败", err)
		return nil, err
	}
	f.inputRecord(pi.Name, pi.Sex, fr)
	return &apis.PersonalInformationFatRate{
		Name:    pi.Name,
		FatRate: fr,
	}, nil
}

func (f *FatRateRank) GetFatRate(name string) (*apis.PersonalRank, error) {
	rankId, sex, fr := f.getRank(name)
	return &apis.PersonalRank{
		Name:       name,
		Sex:        sex,
		RankNumber: rankId,
		FatRate:    fr,
	}, nil
}

func (f *FatRateRank) GetTop() ([]*apis.PersonalRank, error) {
	return f.getTopRank(), nil
}

func NewFatRateRank() *FatRateRank {
	return &FatRateRank{
		items: make([]RankItem, 0, 100),
	}
}

func (f *FatRateRank) inputRecord(name, sex string, fatRate ...float64) {
	f.itemsLock.Lock()
	defer f.itemsLock.Unlock()
	minFatRate := math.MaxFloat64
	for _, item := range fatRate {
		minFatRate = item
	}

	found := false
	for i, item := range f.items {
		if item.Name == name {
			if item.FatRate >= minFatRate {
				item.FatRate = minFatRate
			}
			f.items[i] = item
			found = true
			break
		}
	}
	if !found {
		f.items = append(f.items, RankItem{
			Name:    name,
			Sex:     sex,
			FatRate: minFatRate,
		})
	}
}

func (f *FatRateRank) getRank(name string) (rank int, sex string, fatRate float64) {
	f.itemsLock.Lock()
	defer f.itemsLock.Unlock()
	sort.Slice(f.items, func(i, j int) bool {
		return f.items[i].FatRate < f.items[j].FatRate
	})
	frs := map[float64]struct{}{}
	for _, item := range f.items {
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

func (f *FatRateRank) getTopRank() []*apis.PersonalRank {
	f.itemsLock.Lock()
	defer f.itemsLock.Unlock()
	sort.Slice(f.items, func(i, j int) bool {
		return f.items[i].FatRate < f.items[j].FatRate
	})
	out := make([]*apis.PersonalRank, 0, 10)
	for i := 0; i < 10 && i < len(f.items); i++ {
		out = append(out, &apis.PersonalRank{
			Name:       f.items[i].Name,
			Sex:        f.items[i].Sex,
			RankNumber: i,
			FatRate:    f.items[i].FatRate,
		})
	}
	return out
}
