package handlers

import (
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)



var ErrorMethodNotAllowed="method not allowed"


type ErrorBody struct{
	ErrorMsg *string `json:"error,ommitempty" `
}

func GetUser(){}



func CreateUser(){}



func UpdateUser(){}

func DeleteUser(){

}


func UnhandledMethod()(*events.APIGatewayProxyResponse,error){

return apiResponse(http.StatusMethodNotAllowed,ErrorMethodNotAllowed)

}