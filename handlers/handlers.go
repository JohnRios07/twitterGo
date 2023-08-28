package handlers

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/johnrios07/twitterGo/models"
)

func Manejadores(ctx context.Context, request events.APIGatewayProxyRequest) models.RespApi {
	fmt.Println("Voy a procesar " + ctx.Value(models.Key("path")).(string) + " > " + ctx.Value(models.Key("method")).(string))

	var resp models.RespApi
	resp.Status = 400

	switch ctx.Value(models.Key("method")).(string) {
	case "POST":
		switch ctx.Value(models.Key("path")).(string) {

		}
	case "get":
		switch ctx.Value(models.Key("path")).(string) {

		}

	case "PUT":
		switch ctx.Value(models.Key("path")).(string) {

		}

	case "DELETE":
		switch ctx.Value(models.Key("path")).(string) {

		}

	}

	resp.Message = "method invalid"
	return resp
}
