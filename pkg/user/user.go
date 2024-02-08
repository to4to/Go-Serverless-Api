package user

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

var ()

type User struct {
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName`
}

func FetchUser(email,tableName string,dynalient dynamodbiface.DynamoDBAPI)(*User,error) {
input:=&dynamodb.GetItemInput{
	Key: map[],
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
