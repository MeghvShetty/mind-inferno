package main

import "fmt"

func main() {
	var s []int
	printSlice(s)

	//append works on nill slices.
	s = append(s, 9)
	printSlice(s)

	// The slice grows as needed.
	s = append(s, 2)
	printSlice(s)

	// we can add more than one element at a time.
	s = append(s, 3, 4, 5, 6)
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d , cap=%d %v\n", len(s), cap(s), s)
}
