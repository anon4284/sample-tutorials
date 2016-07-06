package comments

import (
	"projects/gobook/server/util"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/satori/go.uuid"
)

//Add portfolio item
func (c *Comments) Add(input *AddInput, userID string) *AddOutput {

	o := c.validateAdd(input)

	if !o.Valid {
		return o
	}

	commentID := uuid.NewV4().String()
	params := &dynamodb.PutItemInput{
		TableName: aws.String(c.DBName),
		Item: map[string]*dynamodb.AttributeValue{
			"commentID": &dynamodb.AttributeValue{
				S: aws.String(commentID),
			},
			"projectID": &dynamodb.AttributeValue{
				S: aws.String(input.ProjectID),
			},
			"userID": &dynamodb.AttributeValue{
				S: aws.String(userID),
			},
			"content": &dynamodb.AttributeValue{
				S: aws.String(input.Content),
			},
			"createdAt": &dynamodb.AttributeValue{
				S: aws.String(time.Now().String()),
			},
		},
	}
	resp, err := c.svc.PutItem(params)
	util.LogErrOrResp(resp, err)

	if err != nil {
		return &AddOutput{false, "Adding comment failed due to database error", err.Error()}
	}
	return &AddOutput{true, "Comment added successfully", ""}

}

func (c *Comments) GetByProjectID(projectID string) *GetOutput {
	params := &dynamodb.QueryInput{
		TableName:              aws.String(c.DBName),
		KeyConditionExpression: aws.String("projectID = :pID"),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":pID": {
				S: aws.String(projectID),
			},
		},
	}
	resp, err := c.svc.Query(params)
	util.LogErrOrResp(resp, err)
	if err != nil {
		return &GetOutput{false, "Failed to get projects due to databse error", []Comment{}, err.Error()}
	}

	arr := []Comment{}

	for i := 0; i < len(resp.Items); i++ {
		user := c.user.Get(*resp.Items[i]["userID"].S)
		arr = append(arr, Comment{*resp.Items[i]["userID"].S, *user.Item["username"].S, *user.Item["profilePicture"].S, *resp.Items[i]["content"].S, *resp.Items[i]["createdAt"].S})
	}
	return &GetOutput{true, "Got Comments Successfully", arr, ""}
}
