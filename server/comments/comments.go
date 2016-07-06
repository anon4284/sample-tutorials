package comments

import (
	"projects/sample-tutorials/server/portfolio"
	"projects/sample-tutorials/server/user"

	"github.com/aws/aws-sdk-go/service/dynamodb"
)

//Comments used for storing meta data
type Comments struct {
	svc       *dynamodb.DynamoDB
	DBName    string
	portfolio portfolio.Portfolio
	user      user.User
}

//New Comments
func New(svc *dynamodb.DynamoDB, dbName string, p portfolio.Portfolio, u user.User) *Comments {
	return &Comments{svc: svc, DBName: dbName, portfolio: p, user: u}
}

//Comment comment date for get operation
type Comment struct {
	UserID             string
	Username           string
	UserProfilePicture string
	Content            string
	CreatedAt          string
}

//AddInput input to add portfolio item
type AddInput struct {
	ProjectID string
	Content   string
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
	Items []Comment
	Err   string
}
