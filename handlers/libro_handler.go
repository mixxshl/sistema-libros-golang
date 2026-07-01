package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"sistema-libros/models"
	"sistema-libros/services"
)

// GET /libros
func ObtenerLibros(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(services.Libros)
}

// POST /libros/agregar
func AgregarLibroAPI(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Metodo no permitido", http.StatusMethodNotAllowed)
		return
	}

	var libro models.Libro

	err := json.NewDecoder(r.Body).Decode(&libro)

	if err != nil {
		http.Error(w, "Datos invalidos", http.StatusBadRequest)
		return
	}

	libro.Disponible = true
	services.Libros = append(services.Libros, libro)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(libro)
}

// PUT /libros?id=1
func ActualizarLibroAPI(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPut {
		http.Error(w, "Metodo no permitido", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		http.Error(w, "ID invalido", http.StatusBadRequest)
		return
	}

	var libroActualizado models.Libro

	err = json.NewDecoder(r.Body).Decode(&libroActualizado)

	if err != nil {
		http.Error(w, "Datos invalidos", http.StatusBadRequest)
		return
	}

	for i, libro := range services.Libros {
		if libro.ID == id {
			services.Libros[i].Titulo = libroActualizado.Titulo
			services.Libros[i].Autor = libroActualizado.Autor
			services.Libros[i].Categoria = libroActualizado.Categoria
			json.NewEncoder(w).Encode(services.Libros[i])
			return
		}
	}

	http.Error(w, "Libro no encontrado", http.StatusNotFound)
}

// DELETE /libros?id=1
func EliminarLibroAPI(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodDelete {
		http.Error(w, "Metodo no permitido", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		http.Error(w, "ID invalido", http.StatusBadRequest)
		return
	}

	for i, libro := range services.Libros {
		if libro.ID == id {
			services.Libros = append(services.Libros[:i], services.Libros[i+1:]...)
			w.Write([]byte("Libro eliminado correctamente"))
			return
		}
	}

	http.Error(w, "Libro no encontrado", http.StatusNotFound)
}
