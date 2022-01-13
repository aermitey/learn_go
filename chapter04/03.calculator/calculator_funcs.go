package main

func (c *Calculator) add() int {
	c.C = c.A + c.B
	return c.C
}
func (c Calculator) sub() int {
	return c.add() - c.B
}
func (c Calculator) Multiple() int {
	return c.A * c.B
}
func (c Calculator) Divide() float64 {
	return float64(c.A / c.B)
}
func (c Calculator) Reminder() int {
	return c.A % c.B
}
