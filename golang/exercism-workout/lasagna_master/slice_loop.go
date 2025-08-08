package main

import "fmt"

func main() {
	layers := []string{"sauce", "noodles", "sauce", "meat", "mozzarella", "noodles"}
	var dataCollection int
	for i := 0; i < len(layers); i++ {
		if layers[i] == "noodles" {
			dataCollection++
		}
	}
	fmt.Println(dataCollection)
}
