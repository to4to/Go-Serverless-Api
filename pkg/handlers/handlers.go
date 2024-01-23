package handlers

import "github.com/aws/aws-lambda-go/events"



var ErrorMethodNotAllowed="method not allowed"

func GetUser(){}



func CreateUser(){}



func UpdateUser(){}

func DeleteUser(){

}


func UnhandledMethod()(*events.APIGatewayProxyResponse,error){



}