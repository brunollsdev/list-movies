package domain

import (
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/brunollsdev/like-movies/application/facades"
	"github.com/brunollsdev/like-movies/infra/utils"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type Movie struct {
	Id        int       `json:"id" valid:"required~id é um campo obrigatório,movieExist~Filme não encontrado,type(int)" gorm:"primaryKey"`
	Rating    int       `json:"rating" valid:"range(0|10)~A avaliação pode ser entre 0 e 10,optional" gorm:"type:int"`
	Comment   string    `json:"comment" valid:"-" gorm:"type(varchar)"`
	UserId    int       `json:"user_id" valid:"required~É necessário informar um usuário" gorm:"type:varchar;not null;foreignKey:FkUserId;references:UserId"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime:milli"`
}

type Error interface {
	Error() string
}

type ValidateError struct {
	Errors map[string]string `json:"errors"`
}

func (e *ValidateError) Error() string {
	return "Validate fields error"
}

func NewMovie() *Movie {
	return &Movie{}
}

func (mv *Movie) Validate() error {
	govalidator.CustomTypeTagMap.Set("movieExist", func(i interface{}, context interface{}) bool {
		switch v := context.(type) {
		case Movie:
			return checkIdExist(v.Id)
		}

		return false
	})
	_, err := govalidator.ValidateStruct(mv)

	logId := uuid.NewString()
	if err != nil {
		utils.Log().WithFields(logrus.Fields{
			"id":     logId,
			"layer":  "domain",
			"method": "domain.movie.validate",
			"data":   err,
		}).Error(err.Error())

		return &ValidateError{
			govalidator.ErrorsByField(err),
		}
	}

	return nil
}

func (mv *Movie) TableName() string {
	return "favorite_movies"
}

func checkIdExist(id int) bool {
	mf := facades.NewMoviesFacade()

	err := mf.FindById(id)

	return err == nil
}
