package repositories_test

import (
	"testing"

	"github.com/brunollsdev/like-movies/application/repositories"
	"github.com/brunollsdev/like-movies/core/domain"
	"github.com/brunollsdev/like-movies/infra/database"
)

func TestInsertUser(t *testing.T) {
	db := database.NewDBTest()

	repo := repositories.UserRepository{Db: db}

	user := domain.NewUser()
	user.Name = "User Test"
	user.Username = "user_test"
	user.Password = "Teste@123"

	err := repo.Insert(user)

	if err != nil {
		t.Errorf("%s: %s", err.Error(), err)
	}

	if user.Id == 0 {
		t.Errorf("%s: %s", err.Error(), err)
	}

}

func TestFindByUsername(t *testing.T) {
	db := database.NewDBTest()

	repo := repositories.UserRepository{Db: db}

	user := domain.NewUser()
	user.Name = "User Test"
	user.Username = "user_test"
	user.Password = "Teste@123"

	repo.Insert(user)

	res, err := repo.FindByUsername(user.Username)

	if err != nil {
		t.Errorf("%s: %s", err.Error(), err)
	}

	if res.Id != user.Id {
		t.Error("Usuário encontrado diferente do usuário informado.")
	}
}
