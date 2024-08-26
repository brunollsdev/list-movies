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

type UserRepository struct {
	Db *gorm.DB
}

func (repo UserRepository) Insert(user *domain.User) error {
	res := repo.Db.Table(user_table).Select("id", "name", "username", "password", "created_at", "updated_at").Create(&user)

	logId := uuid.NewString()
	if res.Error != nil {
		utils.Log().WithFields(logrus.Fields{
			"id":     logId,
			"layer":  "repository",
			"method": "repositories.user_repository.insert",
			"data":   res.Error,
		}).Error(res.Error.Error())

		return &utils.LogDesc{
			LogId:     logId,
			EventDate: time.Now(),
		}
	}

	return nil
}

func (repo UserRepository) FindByUsername(username string) (*domain.User, error) {
	var user *domain.User
	res := repo.Db.Raw(fmt.Sprintf("SELECT * FROM %s WHERE username = ?", user_table), username).Scan(&user)

	logId := uuid.NewString()
	if res.Error != nil {
		utils.Log().WithFields(logrus.Fields{
			"id":     logId,
			"layer":  "repository",
			"method": "repositories.user_repository.find_by_username",
			"data":   res.Error,
		}).Error(res.Error.Error())

		return nil, &utils.LogDesc{
			LogId:     logId,
			EventDate: time.Now(),
		}
	}
	return user, nil
}
