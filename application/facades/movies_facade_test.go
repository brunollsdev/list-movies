package facades_test

import (
	"testing"

	"github.com/brunollsdev/like-movies/application/facades"
)

func TestFindById(t *testing.T) {
	t.Setenv("TMDB_API_KEY", "7c6fc6f5076d334bed929a0aeb6fc4cf")
	mf := facades.NewMoviesFacade()

	err := mf.FindById(238)

	if err != nil {
		t.Errorf("%s: %s", err.Error(), err)
	}
}

func TestSearch(t *testing.T) {
	t.Setenv("TMDB_API_KEY", "7c6fc6f5076d334bed929a0aeb6fc4cf")
	mf := facades.NewMoviesFacade()

	res, err := mf.Search("the godfa")

	if err != nil {
		t.Errorf("%s: %s", err.Error(), err)
	}

	if len(res) <= 0 {
		t.Errorf("Filme não foi buscado com sucesso")
	}
}
