package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	user     = "postgres"
	password = "PzLS-1XJd"
	host     = "localhost"
	dbname   = "go_auth"
	port     = 5432
)

func Connect() {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// open database
	_, err := gorm.Open(postgres.Open(psqlconn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}
