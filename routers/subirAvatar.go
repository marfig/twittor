package routers

import (
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/marfig/twittor/bd"
	"github.com/marfig/twittor/models"
)

func SubirAvatar(w http.ResponseWriter, r *http.Request) {
	file, handler, err := r.FormFile("avatar")
	var extension = strings.Split(handler.Filename, ".")[1]
	var nombre string = IDUsuario + "." + extension
	var archivo string = "uploads/avatars/" + nombre

	f, err := os.OpenFile(archivo, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(w, "Error al subir la imagen "+err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = io.Copy(f, file)
	if err != nil {
		http.Error(w, "Error al copiar la imagen "+err.Error(), http.StatusInternalServerError)
		return
	}

	var usuario models.Usuario
	var status bool

	usuario.Avatar = nombre
	status, err = bd.EditarPerfil(usuario, IDUsuario)
	if err != nil || status == false {
		http.Error(w, "Error al grabar el avatar en la BD "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
