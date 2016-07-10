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

	o := validateAdd(input)

	if !o.Valid {
		return o
	}

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

//GetByUser finds user by email and returns info
func (p *Portfolio) GetByUser(userID string) *dynamodb.QueryOutput {
	params := &dynamodb.QueryInput{
		TableName:              aws.String(p.DBName),
		IndexName:              aws.String("userID-index"),
		KeyConditionExpression: aws.String("userID = :uID"),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":uID": {
				S: aws.String(userID),
			},
		},
	}

	item, err := p.svc.Query(params)
	util.LogErrOrResp(item, err)
	return item
}

func (p *Portfolio) GetByProjectID(projectID string) *GetSingleOutput {
	params := &dynamodb.QueryInput{
		TableName:              aws.String(p.DBName),
		KeyConditionExpression: aws.String("projectID = :pID"),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":pID": {
				S: aws.String(projectID),
			},
		},
	}

	resp, err := p.svc.Query(params)
	util.LogErrOrResp(resp, err)
	if err != nil {
		return &GetSingleOutput{false, "Failed to get projects due to databse error", ProjectSingle{}, err.Error()}
	}
	return &GetSingleOutput{true, "Retrieved Project successfully", ProjectSingle{*resp.Items[0]["projectID"].S, *resp.Items[0]["title"].S, *resp.Items[0]["description"].S, *resp.Items[0]["content"].S}, ""}
}

func (p *Portfolio) Scan() *GetOutput {
	params := &dynamodb.ScanInput{
		TableName: aws.String(p.DBName),
		AttributesToGet: []*string{
			aws.String("projectID"),
			aws.String("title"),
			aws.String("description"),
		},
	}
	resp, err := p.svc.Scan(params)
	util.LogErrOrResp(resp, err)
	if err != nil {
		return &GetOutput{false, "Failed to get projects due to databse error", nil, err.Error()}
	}

	arr := []Project{}

	for i := 0; i < len(resp.Items); i++ {
		arr = append(arr, Project{*resp.Items[i]["projectID"].S, *resp.Items[i]["title"].S, *resp.Items[i]["description"].S})
	}

	return &GetOutput{true, "Successfully got projects", arr, ""}

}
