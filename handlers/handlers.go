package handlers

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/johnrios07/twitterGo/jwt"
	"github.com/johnrios07/twitterGo/models"
)

func Manejadores(ctx context.Context, request events.APIGatewayProxyRequest) models.RespApi {
	fmt.Println("Voy a procesar " + ctx.Value(models.Key("path")).(string) + " > " + ctx.Value(models.Key("method")).(string))

	var resp models.RespApi
	resp.Status = 400

	isOk, statusCode, msg, claim := validoAuthorization(ctx, request)

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

func validoAuthorization(ctx context.Context, request events.APIGatewayProxyRequest) (bool, int, string, models.Claim) {
	path := ctx.Value(models.Key("path"))
	if path == "registro" || path == "login" || path == "obtnerAvatar" || path == "obtenerBanner" {
		return true, 200, "", models.Claim{}
	}

	token := request.Headers["Authorization"]

	if len(token) == 0 {
		return false, 401, "Token requerido", models.Claim{}
	}

	claim, todoOk, msg, err := jwt.ProcesoToken(token, ctx.Value(models.Key("jwtSign")).(string))
	if !todoOk {
		if err != nil {
			fmt.Println("Error en el token " + err.Error())
			return false, 401, err.Error(), models.Claim{}
		} else {
			fmt.Println("Error en el token " + msg)
			return false, 401, msg, models.Claim{}
		}
	}

	fmt.Println("Token Ok")
	return true, 200, msg, *claim

}
