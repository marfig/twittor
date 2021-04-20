package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/marfig/twittor/middlew"
	"github.com/marfig/twittor/routers"
	"github.com/rs/cors"
)

// Manejadores seteo mi puerto, el Handler y pongo a escuchar al servidor
func Manejadores() {
	router := mux.NewRouter()

	router.HandleFunc("/registro", middlew.ChequeoBD(routers.Registro)).Methods("POST")
	router.HandleFunc("/login", middlew.ChequeoBD(routers.Login)).Methods("POST")
	router.HandleFunc("/perfil", middlew.ChequeoBD(middlew.ValidarJWT(routers.Perfil))).Methods("GET")
	router.HandleFunc("/perfil", middlew.ChequeoBD(middlew.ValidarJWT(routers.ModificarPerfil))).Methods("PUT")
	router.HandleFunc("/tweet", middlew.ChequeoBD(middlew.ValidarJWT(routers.InsertarTweet))).Methods("POST")
	router.HandleFunc("/tweets", middlew.ChequeoBD(middlew.ValidarJWT(routers.ObtenerTweets))).Methods("GET")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
