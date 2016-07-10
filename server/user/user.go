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

//UserInfo used for GetAll operation
type UserInfo struct {
	UserID         string
	Username       string
	ProfilePicture string
}

type MeInfo struct {
	Useremail      string
	Username       string
	ProfilePicture string
}

//AddInput input required to add the user
type AddInput struct {
	Useremail      string
	Username       string
	Password       string
	PasswordRepeat string
}

type DeleteInput struct {
	UserID string
}

//Output output of add operation
type Output struct {
	Valid bool
	Msg   string
	Err   string
}

//ChangePasswordInput needed to change password
type ChangePasswordInput struct {
	Old       string
	New       string
	NewRepeat string
}

//GetMeOutput of get logged in users profile
type GetMeOutput struct {
	Useremail string
	Username  string
}

//GetAllOutput get all users output
type GetAllOutput struct {
	Valid bool
	Msg   string
	Users []UserInfo
	Err   string
}

//LoginInput used to login to app
type LoginInput struct {
	Useremail string
	Password  string
}

//LoginOutput output returned from login
type LoginOutput struct {
	Valid   bool
	IsAdmin bool
	Token   string
	UserID  string
	Msg     string
}
