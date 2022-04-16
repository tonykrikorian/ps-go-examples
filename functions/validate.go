package main

import (
	"fmt"
	"os"
)

func main(){
_, err := os.Open("./test.txt")

if err != nil{
	fmt.Println("This is an error code", err)
}
}