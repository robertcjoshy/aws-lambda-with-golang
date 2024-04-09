package config

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "database-robert.chsu6q8eqfk8.ap-south-1.rds.amazonaws.com"
	port     = 5432
	user     = "postgres"
	password = "daisyjoshy"
)

func connect_datbase() {
	var err error

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s", host, port, user, password)

	Database, err = gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
	if err != nil {
		panic(err)
	} else {
		return MyResponse{Message: fmt.Sprintf("%s id %d years old", event.Name, event.Age)}, nil
	}
}
