package domain

import "github.com/asaskevich/govalidator"

type User struct {
	Id       int    `json:"id" valid:"-" gorm:"primaryKey"`
	Name     string `json:"name" valid:"required~Nome é um campo obrigatório" gorm:"not null"`
	Username string `json:"username" valid:"required~Username é um campo obrigatório" gorm:"unique,not null"`
	Password string `json:"password" valid:"required~Password é um campo obrigatório" gorm:"varchar, not null"`
}

func NewUser() *User {
	return &User{}
}

func (user *User) Validate() error {
	_, err := govalidator.ValidateStruct(user)

	if err != nil {
		return &ValidateError{
			Errors: govalidator.ErrorsByField(err),
		}
	}

	return nil
}
