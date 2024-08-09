package main

import (
	"fmt"
	"log"
	"os"
)

func validateZero(divisor int) {
	if divisor == 0 {
		panic("No se puede dividir entre cero")
	}
}

func division(dividir, divisor int) (int, error) {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()

	validateZero(divisor)

	return dividir / divisor, nil
}

func main() {

	result, err := division(10, 0)

	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	fmt.Println(result)

	// Defer example with file creation

	file, err := os.Create("hola.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	_, err = file.Write([]byte("Hola mundo"))
	if err != nil {
		fmt.Println(err)
		return
	}

	// Error registration
	log.SetPrefix("main()")
	file, err = os.OpenFile("info.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	log.SetOutput(file)
	log.Println("!Oye, soy un Log!")
}
