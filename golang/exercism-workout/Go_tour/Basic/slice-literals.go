package main

import "fmt"

func main() {
	q := []int{1, 2, 3, 4, 5}
	fmt.Println(q)

	r := []bool{true, false, true, true, true, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, true},
		{4, false},
		{5, false},
	}
	fmt.Println(s)
}
