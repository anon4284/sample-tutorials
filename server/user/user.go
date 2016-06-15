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
