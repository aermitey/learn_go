package main

import "fmt"

func main() {
	a, b := 1, 2
	c := Calculator{
		A: a,
		B: b,
	}
	fmt.Println(c.add(), c.C)
	newC := &NewCalculator{&Calculator{}}
	newC.Calculator.A = 12
	newC.Calculator.B = 20

	m := &MyCommand{commandOptions: map[string]string{}}
	m.commandOptions["aa"] = "AAA"
	m.abc = 2
	m.ToCmdStr(3)
	fmt.Println(m)
}

type MyCommand struct {
	commandOptions map[string]string
	abc            int
}

func (m *MyCommand) ToCmdStr(l int) string {
	out := ""
	for s, s2 := range m.commandOptions {
		out = out + fmt.Sprintf("--%s=%s+%d", s, s2, l)
	}
	m.abc = l
	return out
}
