package handlers

import (
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)



var ErrorMethodNotAllowed="method not allowed"


type ErrorBody struct{
	ErrorMsg *string `json:"error,ommitempty" `
}

func GetUser(req events.APIGatewayProxyRequest,tableName string,dynaClient dynamodbiface.DynamoDBAPI)
(*events.APIGatewayProxyResponse,error){



}



func CreateUser(req events.APIGatewayProxyRequest,tableName string,dynaClient dynamodbiface.DynamoDBAPI)
(*events.APIGatewayProxyResponse,error){}



func UpdateUser(){}

func DeleteUser(){

}


func UnhandledMethod()(*events.APIGatewayProxyResponse,error){

return apiResponse(http.StatusMethodNotAllowed,ErrorMethodNotAllowed)

}