package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Database *gorm.DB

const (
	host     = "database-robert.chsu6q8eqfk8.ap-south-1.rds.amazonaws.com"
	port     = 5432
	user     = "postgres"
	password = "daisyjoshy"
)

type MyEvent struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type MyResponse struct {
	Message string `json:"answer"`
}

func HandleLambdaEvent(event MyEvent) (MyResponse, error) {
	var err error

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s", host, port, user, password)

	Database, err = gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
	if err != nil {
		panic(err)
	} else {
		return MyResponse{Message: fmt.Sprintf("%s id %d years old", event.Name, event.Age)}, nil
	}
}
func main() {
	lambda.Start(HandleLambdaEvent)
}
