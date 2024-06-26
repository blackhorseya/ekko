package main

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/blackhorseya/ekko/adapter/platform/rest"
	"github.com/blackhorseya/ekko/pkg/adapterx"
	"github.com/spf13/viper"
)

var ginLambda *ginadapter.GinLambdaV2

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(ctx context.Context, req events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	return ginLambda.ProxyWithContext(ctx, req)
}

func main() {
	service, err := rest.New(viper.GetViper())
	if err != nil {
		panic(err)
	}

	restful, ok := service.(adapterx.Restful)
	if !ok {
		panic("service is not restful")
	}

	err = restful.InitRouting()
	if err != nil {
		panic(err)
	}

	ginLambda = ginadapter.NewV2(restful.GetRouter())

	lambda.Start(Handler)
}
