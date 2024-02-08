package user

import (
	"errors"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

var (

	ErrorFailedToFetchRecord="Failed To Fetch Record"
)

type User struct {
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func FetchUser(email, tableName string, dynalient dynamodbiface.DynamoDBAPI) (*User, error) {
	input := &dynamodb.GetItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			email: {
				S: aws.String(email),
			},
		},
		TableName: aws.String(tableName),
	}

	result, err := dynalient.GetItem(input)

	if err != nil {

	return nil,errors.New(ErrorFailedToFetchRecord)
}

item :=new(User)

err=dynamodbattribute.UnmarshalMap(result.Item,item)


if err!=nil{

}


}

// Fetching Multiple USers
func FetchUsers() {

}

func CreateUser() {

}

func UpdateMethod() {

}

func DeleteUser() error {

	return nil
}
