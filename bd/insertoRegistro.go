package bd

import (
	"context"

	"github.com/johnrios07/twitterGo/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func InsertoRegistro(u models.Usuario) (string, bool, error) {
	ctx := context.TODO()

	db := MongoCN.Database(DatabaseName)
	col := db.Collection("usuario")

	u.Password, _ = EncriptarPassword(u.Password)

	result, err := col.InsertOne(ctx, u)
	if err != nil {
		return "", false, err
	}

	ObjId, _ := result.InsertedID.(primitive.ObjectID)

	return ObjId.String(), true, nil
}
