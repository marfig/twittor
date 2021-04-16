package bd

import (
	"context"
	"time"

	"github.com/marfig/twittor/models"
	"go.mongodb.org/mongo-driver/bson"
)

// ChequeoYaExisteUsuario chequea si el email ya fue registrado anteriormente
func ChequeoYaExisteUsuario(email string) (models.Usuario, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor-db")
	col := db.Collection("usuarios")

	condicion := bson.M{"email": email}

	var resultado models.Usuario

	err := col.FindOne(ctx, condicion).Decode(&resultado)
	ID := resultado.ID.Hex()
	if err != nil {
		return resultado, false, ID
	}
	return resultado, true, ID
}
