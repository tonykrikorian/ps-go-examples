package main

import "fmt"

func main(){
	  dddLengthMins := 275 //Docker Deep Dive Course
    cawLengthMins := 275 //Containers on AWS Wavelength course

	if dddLengthMins > cawLengthMins {
		fmt.Println("Docker course is Dive and dive than AWS")
	} else if dddLengthMins == cawLengthMins {
		fmt.Println("Both courses are equal")
	}else{
		fmt.Println("Ninguno de los dos cursos cuadra")
	}
}