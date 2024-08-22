package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

func printList(list ...any) {
	for _, value := range list {
		fmt.Printf("%v\n", value)
	}
}

// Funcion generica
type integer int

type Integer interface {
	~int | ~float64 | ~float32 | ~uint
}

func include[T comparable](list []T, value T) bool {

	for _, item := range list {
		if item == value {
			return true
		}
	}

	return false
}

func filter[T constraints.Ordered](list []T, callBack func(T) bool) []T {
	result := make([]T, 0, len(list))

	for _, item := range list {
		if callBack(item) {
			result = append(result, item)
		}
	}

	return result
}

func sum[T constraints.Integer | constraints.Float](list ...T) T {
	var totalSum T

	for _, value := range list {
		totalSum += value
	}

	return totalSum
}

// Structuras genericas
type Product[T string | uint] struct {
	Id    T
	Desc  string
	Price float32
}

func main() {
	printList("hola", "Chao", "Bye", "Aloja")

	var inter integer = 100
	var inter2 integer = 300
	fmt.Println(sum[uint](100, 200, 1000, 111, 300))
	fmt.Println(sum[float32](55.6, 60.1, 22.9, 33.23, 10.11))
	fmt.Println(sum(inter, inter2))

	// Restriccion y operadores

	strings := []string{"a", "b", "c", "d"}
	numbers := []int{1, 2, 3, 4, 5}

	fmt.Printf("%v \n", include(strings, "a"))
	fmt.Printf("%v \n", include(numbers, 5))

	fmt.Println(filter(numbers, func(value int) bool { return value > 3 }))
	fmt.Println(filter(strings, func(value string) bool { return value > "b" }))

	// Structuras genericas

	product := Product[uint]{1, "Cuaderno", 13.42}
	product2 := Product[string]{"aklsdjaskjdsak", "Cuaderno", 20}

	fmt.Println(product, product2)

}
