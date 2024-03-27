package db

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {

	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	db, err := gorm.Open(
		postgres.New(
			postgres.Config{
				DSN:                  fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=America/Sao_Paulo", host, user, password, dbname, port),
				PreferSimpleProtocol: true, // disables implicit prepared statement usage
			},
		), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
