package usecases_test

import (
	"testing"

	"github.com/brunollsdev/like-movies/application/dto"
	"github.com/brunollsdev/like-movies/application/repositories"
	"github.com/brunollsdev/like-movies/core/usecases"
	"github.com/brunollsdev/like-movies/infra/database"
)

func prepareMovie() (repositories.MovieRepository, repositories.UserRepository) {
	db := database.NewDBTest()

	movieRepo := repositories.MovieRepository{Db: db}
	userRepo := repositories.UserRepository{Db: db}

	return movieRepo, userRepo
}

func TestFavoriteMovie(t *testing.T) {
	t.Setenv("TMDB_API_KEY", "7c6fc6f5076d334bed929a0aeb6fc4cf")
	movieRepo, userRepo := prepareMovie()

	movieUseCase := usecases.MovieUsecase{MovieRepo: &movieRepo, UserRepo: userRepo}

	payload := dto.FavoriteMovieRequest{
		Id:      238,
		Comment: "",
		Rating:  0,
	}

	movie, err := movieUseCase.FavoriteMovie(&payload, 1)

	if err != nil {
		t.Errorf("%s: %s", err.Error(), err)
	}

	if movie.Id == 0 {
		t.Error("Filme não foi cadastrado com sucesso na lista.")
	}
}

func TestUpdateMovie(t *testing.T) {
	t.Setenv("TMDB_API_KEY", "7c6fc6f5076d334bed929a0aeb6fc4cf")
	movieRepo, _ := prepareMovie()

	movieUseCase := usecases.MovieUsecase{MovieRepo: &movieRepo}

	payload := dto.MovieDto{
		Id:      238,
		Comment: "Comment test",
		Rating:  5,
	}

	movie, err := movieUseCase.UpdateMovie(payload, 1)

	if err != nil {
		t.Errorf("%s: %s", err.Error(), err)
	}

	if movie.Id == 0 {
		t.Error("Filme não foi cadastrado com sucesso na lista.")
	}
}

func TestAllFavoriteMovie(t *testing.T) {
	movieRepo, _ := prepareMovie()

	movieUseCase := usecases.MovieUsecase{MovieRepo: &movieRepo}

	res, err := movieUseCase.AllFavoriteMovie(1)

	if err != nil {
		t.Errorf("%s: %s", err.Error(), err)
	}

	if res != nil {
		t.Error("Nenhum filme foi encontrado")
	}
}
