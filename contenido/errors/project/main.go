package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

// Contact structure
type Contact struct {
	Name  string `json: name`
	Email string `json: email`
	Phone string `json: phone`
}

// Save contact into a json file
func saveContacts(contacts []Contact) error {

	file, err := os.Create("contacts.json")
	if err != nil {
		return err
	}

	defer file.Close()
	encoder := json.NewEncoder(file)
	err = encoder.Encode(contacts)
	if err != nil {
		return err
	}

	return nil
}

// Load contacts from file
func loadContactsFromFile(contacs *[]Contact) error {
	file, err := os.Open("contacts.json")

	if err != nil {
		return err
	}

	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&contacs)
	if err != nil {
		return err
	}

	return nil
}

func main() {

	var contacts []Contact

	// Load contacts from the file
	err := loadContactsFromFile(&contacts)
	if err != nil {
		fmt.Println("Error al cargar el archivo")
	}

	//Crear instancia de fubio
	reader := bufio.NewReader(os.Stdin)

	for {

		fmt.Print("====Gestor de contactos =====\n",
			"1. Agregar un contacto\n",
			"2. Mostrar todos los contactos\n",
			"3. Salir \n",
			"Elige una opcion: ",
		)

		var option int
		_, err = fmt.Scanln(&option)
		if err != nil {
			fmt.Println("Error al leer la opcion:", err)
			return
		}

		switch option {
		case 1:
			var c Contact
			fmt.Print("Nomber: ")
			c.Name, _ = reader.ReadString('\n')
			fmt.Print("Email: ")
			c.Email, _ = reader.ReadString('\n')
			fmt.Print("Phone: ")
			c.Phone, _ = reader.ReadString('\n')

			// ADD to file
			contacts = append(contacts, c)

			//Guardar en un archivo JSON
			if err := saveContacts(contacts); err != nil {
				fmt.Println("Error al guardar archivo")
			}

		case 2:

			fmt.Println("============================")
			for index, contact := range contacts {
				fmt.Printf("%d Nombre: %s Email: %s Telefono: %s\n",
					index+1, contact.Name, contact.Email, contact.Phone)
			}
			fmt.Println("==========================")
		case 3:
			fmt.Println("Adios!")
			return
		default:
			fmt.Println("Opcion invalida")

		}

	}
}
