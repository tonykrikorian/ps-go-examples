package main

import "fmt"

func main() {
courses := []string{
	"ddd",
	"xxxx",
	"fdss",
	"uyu",
}

for _, course := range courses {
	
	fmt.Println(course)
	
}

fmt.Println(append(courses,"mmmm"))
}

