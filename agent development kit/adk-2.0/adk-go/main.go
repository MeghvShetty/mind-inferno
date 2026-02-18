package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("This is a Megh Bot...")
	resp, err := http.Get("http://localhost:11434/api/chat")
	if err != nil {
		fmt.Println("Error")
	}
	fmt.Println(resp.Status)

}
