package routers

import (
	"encoding/json"
	"net/http"

	"github.com/marfig/twittor/bd"
	"github.com/marfig/twittor/models"
)

// Registro es la función para crear en la BD el usuario
func Registro(w http.ResponseWriter, r *http.Request) {
	var t models.Usuario
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Error en los datos recibidos "+err.Error(), http.StatusBadRequest)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "El email de usuario es requerido", http.StatusBadRequest)
		return
	}

	if len(t.Password) < 6 {
		http.Error(w, "Debe especificar una contraseña de al menos 6 caracteres", http.StatusBadRequest)
		return
	}

	_, encontrado, _ := bd.ChequeoYaExisteUsuario(t.Email)
	if encontrado == true {
		http.Error(w, "Ya existe un usuario registrado con ese email", http.StatusBadRequest)
		return
	}

	_, status, err := bd.InsertarRegistro(t)
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar realizar el registro de usuario "+err.Error(), http.StatusInternalServerError)
		return
	}

	if status == false {
		http.Error(w, "No se ha logrado insertar el registro del usuario", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
