package user

import (
	"projects/gobook/server/util"
	"projects/shir-website/server/pkg/crypt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/satori/go.uuid"
)

//Add user to database
func (u *User) Add(input *AddInput) *AddOutput {
	userID := uuid.NewV4().String()
	password := crypt.HashPassword(input.Password)

	params := &dynamodb.PutItemInput{
		TableName: aws.String(u.DBName),
		Item: map[string]*dynamodb.AttributeValue{
			"userID": &dynamodb.AttributeValue{
				S: aws.String(userID),
			},
			"useremail": &dynamodb.AttributeValue{
				S: aws.String(input.Useremail),
			},
			"password": &dynamodb.AttributeValue{
				S: aws.String(password),
			},
			"createdAt": &dynamodb.AttributeValue{
				S: aws.String(time.Now().String()),
			},
		},
	}
	resp, err := u.svc.PutItem(params)
	util.LogErrOrResp(resp, err)

	if err != nil {
		return &AddOutput{false, "Signup failed due to database error", err.Error()}
	}
	return &AddOutput{true, "Signup successfull", ""}

}

//Get user info from id
func (u *User) Get(userID string) *dynamodb.GetItemOutput {
	params := &dynamodb.GetItemInput{
		TableName: aws.String(u.DBName),
		Key: map[string]*dynamodb.AttributeValue{
			"userID": &dynamodb.AttributeValue{
				S: aws.String(userID),
			},
		},
	}
	item, err := u.svc.GetItem(params)
	util.LogErrOrResp(item, err)
	return item
}

//GetByEmail finds user by email and returns info
func (u *User) GetByEmail(useremail string) *dynamodb.QueryOutput {
	params := &dynamodb.QueryInput{
		TableName:              aws.String(u.DBName),
		IndexName:              aws.String("useremail-index"),
		KeyConditionExpression: aws.String("useremail = :ue"),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":ue": {
				S: aws.String(useremail),
			},
		},
	}

	item, err := u.svc.Query(params)
	util.LogErrOrResp(item, err)
	return item
}

//GetAmount gets amount of users
func (u *User) GetAmount() int64 {
	params := &dynamodb.ScanInput{
		TableName:       aws.String(u.DBName),
		AttributesToGet: []*string{},
	}
	resp, err := u.svc.Scan(params)
	util.LogErrOrResp(resp, err)
	return *resp.Count
}

//EmailUnique checks if email is unique
func (u *User) EmailUnique(useremail string) bool {
	return *u.GetByEmail(useremail).Count == 0
}