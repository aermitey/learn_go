package main

import "testing"

func TestFatRateSuggestion_Suggestion(t *testing.T) {
	inputPerson := &Person{}
	inputPerson.sex = "女"
	inputPerson.age = 23
	inputPerson.fatRate = 0.2
	expectedSuggestion := "便瘦"
	t.Logf("开始计算，输入：sex:%s, age:%d,fatRate:%f, 期望结果：%s", inputPerson.sex, inputPerson.age, inputPerson.fatRate, expectedSuggestion)
	actualSuggestion, err := FatRateSuggestion{}.GetSuggestion(inputPerson)
	if err != nil {
		t.Fatalf("报错，错误信息为%v", err)
	}
	if actualSuggestion != expectedSuggestion {
		t.Errorf("预期为%s，得到%s", expectedSuggestion, actualSuggestion)
	}
}
