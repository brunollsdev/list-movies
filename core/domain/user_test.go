package domain_test

import (
	"testing"

	"github.com/brunollsdev/like-movies/core/domain"
)

func TestUserValidate(t *testing.T) {

	user := domain.NewUser()
	user.Name = "User Test"
	user.Username = "user_test"
	user.Password = "Teste@123"

	err := user.Validate()

	if err != nil {
		t.Errorf("%s: %s", err.Error(), err)
	}
}
