package database

import (
	"github.com/brunollsdev/like-movies/core/domain"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewDBTest() *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&domain.User{}, &domain.Movie{})

	return db
}
