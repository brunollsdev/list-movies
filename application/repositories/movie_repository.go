package repositories

import (
	"fmt"
	"time"

	"github.com/brunollsdev/like-movies/core/domain"
	"github.com/brunollsdev/like-movies/infra/utils"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type MovieRepository struct {
	Db *gorm.DB
}

func (repo *MovieRepository) Insert(movie *domain.Movie) error {
	res := repo.Db.Table(favorite_movies_table).Select("id", "comment", "rating", "user_id", "created_at", "updated_at").Create(&movie)

	logId := uuid.NewString()
	if res.Error != nil {
		utils.Log().WithFields(logrus.Fields{
			"id":     logId,
			"layer":  "repository",
			"method": "repositories.movie_repository.insert",
			"data":   res.Error,
		}).Error(res.Error.Error())

		return &utils.LogDesc{
			LogId:     logId,
			EventDate: time.Now(),
		}
	}

	return nil
}

func (repo *MovieRepository) Get(userId int) ([]domain.Movie, error) {
	var movies []domain.Movie
	res := repo.Db.Raw(fmt.Sprintf("SELECT * FROM %s WHERE user_id = ?", favorite_movies_table), userId).Scan(&movies)

	logId := uuid.NewString()
	if res.Error != nil {
		utils.Log().WithFields(logrus.Fields{
			"id":     logId,
			"layer":  "repository",
			"method": "repositories.movie_repository.get",
			"data":   res.Error,
		}).Error(res.Error.Error())

		return nil, &utils.LogDesc{
			LogId:     logId,
			EventDate: time.Now(),
		}
	}

	return movies, nil
}

func (repo *MovieRepository) Delete(movieId int, userId int) error {
	res := repo.Db.Exec(fmt.Sprintf("DELETE FROM %s WHERE id = ? AND user_id = ?", favorite_movies_table), movieId, userId)

	logId := uuid.NewString()
	if res.Error != nil {
		utils.Log().WithFields(logrus.Fields{
			"id":     logId,
			"layer":  "repository",
			"method": "repositories.movie_repository.delete",
			"data":   res.Error,
		}).Error(res.Error.Error())

		return &utils.LogDesc{
			LogId:     logId,
			EventDate: time.Now(),
		}
	}

	return nil
}

func (repo *MovieRepository) Update(movie *domain.Movie) error {
	res := repo.Db.Table(favorite_movies_table).Omit("created_at").Select(
		"id",
		"comment",
		"rating",
		"user_id",
		"updated_at",
		"created_at",
	).Where("id = ? AND user_id = ?", movie.Id, movie.UserId).Save(&movie)

	logId := uuid.NewString()
	if res.Error != nil {
		utils.Log().WithFields(logrus.Fields{
			"id":     logId,
			"layer":  "repository",
			"method": "repositories.movie_repository.update",
			"data":   res.Error,
		}).Error(res.Error.Error())

		return &utils.LogDesc{
			LogId:     logId,
			EventDate: time.Now(),
		}
	}

	return nil
}
