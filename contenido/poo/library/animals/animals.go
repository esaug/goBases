package animal

import "fmt"

type Animal interface {
	Sound()
}

type Dog struct {
	Name string
}

func (p *Dog) Sound() {
	fmt.Printf("El perro %s hace: Guau gua guau \n", p.Name)
}

type Cat struct {
	Name string
}

func (p *Cat) Sound() {
	fmt.Print("El gato %s hace: Miau, miua, miua \n", p.Name)
}

func MakeNoise(animal Animal) {
	animal.Sound()
}
