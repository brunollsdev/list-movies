package usecases

import (
	"time"

	"github.com/brunollsdev/like-movies/application/dto"
	"github.com/brunollsdev/like-movies/application/repositories"
	"github.com/brunollsdev/like-movies/core/domain"
)

type MovieUsecase struct {
	MovieRepo repositories.MovieRepositoryI
	UserRepo  repositories.UserRepositoryI
}

func (uc *MovieUsecase) FavoriteMovie(payload *dto.FavoriteMovieRequest, useId int) (*domain.Movie, error) {
	movie := domain.NewMovie()

	movie.Id = payload.Id
	movie.Comment = payload.Comment
	movie.Rating = payload.Rating
	movie.UserId = useId

	err := movie.Validate()

	if err != nil {
		return nil, err
	}

	err = uc.MovieRepo.Insert(movie)

	if err != nil {
		return nil, err
	}

	return movie, nil
}

func (uc *MovieUsecase) AllFavoriteMovie(userId int) ([]domain.Movie, error) {
	movies, err := uc.MovieRepo.Get(userId)

	if err != nil {
		return nil, err
	}

	return movies, nil
}

func (uc *MovieUsecase) RemoveFavoriteMovie(movieId int, userId int) error {
	return uc.MovieRepo.Delete(movieId, userId)
}

func (uc *MovieUsecase) UpdateMovie(payload dto.MovieDto, userId int) (*domain.Movie, error) {
	movie := domain.NewMovie()

	movie.Id = payload.Id
	movie.Comment = payload.Comment
	movie.Rating = payload.Rating
	movie.UserId = userId
	movie.UpdatedAt = time.Now()

	err := movie.Validate()

	if err != nil {
		return nil, err
	}

	return movie, uc.MovieRepo.Update(movie)
}
