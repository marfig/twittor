package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/marfig/twittor/bd"
	"github.com/marfig/twittor/models"
)

func InsertarTweet(w http.ResponseWriter, r *http.Request) {
	var mensaje models.Tweet
	err := json.NewDecoder(r.Body).Decode(&mensaje)
	if err != nil {
		http.Error(w, "Error en los datos recibidos "+err.Error(), http.StatusBadRequest)
		return
	}

	registro := models.GuardarTweet{
		UserID:  IDUsuario,
		Mensaje: mensaje.Mensaje,
		Fecha:   time.Now(),
	}

	var status bool
	_, status, err = bd.InsertarTweet(registro)
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar insertar el registro "+err.Error(), http.StatusBadRequest)
		return
	}

	if status == false {
		http.Error(w, "No se ha logrado insertar el registro ", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
