package main

import "fmt"

// Function variatica
func sum(numbers ...int) int {

	var total int

	for _, num := range numbers {
		total += num
	}

	return total
}

func printData(allData ...interface{}) {
	for _, data := range allData {
		fmt.Printf("The data is %v \n", data)
	}
}

func factorial(num int) int {
	if num == 0 {
		return 1
	}

	return num * factorial(num-1)
}

// Anom func
func Saludar(name string, f func(string)) {
	f(name)
}

func Duplicar(n int) int {
	return n * 2
}

func Triplicar(n int) int {
	return n * 3
}

func Aplicar(f func(int) int, n int) int {
	return f(n)
}

func main() {
	fmt.Printf("%v \n", sum(20, 30, 52, 67, 98))
	printData("Hola", 28, true, "Halo")
	fmt.Printf("%v \n", factorial(5))

	greeting := func(name string) {
		fmt.Printf("Hola como estas %s \n", name)
	}

	Saludar("Sau", greeting)
	r1 := Aplicar(Duplicar, 3)
	r2 := Aplicar(Triplicar, 5)

	fmt.Println(r1, r2)
}
