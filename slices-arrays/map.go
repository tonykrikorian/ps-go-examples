package main

import "fmt"

func main() {

	months := map[string]int{
		"Enero":     1,
		"Febrero":   2,
		"Marzo":     3,
		"Abril":     4,
		"Mayo":      5,
		"Junio":     6,
		"Julio":     7,
		"Agosto":    8,
		"Setiembre": 9,
		"Octubre":   10,
		"Noviembre": 11,
		"Diciembre": 12,
	}

	// Para este caso no los devuelve en ningun orden especifico
	for mapKey, mapValue := range months {
	fmt.Printf("The key is %v and value is %v\n",mapKey,mapValue)	
	}

	//Para acceder a un valor
	fmt.Println(months["Marzo"])

	//Para actualizar el valor
	months["Enero"]=100

	//Para agregar un nuevo valor
	months["Febrero-Marzo"]=1000

	fmt.Println(months)

}