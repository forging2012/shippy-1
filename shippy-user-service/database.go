package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func CreateConnection() (*gorm.DB, error) {

	// Get database details from environment variables
	host := "127.0.0.1"
	user := "postgres"
	DBName := "shippy"
	password := "123qwe"

	return gorm.Open(
		"postgres",
		fmt.Sprintf(
			"host=%s user=%s dbname=%s sslmode=disable password=%s",
			host, user, DBName, password,
		),
	)
}
