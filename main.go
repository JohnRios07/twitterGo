package main

import (
	"context"
	"os"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	lambda "github.com/aws/aws-lambda-go/lambda"

	"github.com/johnrios07/twitterGo/awsgo"
	"github.com/johnrios07/twitterGo/bd"
	"github.com/johnrios07/twitterGo/handlers"
	"github.com/johnrios07/twitterGo/models"
	"github.com/johnrios07/twitterGo/secretmanager"
)

func main() {
	lambda.Start(EjecutoLambda)
}

func EjecutoLambda(ctx context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	var resp *events.APIGatewayProxyResponse

	awsgo.InicializoAWS()

	if !Validoparametros() {
		resp = &events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       "Error en las variables de entorno, deben incluir 'SecretName', 'BucketName', 'UrlPrefix'",
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
		}
		return resp, nil
	}

	SecretModel, err := secretmanager.GetSecret(os.Getenv("SecretName"))
	if err != nil {
		resp = &events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       "Error en la lectura de Secret " + err.Error(),
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
		}
		return resp, nil
	}

	path := strings.Replace(request.PathParameters["twitterGo"], os.Getenv("UrlPrefix"), "", -1)

	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("path"), path)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("method"), request.HTTPMethod)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("user"), SecretModel.Username)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("password"), SecretModel.Password)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("host"), SecretModel.Host)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("database"), SecretModel.Database)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("jwtsign"), SecretModel.JWTSign)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("body"), request.Body)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("bucketName"), os.Getenv("BucketName"))

	// Chequeo conexion a la BD o conecto a la BD
	err = bd.ConectarBD(awsgo.Ctx)
	if err != nil {
		resp = &events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       "Error conectando a la BD " + err.Error(),
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
		}
		return resp, nil
	}

	respApi := handlers.Manejadores(awsgo.Ctx, request)

	if respApi.CustomResp == nil {
		resp = &events.APIGatewayProxyResponse{
			StatusCode: respApi.Status,
			Body:       respApi.Message,
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
		}
		return resp, nil
	} else {
		return respApi.CustomResp, nil
	}
}

func Validoparametros() bool {
	_, traeParametro := os.LookupEnv("SecretName")
	if !traeParametro {
		return traeParametro
	}

	_, traeParametro = os.LookupEnv("BucketName")
	if !traeParametro {
		return traeParametro
	}

	_, traeParametro = os.LookupEnv("UrlPrefix")
	if !traeParametro {
		return traeParametro
	}

	return traeParametro
}
