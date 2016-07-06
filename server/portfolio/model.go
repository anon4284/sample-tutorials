package portfolio

//CreateTable creates table for users
import (
	"projects/gobook/server/util"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

//CreateTable creates portfolio table
func (p *Portfolio) CreateTable() *dynamodb.CreateTableOutput {
	params := &dynamodb.CreateTableInput{
		TableName: aws.String(p.DBName),
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("projectID"),
				KeyType:       aws.String("HASH"),
			},
			{
				AttributeName: aws.String("createdAt"),
				KeyType:       aws.String("RANGE"),
			},
		},
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("projectID"),
				AttributeType: aws.String("S"),
			},
			{
				AttributeName: aws.String("userID"),
				AttributeType: aws.String("S"),
			},
			{
				AttributeName: aws.String("createdAt"),
				AttributeType: aws.String("S"),
			},
		},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(1),
			WriteCapacityUnits: aws.Int64(1),
		},
		GlobalSecondaryIndexes: []*dynamodb.GlobalSecondaryIndex{
			{
				IndexName: aws.String("userID-index"),
				KeySchema: []*dynamodb.KeySchemaElement{
					{
						AttributeName: aws.String("userID"),
						KeyType:       aws.String("HASH"),
					},
					{
						AttributeName: aws.String("createdAt"),
						KeyType:       aws.String("RANGE"),
					},
				},
				Projection: &dynamodb.Projection{
					NonKeyAttributes: []*string{
						aws.String("projectID"),
						aws.String("createdAt"),
						aws.String("title"),
						aws.String("description"),
						aws.String("views"),
						aws.String("likes"),
						aws.String("dislikes"),
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

	resp, err := p.svc.CreateTable(params)
	util.LogErrOrResp(resp, err)
	return resp
}
