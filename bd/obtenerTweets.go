package bd

import (
	"context"
	"log"
	"time"

	"github.com/marfig/twittor/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ObtenerTweets(ID string, pagina int64) ([]*models.RespuestaTweet, bool) {
	var pageSize int64 = 20

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()
	db := MongoCN.Database("twittor-db")
	col := db.Collection("tweet")

	var resultados []*models.RespuestaTweet

	condicion := bson.M{
		"userid": ID,
	}

	opciones := options.Find()
	opciones.SetLimit(20)
	opciones.SetSort(bson.D{{Key: "fecha", Value: -1}})
	opciones.SetSkip((pagina - 1) * pageSize)

	cursor, err := col.Find(ctx, condicion, opciones)
	if err != nil {
		log.Fatal(err.Error())
		return resultados, false
	}

	for cursor.Next(context.TODO()) {
		var registro models.RespuestaTweet
		err := cursor.Decode(&registro)
		if err != nil {
			return resultados, false
		}
		resultados = append(resultados, &registro)
	}

	return resultados, true
}
