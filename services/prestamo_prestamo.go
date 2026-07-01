package services

import (
	"fmt"
	"sistema-libros/models"
	"sync"
)

var mu sync.Mutex
var Prestamos []models.Prestamo

func RegistrarPrestamo() {

	mu.Lock()
	defer mu.Unlock()

	var prestamo models.Prestamo

	fmt.Println("=== REGISTRO DE PRESTAMO ===")

	fmt.Print("Usuario: ")
	fmt.Scanln(&prestamo.Usuario)

	fmt.Print("Libro: ")
	fmt.Scanln(&prestamo.Libro)

	prestamo.Estado = "Prestado"

	Prestamos = append(Prestamos, prestamo)

	fmt.Println("Prestamo registrado")
}

func MostrarPrestamos() {

	mu.Lock()
	defer mu.Unlock()

	fmt.Println("=== LISTA DE PRESTAMOS ===")

	for _, prestamo := range Prestamos {
		fmt.Println("------------------")
		fmt.Println("Usuario:", prestamo.Usuario)
		fmt.Println("Libro:", prestamo.Libro)
		fmt.Println("Estado:", prestamo.Estado)
	}
}
