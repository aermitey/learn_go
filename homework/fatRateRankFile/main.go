package main

import (
	"fmt"
)

func main() {
	record := &Record{
		FilePath: "/Users/chenxi/go/src/learngo/chapter08/testFatRate",
	}
	input := &InputFromStd{}
	//rankItem := &RankItem{}
	rank := &FatRateRank{}
	for i := 0; i < 3; i++ {
		personInfo := input.fakeRegister()
		record.SaveInformation(personInfo)
		fatRate, err := personInfo.CalcFatRate()
		if err != nil {
			return
		}
		rank.InputRecord(personInfo.Name, fatRate)
		personRank, fatRate := rank.GetRank(personInfo.Name)
		fmt.Println(personRank, fatRate)
	}
	fmt.Println("修改")
	personInfo := input.Update()
	fatRate, err := record.UpdateInfo(personInfo.Name, personInfo.Tall, personInfo.Weight, personInfo.Age)
	if err != nil {
		return
	}
	for i, _ := range *rank {
		if (*rank)[i].Name == personInfo.Name {
			(*rank)[i].FatRate = fatRate
		}
	}
	personRank, fatRate := rank.GetRankWithQuickSort(personInfo.Name)
	fmt.Println(personRank, fatRate)
}
