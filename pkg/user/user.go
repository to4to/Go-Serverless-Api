package user

import (
	"errors"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

var (
	ErrorFailedToFetchRecord = "Failed To Fetch Record"

	ErrorFailedToUnMArshalRecord = "Failed To UnMrashal"
)

type User struct {
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func FetchUser(email, tableName string, dynaClient dynamodbiface.DynamoDBAPI) (*User, error) {
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

		return nil, errors.New(ErrorFailedToFetchRecord)
	}

	item := new(User)

	err = dynamodbattribute.UnmarshalMap(result.Item, item)

	if err != nil {
		return nil, errors.New(ErrorFailedToUnMArshalRecord)
	}
	return item, nil

}

// Fetching Multiple USers
func FetchUsers(tableName string, dynaClient dynamodbiface.DynamoDBAPI) *[]User {
	input := &dynamodb.ScanInput{
		TableName: aws.String(tableName),
	}

}

func CreateUser() {

}

func UpdateMethod() {

}

func DeleteUser() error {

	return nil
}
