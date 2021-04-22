package routers

import (
	"net/http"

	"github.com/marfig/twittor/bd"
	"github.com/marfig/twittor/models"
)

func EliminarRelacion(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parámetro ID", http.StatusBadRequest)
		return
	}

	var t models.Relacion
	t.UsuarioID = IDUsuario
	t.UsuarioRelacionID = ID

	status, err := bd.EliminarRelacion(t)
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar borrar el registro"+err.Error(), http.StatusInternalServerError)
		return
	}

	if status == false {
		http.Error(w, "No se ha logrado eliminar el registro", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
