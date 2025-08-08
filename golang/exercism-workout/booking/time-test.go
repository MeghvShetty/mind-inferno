package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	date := "7/25/2019 13:45:00"
	time, err := time.Parse("1/02/2006 15:04:05", date)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(time)
}
