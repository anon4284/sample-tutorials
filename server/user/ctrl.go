package user

import (
	"projects/gobook/server/util"
	"projects/sample-tutorials/server/pkg/crypt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/satori/go.uuid"
)

//Add user to database
func (u *User) Add(input *AddInput) *Output {
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
			"username": &dynamodb.AttributeValue{
				S: aws.String(input.Username),
			},
			"profilePicture": &dynamodb.AttributeValue{
				S: aws.String("https://ssl.gstatic.com/accounts/ui/avatar_2x.png"),
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
		return &Output{false, "Signup failed due to database error", err.Error()}
	}
	return &Output{true, "Signup successfull", ""}

}

func (u *User) Delete(userID string) *Output {
	params := &dynamodb.DeleteItemInput{
		TableName: aws.String(u.DBName),
		Key: map[string]*dynamodb.AttributeValue{
			"userID": &dynamodb.AttributeValue{
				S: aws.String(userID),
			},
		},
	}
	item, err := u.svc.DeleteItem(params)
	util.LogErrOrResp(item, err)

	if err != nil {
		return &Output{false, "User deleted failed due to database", err.Error()}
	}
	return &Output{true, "User deleted successfull", ""}
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

//ChangePassword changes user password in database
func (u *User) ChangePassword(userID string, password string) *Output {
	pw := crypt.HashPassword(password)

	params := &dynamodb.UpdateItemInput{
		TableName: aws.String(u.DBName),
		Key: map[string]*dynamodb.AttributeValue{
			"userID": &dynamodb.AttributeValue{
				S: aws.String(userID),
			},
		},
		AttributeUpdates: map[string]*dynamodb.AttributeValueUpdate{
			"password": &dynamodb.AttributeValueUpdate{
				Value: &dynamodb.AttributeValue{
					S: aws.String(pw),
				},
			},
		},
	}

	item, err := u.svc.UpdateItem(params)
	util.LogErrOrResp(item, err)

	if err != nil {
		return &Output{false, "Password Change failed", err.Error()}
	}
	return &Output{true, "Password Changed Successfully", ""}
}

//GetAll gets all users
func (u *User) GetAll(limit int64) *GetAllOutput {
	params := &dynamodb.ScanInput{
		TableName: aws.String(u.DBName),
		AttributesToGet: []*string{
			aws.String("userID"),
			aws.String("username"),
			aws.String("profilePicture"),
		},
		Limit: aws.Int64(limit),
	}
	resp, err := u.svc.Scan(params)
	util.LogErrOrResp(resp, err)

	if err != nil {
		return &GetAllOutput{false, "Database error", nil, err.Error()}
	}

	arr := []UserInfo{}

	for i := 0; i < len(resp.Items); i++ {
		arr = append(arr, UserInfo{*resp.Items[i]["userID"].S, *resp.Items[i]["username"].S, *resp.Items[i]["profilePicture"].S})
	}

	return &GetAllOutput{true, "Got all users successfully", arr, ""}
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
