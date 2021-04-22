package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/marfig/twittor/bd"
)

func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	typeUser := r.URL.Query().Get("type")
	page := r.URL.Query().Get("page")
	search := r.URL.Query().Get("search")

	pageTemp, err := strconv.Atoi(page)
	if err != nil {
		http.Error(w, "Debe enviar el parámetro página como entero mayor a 0", http.StatusBadRequest)
		return
	}

	pag := int64(pageTemp)

	result, status := bd.BuscarUsuarios(IDUsuario, pag, search, typeUser)
	if status == false {
		http.Error(w, "Error al leer los usuarios", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}
