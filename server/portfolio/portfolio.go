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

//GetOutput output of get operation
type GetOutput struct {
	Valid bool
	Msg   string
	Items []Project
	Err   string
}

//GetSingleOutput output of get operation
type GetSingleOutput struct {
	Valid bool
	Msg   string
	Item  ProjectSingle
	Err   string
}

//Project to define databse Project
type Project struct {
	ProjectID   string
	Title       string
	Description string
}

//ProjectSingle like project but with content
type ProjectSingle struct {
	ProjectID   string
	Title       string
	Description string
	Content     string
}
