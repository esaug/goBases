package main

import (
	"bufio"
	"fmt"
	"os"
)

type Task struct {
	name        string
	description string
	complete    bool
}

type TaskList struct {
	tasks []Task
}

func (l *TaskList) addTask(task Task) {
	l.tasks = append(l.tasks, task)
}

func (l *TaskList) completeTask(index uint8) {
	l.tasks[index].complete = true
}

func (l *TaskList) editTask(index uint8, task Task) {
	l.tasks[index] = task
}

func (l *TaskList) deleteTask(index uint8) {
	l.tasks = append(l.tasks[:index], l.tasks[index+1:]...)
}

func main() {

	var option uint
	list := TaskList{}

	leer := bufio.NewReader(os.Stdin)

	fmt.Print("Seleccione un opcion \n",
		"1. Agregar tarea\n",
		"2. Marcar tarea como completada\n",
		"3. Editar tarea\n",
		"4. Eliminar tarea\n",
		"5. salir\n")

	fmt.Println("Ingrese la opcion")
	fmt.Scanln(&option)

	switch option {
	case 1:

		var t Task
		fmt.Print("Introduce el nombre de la tarea")
		t.name, _ = leer.ReadString('\n')
		fmt.Println("Introduce la descripcion de la tarea")
		t.description, _ = leer.ReadString('\n')
		list.addTask(t)
		fmt.Println("Tarea agregada correctamente")

	case 2:

		fmt.Print("Introduce el indice de la tare a completar")
		var index uint8
		fmt.Scanln(&index)
		list.completeTask(index)
		fmt.Println("Tarea marcada como completado")

	case 3:

		fmt.Print("Introduce el indice para editar tarea")
		var index uint8
		var t Task
		fmt.Scanln(&index)
		fmt.Print("Introduce el nombre de la tarea")
		t.name, _ = leer.ReadString('\n')
		fmt.Println("Introduce la descripcion de la tarea")
		t.description, _ = leer.ReadString('\n')

		list.editTask(index, t)
		fmt.Println("Tarea editada")

	case 4:

		fmt.Print("Introduce el indice para eliminar tarea")
		var index uint8
		list.deleteTask(index)
		fmt.Println("Tarea eliminada")

	case 5:
		fmt.Println("BYE!")
		return
	default:
		fmt.Println("Seleccione una opcion correcta")
	}

	fmt.Println("Lista de tareas")
	fmt.Println("==================================================")

	for index, task := range list.tasks {
		fmt.Printf("%d. %s - %s - Completado: %t\n", index, task.name, task.description, task.complete)
	}
	fmt.Println("==================================================")

}
