package main

import (
	"log"

	"github.com/marfig/twittor/bd"
	"github.com/marfig/twittor/handlers"
)

func main() {
	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin conexión a la BD")
		return
	}
	handlers.Manejadores()
}
