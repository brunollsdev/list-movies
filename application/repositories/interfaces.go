package repositories

import "github.com/brunollsdev/like-movies/core/domain"

const favorite_movies_table = "favorite_movies"
const user_table = "users"

type MovieRepositoryI interface {
	Insert(movie *domain.Movie) error
	Get(userId int) ([]domain.Movie, error)
	Delete(movieId int, userId int) error
	Update(movie *domain.Movie) error
}

type UserRepositoryI interface {
	Insert(user *domain.User) error
	FindByUsername(username string) (*domain.User, error)
}
