package utils

import (
	"fmt"
	"sistema-libros/services"
)

func MenuPrincipal() {

	var opcion int

	for {
		fmt.Println("\n===== SISTEMA DE LIBROS =====")
		fmt.Println("1. Agregar libro")
		fmt.Println("2. Mostrar libros")
		fmt.Println("3. Registrar usuario")
		fmt.Println("4. Mostrar usuarios")
		fmt.Println("5. Registrar prestamo")
		fmt.Println("6. Mostrar prestamos")
		fmt.Println("7. Salir")

		fmt.Print("Seleccione una opcion: ")
		fmt.Scanln(&opcion)

		switch opcion {

		case 1:
			services.AgregarLibro()

		case 2:
			services.MostrarLibros()

		case 3:
			services.RegistrarUsuario()

		case 4:
			services.MostrarUsuarios()

		case 5:
			services.RegistrarPrestamo()

		case 6:
			services.MostrarPrestamos()

		case 7:
			fmt.Println("Saliendo del sistema...")
			return

		default:
			fmt.Println("Opcion invalida")
		}
	}
}
