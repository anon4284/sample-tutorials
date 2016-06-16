package portfolio

import "github.com/aws/aws-sdk-go/service/dynamodb"

//Portfolio backend component
type Portfolio struct {
	svc         *dynamodb.DynamoDB
	DBName      string
	amountLimit int64
}

//New creates new Portfolio
func New(svc *dynamodb.DynamoDB, dbName string) *Portfolio {
	return &Portfolio{svc: svc, DBName: dbName}
}

//AddInput input to add portfolio item
type AddInput struct {
	Title       string
	Description string
	Content     string
}

//AddOutput output of add operation
type AddOutput struct {
	Valid bool
	Msg   string
	Err   string
}