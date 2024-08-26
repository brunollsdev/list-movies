package usecases

import (
	"crypto/md5"
	"encoding/hex"
	"errors"

	"github.com/brunollsdev/like-movies/application/dto"
	"github.com/brunollsdev/like-movies/application/repositories"
	"github.com/brunollsdev/like-movies/core/domain"
)

type UserUsecase struct {
	UserRepo repositories.UserRepositoryI
}

func (uc *UserUsecase) RegisterUser(payload dto.UserCreateRequest) (*domain.User, error) {
	user := domain.NewUser()

	user.Name = payload.Name
	user.Username = payload.Username
	user.Password = cryptPassowrd(payload.Password)

	err := user.Validate()

	if err != nil {
		return nil, err
	}

	err = uc.UserRepo.Insert(user)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (uc *UserUsecase) AuthenticateUser(username string, password string) (*domain.User, error) {

	user, err := uc.UserRepo.FindByUsername(username)

	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.New("user not found")
	}

	if user.Password != cryptPassowrd(password) {
		return nil, errors.New("password incorrect")
	}

	return user, nil
}

func cryptPassowrd(password string) string {
	hash := md5.New()
	pass := []byte(password)

	return hex.EncodeToString(hash.Sum(pass))

}
