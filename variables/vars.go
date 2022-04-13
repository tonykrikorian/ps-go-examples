package main

import (
	"fmt"
)



func main(){

	    name := "Tony"
		course:= "Starting with Go"

	fmt.Println("Hi ",name, " your course is ", course)
	updateCourseValue(&course)
	fmt.Println("You are currently watching",course)
}

func updateCourseValue(course *string) string{
	*course = "Starting with Kubernetes"
	fmt.Println("Updated course to",*course)
	return *course
}