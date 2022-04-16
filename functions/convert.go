package main

import (
	"fmt"
	"strings"
)
func main() {

	author := "tony krikorian"
	course := "getting started with kubernetes"

	fmt.Println(converter(author,course))
}

func converter(author string ,course string) (a string, c string){
	
  author = strings.ToUpper(author)
  course = strings.Title(course)
	
	return author,course
}