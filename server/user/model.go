package user

import (
	"projects/gobook/server/util"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

//CreateTable creates table for users
func (u *User) CreateTable() *dynamodb.CreateTableOutput {
	params := &dynamodb.CreateTableInput{
		TableName: aws.String(u.DBName),
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("userID"),
				KeyType:       aws.String("HASH"),
			},
		},
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("userID"),
				AttributeType: aws.String("S"),
			},
			{
				AttributeName: aws.String("useremail"),
				AttributeType: aws.String("S"),
			},
		},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(1),
			WriteCapacityUnits: aws.Int64(1),
		},
		GlobalSecondaryIndexes: []*dynamodb.GlobalSecondaryIndex{
			{ // Required
				IndexName: aws.String("useremail-index"), // Required
				KeySchema: []*dynamodb.KeySchemaElement{ // Required
					{ // Required
						AttributeName: aws.String("useremail"), // Required
						KeyType:       aws.String("HASH"),      // Required
					},
					// More values...
				},
				Projection: &dynamodb.Projection{ // Required
					NonKeyAttributes: []*string{
						aws.String("password"), // Required
						// More values...
					},
					ProjectionType: aws.String("INCLUDE"),
				},
				ProvisionedThroughput: &dynamodb.ProvisionedThroughput{ // Required
					ReadCapacityUnits:  aws.Int64(1), // Required
					WriteCapacityUnits: aws.Int64(1), // Required
				},
			},
			// More values...
		},
	}
	resp, err := u.svc.CreateTable(params)
	util.LogErrOrResp(resp, err)
	return resp
}

//AddInput input required to add the user
type AddInput struct {
	Useremail string
	Password  string
}

//AddOutput output of add operation
type AddOutput struct {
	Valid bool
	Msg   string
	Err   string
}

//LoginInput used to login to app
type LoginInput struct {
	Useremail string
	Password  string
}

//LoginOutput output returned from login
type LoginOutput struct {
	Valid  bool
	Token  string
	UserID string
	Msg    string
}
