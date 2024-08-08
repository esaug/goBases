package main

import (
	"fmt"
	"strconv"

	"rsc.io/quote"
)

// CONSTANTS declaration

const (
	Domingo = iota + 1
	Lunes
	Martes
	Miercoles
	Jueves
	Viernes
	Sabado
)

func main() {

	fmt.Println("Hola")
	fmt.Println(quote.Hello())

	// Variable declaration inside of a func, var is for outside and inside func declaration
	firstName, secondName, age := "Sau", "Guerra", 31

	fmt.Println(firstName, secondName, age)
	fmt.Println(Viernes, "to Dia de la semana")

	// DATA TYPES
	firstName = "Alojo mora \t(alias \"Rolex\" )"
	fmt.Println(firstName)

	s := "Hola"
	fmt.Println(s[0])

	var unicode rune = '♥'
	fmt.Println(unicode)

	// Variable without a initial value in the declaration they have a default value

	// Convert string to int
	one := "100"
	i, _ := strconv.Atoi(one)
	fmt.Println(i)

	// Number to string
	two := 200
	s = strconv.Itoa(two)
	fmt.Print(s)

	// SCAN LINE

	var nombre string
	var apellido string
	var edad int8

	fmt.Print("Ingrese su Nombre y Apellido")
	fmt.Scanln(&nombre, &apellido)

	fmt.Print("Ingrese la edad")
	fmt.Scanln(&edad)

	fmt.Printf("Usuario: %s %s edad: %d", nombre, apellido, edad)

}
