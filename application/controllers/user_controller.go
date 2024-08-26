package controllers

import (
	"net/http"

	"github.com/brunollsdev/like-movies/application/dto"
	"github.com/brunollsdev/like-movies/application/repositories"
	"github.com/brunollsdev/like-movies/core/usecases"
	"github.com/brunollsdev/like-movies/infra/utils"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type UserController struct {
	Db *gorm.DB
}

func NewUserController() *UserController {
	return &UserController{}
}

func (c *UserController) Register(ctx echo.Context) error {
	var payload dto.UserCreateRequest

	ctx.Bind(&payload)

	usecase := usecases.UserUsecase{
		UserRepo: repositories.UserRepository{
			Db: c.Db,
		},
	}

	user, err := usecase.RegisterUser(payload)

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, utils.Response(false, "Erro ao cadastrar usuário.", err))
	}

	return ctx.JSON(http.StatusOK, utils.Response(true, "Usuário cadastrado com sucesso.", user))
}
