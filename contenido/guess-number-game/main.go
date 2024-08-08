package main

import (
	"fmt"
	"math/rand"
)

func main() {
	game()
}

func game() {

	randomNumber := rand.Intn(100)

	var numeroIngresado, intentos int
	var maxIntentos = 10

	for intentos < maxIntentos {
		intentos++
		fmt.Printf("Ingrese un numero (intentos restantes %d): ", maxIntentos-intentos+1)
		fmt.Scanln(&numeroIngresado)

		if numeroIngresado == randomNumber {
			fmt.Println("Felicidades haz acertado el numero ")
			playAgain()
			return
		} else if numeroIngresado > randomNumber {
			fmt.Printf("El numero aleatorio es mayor \n")
		} else {
			fmt.Printf("El numero aleatorio es menor \n")
		}

	}

	fmt.Printf("Se acabo los intentos y el numero era %d ", randomNumber)
	playAgain()
}

func playAgain() {
	var option string

	fmt.Println("Ingrese s/n si desea jugar nuevamente")
	fmt.Scanln(option)

	switch option {
	case "s":
		game()
	case "n":
		fmt.Println("Gracias por jugar!")
	default:
		fmt.Println("Haz seleccionado una opcion incorrecta")
		playAgain()
	}

}
