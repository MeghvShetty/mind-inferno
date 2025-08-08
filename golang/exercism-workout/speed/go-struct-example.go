package main

import "fmt"

type Engineer struct {
	Name    string
	Age     int
	Project Project //nested struct
}

type Project struct {
	Name         string
	Priority     string
	Technologies []string //passing a slice
}

func main() {
	fmt.Println("Structs Tutorial")
	// the verbos approach is good for production for readablity
	engineer := Engineer{
		Name: "Emeg",
		Age:  28,
		Project: Project{
			Name:         "Guide to learning go",
			Priority:     "High",
			Technologies: []string{"Go", "VS-Code"},
		},
	}
	// Can also be used in this way
	//engineer := Engineer{"Emeg"}

	//Common print statment
	fmt.Println(engineer)

	// Print a given field in struct
	fmt.Println(engineer.Project.Name)
	// Gives you the struct type when printing
	//fmt.Printf("%+v\n", engineer)
}
