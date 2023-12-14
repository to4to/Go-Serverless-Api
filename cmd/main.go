package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/to4to/go-serverless-api/pkg/handlers"
	"os"
	//"github.com/aws/aws-sdk-go/service/lambda"
)

func main() {

	// region := os.Getenv("AWS_REGION")
	// awsSession, err := session.NewSession(&aws.Config{
	// 	Region: aws.String(region),
	// })

	// if err != nil {
	// 	return
	// }

	// dynaClient := dynamodb.New(awsSession)

	lambda.Start(handler)
}

const tableName = "LambdaInGoUser"

func handler(req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	region := os.Getenv("AWS_REGION")
	awsSession, err := session.NewSession(&aws.Config{
		Region: aws.String(region),
	})

	if err != nil {
		return nil, err
	}

	dynaClient := dynamodb.New(awsSession)
	switch req.HTTPMethod {

	case "GET":
		return handlers.GetUser(req, tableName, dynaClient)
	}
}
