package user

import "github.com/aws/aws-sdk-go/service/dynamodb"

//User the user logic struct
type User struct {
	svc         *dynamodb.DynamoDB
	DBName      string
	amountLimit int64
}

//New create new user
func New(svc *dynamodb.DynamoDB, dbName string) *User {
	return &User{svc: svc, DBName: dbName}
}

//LimitAmount limits amount of users
func (u *User) LimitAmount(limit int64) {
	u.amountLimit = limit
}

//AddInput input required to add the user
type AddInput struct {
	Useremail string
	Username  string
	Password  string
}

//AddOutput output of add operation
type AddOutput struct {
	Valid bool
	Msg   string
	Err   string
}

//LoginInput used to login to app
type LoginInput struct {
	Useremail string
	Password  string
}

//LoginOutput output returned from login
type LoginOutput struct {
	Valid  bool
	Token  string
	UserID string
	Msg    string
}
