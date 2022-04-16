package main

import "fmt"


func main() {
	type personData struct {
		name     string
		lastname string
		email    string
	}

	employee := personData{
		name:     "Tony",
		lastname: "Krikorian",
		email:    "tony.krikorian@outlook.com",
	}

fmt.Print(employee)
}