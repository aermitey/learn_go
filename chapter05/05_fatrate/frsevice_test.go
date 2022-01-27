package main

import "testing"

func TestFrServiceSuggestion(t *testing.T) {
	realOutput := &fakeOutput{} //需要定义一个realOutput，不然下面s会取不到值
	frSvc := &fatRateService{
		input:  fakeInput{},
		s:      GetFatRateSuggestions(),
		output: realOutput,
	}
	p := frSvc.input.GetInPut()
	expectedSuggestion := &fakeOutput{
		p: p,
		s: "偏重",
	}
	frSvc.GiveSuggestionToPerson(&p)
	suggestion := realOutput.s
	if expectedSuggestion.s != suggestion {
		t.Errorf("预期：%s，实际：%s", expectedSuggestion.s, suggestion)
	}
}
