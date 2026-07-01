package main

import (
	"fmt"
	"net/http"

	"sistema-libros/routes"
)

func main() {

	routes.ConfigurarRutas()

	fmt.Println("====================================")
	fmt.Println("Servidor iniciado correctamente")
	fmt.Println("http://localhost:8080")
	fmt.Println("====================================")

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Println("Error al iniciar el servidor:", err)
	}
}
