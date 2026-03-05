package main

import "fmt"

func main() {
	i, j := 42, 2701
	p := &i
	fmt.Println("read i through the pointer: ", *p)
	*p = 21
	fmt.Println(i) //set i though the pointer

	p = &j
	*p = *p / 37
	fmt.Println(j)
}
