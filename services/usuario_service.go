package services

import (
	"fmt"
	"sistema-libros/models"
)

var Usuarios []models.Usuario

func RegistrarUsuario() {

	mu.Lock()
	defer mu.Unlock()

	var usuario models.Usuario

	fmt.Println("=== REGISTRO DE USUARIO ===")

	fmt.Print("ID: ")
	fmt.Scanln(&usuario.ID)

	fmt.Print("Nombre: ")
	fmt.Scanln(&usuario.Nombre)

	fmt.Print("Correo: ")
	fmt.Scanln(&usuario.Correo)

	Usuarios = append(Usuarios, usuario)

	fmt.Println("Usuario registrado")
}
func MostrarUsuarios() {

	mu.Lock()
	defer mu.Unlock()

	fmt.Println("=== LISTA DE USUARIOS ===")

	for _, usuario := range Usuarios {
		fmt.Println("------------------")
		fmt.Println("ID:", usuario.ID)
		fmt.Println("Nombre:", usuario.Nombre)
		fmt.Println("Correo:", usuario.Correo)
	}
}
