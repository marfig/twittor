package routers

import (
	"encoding/json"
	"net/http"

	"github.com/marfig/twittor/bd"
)

// Perfil devuelve el perfil del usuario
func Perfil(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parámetro ID", http.StatusBadRequest)
		return
	}

	perfil, err := bd.ObtenerPerfil(ID)
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar buscar el registro"+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(perfil)
}
