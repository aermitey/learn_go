package main

type fakeInput struct {
}

func (f fakeInput) GetInPut() Person {
	return Person{
		name:   "小强",
		sex:    "男",
		tall:   1.7,
		weight: 70,
		age:    35,
	}
}

