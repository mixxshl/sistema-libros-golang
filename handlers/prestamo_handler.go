package handlers

import (
	"encoding/json"
	"net/http"

	"sistema-libros/models"
	"sistema-libros/services"
)

// GET /prestamos
func ObtenerPrestamos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(services.Prestamos)
}

// POST /prestamos
func RegistrarPrestamoAPI(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Metodo no permitido", http.StatusMethodNotAllowed)
		return
	}

	var prestamo models.Prestamo

	err := json.NewDecoder(r.Body).Decode(&prestamo)

	if err != nil {
		http.Error(w, "Datos invalidos", http.StatusBadRequest)
		return
	}

	prestamo.Estado = "Prestado"

	services.Prestamos = append(services.Prestamos, prestamo)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(prestamo)
}
