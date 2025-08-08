package main

import "fmt"

type Rectangle struct {
	Lenght int
	width  int
}

func main() {
	p := fmt.Println
	var rect1 Rectangle
	rect1.width = 10
	rect1.Lenght = 12
	p(rect1)
}
