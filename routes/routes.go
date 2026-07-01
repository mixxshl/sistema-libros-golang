package routes

import (
	"net/http"

	"sistema-libros/handlers"
)

func ConfigurarRutas() {

	// LIBROS
	http.HandleFunc("/libros", handlers.ObtenerLibros)
	http.HandleFunc("/libros/agregar", handlers.AgregarLibroAPI)
	http.HandleFunc("/libros/actualizar", handlers.ActualizarLibroAPI)
	http.HandleFunc("/libros/eliminar", handlers.EliminarLibroAPI)

	// USUARIOS
	http.HandleFunc("/usuarios", handlers.ObtenerUsuarios)
	http.HandleFunc("/usuarios/agregar", handlers.RegistrarUsuarioAPI)

	// PRESTAMOS
	http.HandleFunc("/prestamos", handlers.ObtenerPrestamos)
	http.HandleFunc("/prestamos/agregar", handlers.RegistrarPrestamoAPI)
}
