package main

import "fmt"

/*
Las variables declaradas fuera de una función, tienen scope
a nivel de paquete
*/
var(
   name,course string
   module, clip int
)


func main() {
	fmt.Println("Name and course are set to", name, "and",course)
	fmt.Println("module and clip are set to", module, "and",clip)
}
