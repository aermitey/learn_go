package main

type fakeOutput struct {
	p Person
	s string
}

func (f *fakeOutput) OutPut(person Person, s string) {
	f.p = person
	f.s = s
}
