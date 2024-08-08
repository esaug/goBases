package main

import (
	"fmt"
	"math"
)

const decimalsPresicion = "%.2f"

func main() {

	var base, altura float64

	fmt.Println("Ingrese la longitud de la base del triangulo: ")
	fmt.Scanln(&base)

	fmt.Println("Ingrese la longitud de altura del triangulo: ")
	fmt.Scanln(&altura)

	hipotenusaCuadrada := math.Pow(base, 2) + math.Pow(altura, 2)
	hipotenusa := math.Sqrt(hipotenusaCuadrada)

	area := base * altura / 2

	perimetro := base + altura + hipotenusa

	presicionDecimalArea := fmt.Sprintf(decimalsPresicion, area)
	presicionDecimalPerimetro := fmt.Sprintf(decimalsPresicion, perimetro)

	fmt.Printf("Area: %v \n", presicionDecimalArea)
	fmt.Printf("Perimetro: %v \n", presicionDecimalPerimetro)
}
