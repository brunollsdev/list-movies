package main

import (
	"log"

	"github.com/brunollsdev/like-movies/infra/database"
	"github.com/brunollsdev/like-movies/infra/server"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
		panic(err)
	}

	db := database.NewDatabase()
	s := server.Server{Db: db.Connect()}
	s.Start()
}
