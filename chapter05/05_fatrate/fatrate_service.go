package main

import "log"

type fatRateService struct {
	input  InputService
	s      *FatRateSuggestion
	output OutputService
}

func (f *fatRateService) GiveSuggestionToPerson(person *Person) {
	if err := person.calcBmi(); err != nil {
		log.Printf("无法给%s计算BMI:%v", person.name, err)
	}
	person.calcFatRate()
	suggestion, err := f.s.GetSuggestion(person)
	if err != nil {
		log.Printf("无法给%s计算FatRate:%v", person.name, err)
	}
	f.output.OutPut(*person, suggestion)
}

//func (f *fatRateService) GiveSuggestionToPersons(persons ...*Person) (outPut map[*Person]string) {
//	for _, person := range persons {
//		outPut[person] = f.GiveSuggestionToPerson(person)
//	}
//	return
//}
