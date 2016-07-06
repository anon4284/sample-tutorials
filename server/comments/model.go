package comments

import (
	"projects/gobook/server/util"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

//CreateTable creates portfolio table
func (c *Comments) CreateTable() *dynamodb.CreateTableOutput {
	params := &dynamodb.CreateTableInput{
		TableName: aws.String(c.DBName),
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("projectID"),
				KeyType:       aws.String("HASH"),
			},
			{
				AttributeName: aws.String("commentID"),
				KeyType:       aws.String("RANGE"),
			},
		},
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("projectID"),
				AttributeType: aws.String("S"),
			},
			{
				AttributeName: aws.String("commentID"),
				AttributeType: aws.String("S"),
			},
		},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(1),
			WriteCapacityUnits: aws.Int64(1),
		},
	}

	resp, err := c.svc.CreateTable(params)
	util.LogErrOrResp(resp, err)
	return resp
}
