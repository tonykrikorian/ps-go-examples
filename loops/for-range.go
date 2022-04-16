package main

import "fmt"

func main() {
	coursesInProgress := []string{
		"Docker",
		"Kubernetes",
	}

	for _, i := range coursesInProgress {

		fmt.Println("index: ",i)
	
	}
}