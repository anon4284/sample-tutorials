package db

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

//DB database
type DB struct {
	SVC dynamodb.DynamoDB
}

type config struct {
	UserID     string
	UserSecret string
	Region     string
}

//New Create new DB connection
func New(cfg *aws.Config) *DB {
	return &DB{SVC: *dynamodb.New(session.New(cfg))}
}

//NewConfig Create config for db
func NewConfig(userID string, userSecret string, region string) *aws.Config {
	cred := credentials.NewStaticCredentials(userID, userSecret, "")
	config := aws.NewConfig().WithRegion(region).WithCredentials(cred)
	return config
}

//NewConfigFromJSON Create config for db
func NewConfigFromJSON(path string) *aws.Config {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		log.Print(err)
	}
	cfg := config{}
	json.Unmarshal(b, &cfg)

	cred := credentials.NewStaticCredentials(cfg.UserID, cfg.UserSecret, "")
	config := aws.NewConfig().WithRegion(cfg.Region).WithCredentials(cred)
	return config
}
