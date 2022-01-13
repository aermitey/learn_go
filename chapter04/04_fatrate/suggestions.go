package main

func GetFatRateSuggestions() *FatRateSuggestion {
	return &FatRateSuggestion{
		suggestionArr: [][][]int{
			{ //第一个元素，表示男
				{ //18-39
					0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 2, 2, 2, 2, 2, 3, 3, 3, 3, 3, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4,
				},
				{ //40-59

				},
				{ //60-150

				},
			},
			{ //第二个元素，表示女
				{ //18-39
					0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 2, 2, 2, 2, 2, 3, 3, 3, 3, 3, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4,
				},
				{ //40-59

				},
				{ //60-150

				},
			},
		},
	}
}

type FatRateSuggestion struct {
	suggestionArr [][][]int
}

func (s *FatRateSuggestion) GetSuggestion(person *Person) (suggestion string, err error) {
	sexIndex := s.getIndexOfSex(person.sex)
	ageIndex := s.getIndexOfAge(person.age)
	maxFRSupported := len(s.suggestionArr[sexIndex][ageIndex]) - 1
	FrIndex := int(person.fatRate * 100)
	if FrIndex > maxFRSupported {
		FrIndex = maxFRSupported
	}
	suggestionIndex := s.suggestionArr[sexIndex][ageIndex][int(person.fatRate*100)]
	suggestion = s.result(suggestionIndex)
	return
}

func (s *FatRateSuggestion) getIndexOfSex(sex string) int {
	if sex == "男" {
		return 0
	}
	return 1
}

func (s *FatRateSuggestion) getIndexOfAge(age int) int {
	switch {
	case age >= 18 && age <= 39:
		return 0
	case age >= 40 && age <= 59:
		return 1
	case age >= 60 && age <= 150:
		return 2
	default:
		return -1
	}
}

func (s *FatRateSuggestion) result(index int) string {
	switch index {
	case 0:
		return "偏瘦"
	case 1:
		return "正常"
	case 2:
		return "偏重"
	case 3:
		return "肥胖"
	case 4:
		return "非常肥胖"
	}
	return ""
}
