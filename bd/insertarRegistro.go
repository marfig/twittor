package bd

import (
	"context"
	"time"

	"github.com/marfig/twittor/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//InsertarRegistro inserta los datos de usuario
func InsertarRegistro(u models.Usuario) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor-db")
	col := db.Collection("usuarios")

	u.Password, _ = EncriptarPassword(u.Password)

	result, err := col.InsertOne(ctx, u)
	if err != nil {
		return "", false, err
	}

	ObjID, _ := result.InsertedID.(primitive.ObjectID)
	return ObjID.String(), true, nil
}
