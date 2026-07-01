package handlers

import (
	"encoding/json"
	"net/http"

	"sistema-libros/models"
	"sistema-libros/services"
)

// GET /usuarios
func ObtenerUsuarios(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(services.Usuarios)
}

// POST /usuarios/agregar
func RegistrarUsuarioAPI(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Metodo no permitido", http.StatusMethodNotAllowed)
		return
	}

	var usuario models.Usuario

	err := json.NewDecoder(r.Body).Decode(&usuario)

	if err != nil {
		http.Error(w, "Datos invalidos", http.StatusBadRequest)
		return
	}

	services.Usuarios = append(services.Usuarios, usuario)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(usuario)
}
