package usecases_test

import (
	"testing"

	"github.com/brunollsdev/like-movies/application/dto"
	"github.com/brunollsdev/like-movies/application/repositories"
	"github.com/brunollsdev/like-movies/core/usecases"
	"github.com/brunollsdev/like-movies/infra/database"
)

func prepareUser() repositories.UserRepository {
	db := database.NewDBTest()

	userRepo := repositories.UserRepository{Db: db}

	return userRepo
}

func TestRegisterUser(t *testing.T) {
	userRepo := prepareUser()

	userUseCase := usecases.UserUsecase{
		UserRepo: userRepo,
	}

	payload := dto.UserCreateRequest{
		Name:     "User Test",
		Username: "user_test",
		Password: "Teste@123",
	}

	user, err := userUseCase.RegisterUser(payload)

	if err != nil {
		t.Errorf("%s: %s", err.Error(), err)
	}

	if user.Password == payload.Password {
		t.Errorf("Não foi feito o hash da senha")
	}
}

func TestAuthenticateUser(t *testing.T) {
	userRepo := prepareUser()

	userUseCase := usecases.UserUsecase{
		UserRepo: userRepo,
	}

	payload := dto.UserCreateRequest{
		Name:     "User Test",
		Username: "user_test",
		Password: "Teste@123",
	}

	userUseCase.RegisterUser(payload)

	user, err := userUseCase.AuthenticateUser("user_test", "Teste@123")

	if err != nil {
		t.Errorf("%s: %s", err.Error(), err)
	}

	if user.Id == 0 {
		t.Errorf("Usuário não encontrado")
	}
}
