package routers

import "net/http"

func Perfil(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
}
