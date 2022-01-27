package main

import (
	"testing"
)

func TestFatRateSuggestion_GetSuggestion(t *testing.T) {
	suggestion := GetFatRateSuggestions()
	tests := []Person{
		{
			sex:     "男",
			age:     35,
			fatRate: 0.24,
		},
	}
	if gotSuggestion, _ := suggestion.GetSuggestion(&tests[0]); gotSuggestion != "肥胖" {
		t.Fail()
	}

}
