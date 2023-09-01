package bd

import (
	"context"

	"github.com/johnrios07/twitterGo/models"
	"go.mongodb.org/mongo-driver/bson"
)

func ChequeoYaExisteUsuario(email string) (models.Usuario, bool, string) {
	ctx := context.TODO()

	db := MongoCN.Database(DatabaseName)
	col := db.Collection("usuario")

	condition := bson.M{"email": email}

	var resultado models.Usuario

	err := col.FindOne(ctx, condition).Decode(&resultado)
	ID := resultado.ID.Hex()
	if err != nil {
		return resultado, false, ID
	}

	return resultado, true, ID
}
