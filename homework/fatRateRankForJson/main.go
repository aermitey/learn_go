package main

import (
	"fmt"
	"log"
)

func main() {
	record := &Record{
		FilePath: "/Users/chenxi/go/src/learngo/chapter08/testFatRate",
	}
	input := &InputFromStd{}
	//rankItem := &RankItem{}
	rank := &FatRateRank{}
	for i := 0; i < 3; i++ {
		personInfo := input.Register()
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
	err := record.UpdateInfo(personInfo.Name, personInfo.Tall, personInfo.Weight, personInfo.Age)
	if err != nil {
		return
	}
	personInfos, err := record.Unmarshal(err)
	if err != nil {
		return
	}
	for i, _ := range *personInfos {
		fatRate, err := (*personInfos)[i].CalcFatRate()
		if err != nil {
			log.Println(err)
			return
		}
		rank.InputRecord((*personInfos)[i].Name, fatRate)
	}
	personRank, fatRate := rank.GetRankWithQuickSort(personInfo.Name)
	fmt.Println(personRank, fatRate)
	//fatRate, err := personInfo.CalcFatRate()
	//if err != nil {
	//	return
	//}
	//for _, item := range *rank {
	//	if item.Name == personInfo.Name {
	//		item.FatRate = fatRate
	//	}
	//}
	//personRank, fatRate := rank.GetRank(personInfo.Name)
	//fmt.Println(personRank, fatRate)
}
