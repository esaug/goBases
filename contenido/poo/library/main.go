package main

import (
	animal "library/animals"
)

func main() {
	/*var myBook = book.NewBook("La caida", "Sau", 200)

	myBook.PrintInfo()

	myBook.SetTittle("El maduro muerto")
	fmt.Printf("Nuevo titulo %s \n", myBook.GetTitle())

	textBook := book.NewTextBook("Dictador2000", "Sau", 300, "El Mundo", 30)
	textBook.PrintTextBookInfo() */

	/*myDog := animal.Dog{Name: "Pepito"}
	myCat := animal.Cat{Name: "Alberto"}

	animal.MakeNoise(&myDog)
	animal.MakeNoise(&myCat) */

	animals := []animal.Animal{
		&animal.Dog{Name: "Alex"},
		&animal.Cat{Name: "Patroclo"},
		&animal.Dog{Name: "Albarot"},
		&animal.Cat{Name: "Veltreco"},
	}

	for _, animal := range animals {
		animal.Sound()
	}

}
