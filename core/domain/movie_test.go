package domain_test

import (
	"testing"

	"github.com/brunollsdev/like-movies/core/domain"
)

func TestMovieValidate(t *testing.T) {
	t.Setenv("TMDB_API_KEY", "7c6fc6f5076d334bed929a0aeb6fc4cf")

	movie := domain.NewMovie()
	movie.Id = 238
	movie.Rating = 5
	movie.UserId = 1
	movie.Comment = ""

	err := movie.Validate()

	if err != nil {
		t.Errorf("%s: %s", err.Error(), err)
	}
}
