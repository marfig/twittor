package routers

import (
	"encoding/json"
	"net/http"

	"github.com/marfig/twittor/bd"
	"github.com/marfig/twittor/models"
)

func ModificarPerfil(w http.ResponseWriter, r *http.Request) {

	var t models.Usuario

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Datos Incorrectos"+err.Error(), http.StatusBadRequest)
		return
	}

	var status bool
	status, err = bd.EditarPerfil(t, IDUsuario)
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar modificar el registro "+err.Error(), http.StatusBadRequest)
		return
	}

	if status == false {
		http.Error(w, "No se ha logrado modificar el registro del usuario", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
