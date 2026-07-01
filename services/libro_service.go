package services

import (
	"fmt"
	"sistema-libros/models"
)

var Libros []models.Libro

func AgregarLibro() {

	mu.Lock()
	defer mu.Unlock()

	var libro models.Libro

	fmt.Println("=== REGISTRO DE LIBRO ===")

	fmt.Print("ID: ")
	fmt.Scanln(&libro.ID)

	fmt.Print("Titulo: ")
	fmt.Scanln(&libro.Titulo)

	fmt.Print("Autor: ")
	fmt.Scanln(&libro.Autor)

	fmt.Print("Categoria: ")
	fmt.Scanln(&libro.Categoria)

	libro.Disponible = true

	Libros = append(Libros, libro)

	fmt.Println("Libro agregado correctamente")
}

func MostrarLibros() {

	mu.Lock()
	defer mu.Unlock()

	fmt.Println("=== LISTA DE LIBROS ===")

	for _, libro := range Libros {
		fmt.Println("------------------")
		fmt.Println("ID:", libro.ID)
		fmt.Println("Titulo:", libro.Titulo)
		fmt.Println("Autor:", libro.Autor)
		fmt.Println("Categoria:", libro.Categoria)
		fmt.Println("Disponible:", libro.Disponible)
	}
}
