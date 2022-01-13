package main

import "log"

type fatRateService struct {
	s *FatRateSuggestion
}

func (f fatRateService) GiveSuggestionToPerson(person *Person) string {
	if err := person.calcBmi(); err != nil {
		log.Printf("无法给%s计算BMI:%v", person.name, err)
	}
	person.calcFatRate()
	suggestion, err := f.s.GetSuggestion(person)
	if err != nil {
		log.Printf("无法给%s计算FatRate:%v", person.name, err)
	}
	return suggestion
}

func (f fatRateService) GiveSuggestionToPersons(persons ...*Person) (outPut map[*Person]string) {
	for _, person := range persons {
		outPut[person] = f.GiveSuggestionToPerson(person)
	}
	return
}
