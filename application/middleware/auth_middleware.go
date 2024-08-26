package middleware

import (
	"net/http"

	"github.com/brunollsdev/like-movies/application/repositories"
	"github.com/brunollsdev/like-movies/core/usecases"
	"github.com/brunollsdev/like-movies/infra/utils"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

var database *gorm.DB

func AuthMiddleware(db *gorm.DB) echo.MiddlewareFunc {
	database = db
	return middleware.BasicAuthWithConfig(middleware.BasicAuthConfig{
		Validator: auth,
	})
}

func auth(username string, password string, ctx echo.Context) (bool, error) {
	usecase := usecases.UserUsecase{
		UserRepo: repositories.UserRepository{
			Db: database,
		},
	}

	user, err := usecase.AuthenticateUser(username, password)

	if err != nil {
		return false, ctx.JSON(http.StatusUnauthorized, utils.Response(false, err.Error(), err))
	}

	ctx.Set("user", user)
	return true, nil
}
