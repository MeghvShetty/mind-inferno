package main

import "fmt"

type Carloan struct {
	Amount    float64
	ModelYear int
	ModelType string
}

func (c Carloan) loanCal() float64 {
	if c.ModelYear-2026 == 0 {
		return c.Amount / 100 * 6
	} else {
		return c.Amount / 100 * 4
	}
}

func main() {
	c := Carloan{500, 200, "suv"}
	fmt.Println(c.loanCal())
}
