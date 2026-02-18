package main

import (
	"fmt"
	"strings"
)

func main() {
	val := strings.Fields("I ate a donut. Then I ate another donut.")
	m := make(map[string]int)
	fmt.Println(val)
	for _, v := range val {
		m[v]++
		fmt.Println(m)

	}

}
