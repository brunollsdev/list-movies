package database

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Postgres struct {
	username string
	password string
	database string
	port     string
	host     string
}

func NewDatabase() *Postgres {
	return &Postgres{
		username: os.Getenv("DB_USERNAME"),
		password: os.Getenv("DB_PASSWORD"),
		database: os.Getenv("DB_DATABASE"),
		port:     os.Getenv("DB_PORT"),
		host:     os.Getenv("DB_HOST"),
	}
}

func (pt Postgres) Connect() *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=America/Sao_Paulo",
		pt.host,
		pt.username,
		pt.password,
		pt.database,
		pt.port,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(fmt.Sprintf("Error connect database: %s", err))
	}

	return db
}
