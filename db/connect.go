package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {

	db, err := gorm.Open(
		postgres.New(
			postgres.Config{
				DSN:                  "host=localhost user=postgres password=admin123 dbname=portal-vendas-db port=5432 sslmode=disable TimeZone=America/Sao_Paulo",
				PreferSimpleProtocol: true, // disables implicit prepared statement usage
			},
		), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
