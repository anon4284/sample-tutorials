package portfolio

import (
	"projects/gobook/server/util"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/satori/go.uuid"
)

//Add portfolio item
func (p *Portfolio) Add(input *AddInput, userID string) *AddOutput {
	projectID := uuid.NewV4().String()

	params := &dynamodb.PutItemInput{
		TableName: aws.String(p.DBName),
		Item: map[string]*dynamodb.AttributeValue{
			"projectID": &dynamodb.AttributeValue{
				S: aws.String(projectID),
			},
			"userID": &dynamodb.AttributeValue{
				S: aws.String(userID),
			},
			"title": &dynamodb.AttributeValue{
				S: aws.String(input.Title),
			},
			"description": &dynamodb.AttributeValue{
				S: aws.String(input.Description),
			},
			"content": &dynamodb.AttributeValue{
				S: aws.String(input.Content),
			},
			"createdAt": &dynamodb.AttributeValue{
				S: aws.String(time.Now().String()),
			},
		},
	}
	resp, err := p.svc.PutItem(params)
	util.LogErrOrResp(resp, err)

	if err != nil {
		return &AddOutput{false, "Adding project failed due to database error", err.Error()}
	}
	return &AddOutput{true, "Project added successfully", ""}

}
