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
				AttributeName: aws.String("userID"),
				KeyType:       aws.String("HASH"),
			},
			{
				AttributeName: aws.String("projectID"),
				KeyType:       aws.String("RANGE"),
			},
		},
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("userID"),
				AttributeType: aws.String("S"),
			},
			{
				AttributeName: aws.String("projectID"),
				AttributeType: aws.String("S"),
			},
		},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(1),
			WriteCapacityUnits: aws.Int64(1),
		},
	}

	resp, err := p.svc.CreateTable(params)
	util.LogErrOrResp(resp, err)
	return resp
}
