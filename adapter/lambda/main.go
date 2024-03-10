package main

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/blackhorseya/ekko/adapter/restful"
	"github.com/blackhorseya/ekko/pkg/configx"
	"github.com/spf13/viper"
)

var ginLambda *ginadapter.GinLambda

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return ginLambda.ProxyWithContext(ctx, req)
}

func main() {
	err := configx.LoadWithPathAndName("", "ekko")
	if err != nil {
		log.Fatal(err)
	}

	app, err := restful.NewRestful(viper.GetViper())
	if err != nil {
		log.Fatal(err)
	}

	err = app.InitRouting()
	if err != nil {
		log.Fatal(err)
	}

	ginLambda = ginadapter.New(app.GetRouter())

	lambda.Start(Handler)
}
