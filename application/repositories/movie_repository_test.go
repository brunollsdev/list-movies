package repositories_test

import (
	"testing"
	"time"

	"github.com/brunollsdev/like-movies/application/repositories"
	"github.com/brunollsdev/like-movies/core/domain"
	"github.com/brunollsdev/like-movies/infra/database"
)

func TestInsertMovie(t *testing.T) {
	db := database.NewDBTest()

	repo := repositories.MovieRepository{Db: db}

	movie := domain.NewMovie()
	movie.Id = 238
	movie.Comment = "Test"
	movie.Rating = 2
	movie.UserId = 1

	err := repo.Insert(movie)

	if err != nil {
		t.Errorf("%s: %s", err.Error(), err)
	}

	if movie.CreatedAt == time.Now() {
		t.Errorf("O filme não foi criado")
	}
}

func TestGetMovie(t *testing.T) {
	db := database.NewDBTest()

	repo := repositories.MovieRepository{Db: db}

	movie := domain.NewMovie()
	movie.Id = 238
	movie.Comment = "Test"
	movie.Rating = 2
	movie.UserId = 1

	repo.Insert(movie)

	res, err := repo.Get(movie.UserId)

	if err != nil {
		t.Errorf("%s: %s", err.Error(), err)
	}

	if len(res) <= 0 {
		t.Error("Nenhum filme foi encontrado")
	}
}

func TestDeleteMovie(t *testing.T) {
	db := database.NewDBTest()

	repo := repositories.MovieRepository{Db: db}

	movie := domain.NewMovie()
	movie.Id = 238
	movie.Comment = "Test"
	movie.Rating = 2
	movie.UserId = 1

	repo.Insert(movie)

	err := repo.Delete(238, 1)

	if err != nil {
		t.Errorf("%s: %s", err.Error(), err)
	}
}

func TestUpdateMovie(t *testing.T) {
	db := database.NewDBTest()

	repo := repositories.MovieRepository{Db: db}

	movie := domain.NewMovie()
	movie.Id = 238
	movie.Comment = "Test"
	movie.Rating = 2
	movie.UserId = 1

	repo.Insert(movie)

	updDate := movie.UpdatedAt

	movie.Id = 500

	err := repo.Update(movie)

	if err != nil {
		t.Errorf("%s: %s", err.Error(), err)
	}

	if movie.UpdatedAt == updDate {
		t.Error("Filme não foi atualizado")
	}
}
