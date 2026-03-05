package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	z := 1.
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
		if z*z == x {
			return z
		}
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
