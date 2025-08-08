package main

import "fmt"

var birdsPerDay = []int{2, 5, 0, 7, 4, 1, 3, 0, 2, 5, 0, 1, 3, 1}

func main() {
	data := 0

	for i := 0; i < len(birdsPerDay); i++ {
		data += birdsPerDay[i]
	}
	fmt.Println(data)
}
