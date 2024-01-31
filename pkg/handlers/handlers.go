package handlers

import (
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/to4to/go-serverless-api/pkg/user"
)



var ErrorMethodNotAllowed="method not allowed"


type ErrorBody struct{
	ErrorMsg *string `json:"error,ommitempty" `
}

// func GetUser(req events.APIGatewayProxyRequest,tableName string,dynaClient dynamodbiface.DynamoDBAPI)
// (*events.APIGatewayProxyResponse,error)

func GetUser(req events.APIGatewayProxyRequest, tableName string, dynaClient dynamodbiface.DynamoDBAPI) (*events.APIGatewayProxyResponse, error) {
    email :=req.QueryStringParameters["email"]
	if len(email)>0{
		result, err:= user.FetchUser(email,tableName,dynaClient)
		if err!=nil{
			return apiResponse(http.StatusBadRequest,ErrorBody{aws.String(err.Error())})
		}
	}

}






func CreateUser(req events.APIGatewayProxyRequest,tableName string,dynaClient dynamodbiface.DynamoDBAPI)(*events.APIGatewayProxyResponse,error){


}



func UpdateUser(req events.APIGatewayProxyRequest,tableName string,dynaClient dynamodbiface.DynamoDBAPI)(*events.APIGatewayProxyResponse,error){}

func DeleteUser(req events.APIGatewayProxyRequest,tableName string,dynaClient dynamodbiface.DynamoDBAPI)(*events.APIGatewayProxyResponse,error){

}


func UnhandledMethod()(*events.APIGatewayProxyResponse,error){

return apiResponse(http.StatusMethodNotAllowed,ErrorMethodNotAllowed)

}