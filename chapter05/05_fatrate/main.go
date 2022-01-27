package main

func main() {

	frSvc := &fatRateService{
		input:  &inputFromStd{},
		s:      GetFatRateSuggestions(),
		output: StdOut{},
	}

	for {
		p := frSvc.input.GetInPut()
		frSvc.GiveSuggestionToPerson(&p)
	}
}
