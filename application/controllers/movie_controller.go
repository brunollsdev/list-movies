package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/brunollsdev/like-movies/application/dto"
	"github.com/brunollsdev/like-movies/application/facades"
	"github.com/brunollsdev/like-movies/application/repositories"
	"github.com/brunollsdev/like-movies/core/domain"
	"github.com/brunollsdev/like-movies/core/usecases"
	"github.com/brunollsdev/like-movies/infra/utils"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type MovieController struct {
	Db          *gorm.DB
	MovieFacade facades.MoviesFacadeI
}

func NewMovieController() *MovieController {
	return &MovieController{}
}

func (c *MovieController) Favorite(ctx echo.Context) error {
	user := ctx.Get("user").(*domain.User)
	var payload *dto.FavoriteMovieRequest

	ctx.Bind(&payload)

	usecase := usecases.MovieUsecase{
		MovieRepo: &repositories.MovieRepository{Db: c.Db},
		UserRepo:  &repositories.UserRepository{Db: c.Db},
	}

	movie, err := usecase.FavoriteMovie(payload, user.Id)

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, utils.Response(false, "Erro ao adicionar filme na lista.", err))
	}

	return ctx.JSON(http.StatusOK, utils.Response(true, "Filme adicionado à lista com sucesso.", movie))
}

func (c *MovieController) List(ctx echo.Context) error {
	user := ctx.Get("user").(*domain.User)
	usecase := usecases.MovieUsecase{
		MovieRepo: &repositories.MovieRepository{
			Db: c.Db,
		},
	}

	res, err := usecase.AllFavoriteMovie(user.Id)

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, utils.Response(false, "Erro ao buscar lista de favoritos.", err))
	}

	return ctx.JSON(http.StatusOK, utils.Response(true, "Busca realizada com sucesso.", res))
}

func (c *MovieController) Remove(ctx echo.Context) error {
	user := ctx.Get("user").(*domain.User)
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		logId := uuid.NewString()
		utils.Log().WithFields(logrus.Fields{
			"id":     logId,
			"layer":  "controllers",
			"method": "controller.movie_controller.remove",
			"data":   err,
		}).Error(err.Error())

		return ctx.JSON(http.StatusInternalServerError, utils.Response(false, "Erro ao remover filme da lista.", &utils.LogDesc{
			LogId:     logId,
			EventDate: time.Now(),
		}))
	}

	usecase := usecases.MovieUsecase{
		MovieRepo: &repositories.MovieRepository{
			Db: c.Db,
		},
	}

	err = usecase.RemoveFavoriteMovie(id, user.Id)

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, utils.Response(false, "Erro ao remover filme da lista.", err))
	}

	return ctx.JSON(http.StatusOK, utils.Response(true, "Filme removido da lista com sucesso.", nil))
}

func (c *MovieController) Update(ctx echo.Context) error {
	user := ctx.Get("user").(*domain.User)
	var payload dto.MovieDto

	ctx.Bind(&payload)
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, utils.Response(false, "Erro ao atualizar dados do filme.", err))
	}

	payload.Id = id

	usecase := usecases.MovieUsecase{
		MovieRepo: &repositories.MovieRepository{
			Db: c.Db,
		},
	}

	movie, err := usecase.UpdateMovie(payload, user.Id)

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, utils.Response(false, "Erro ao atualizar filme da lista.", err))
	}

	return ctx.JSON(http.StatusOK, utils.Response(true, "Filme atualizado com sucesso.", movie))
}

func (c *MovieController) Search(ctx echo.Context) error {
	res, err := c.MovieFacade.Search(ctx.QueryParam("name"))

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, utils.Response(false, "Erro ao realizar busca.", err))
	}

	return ctx.JSON(http.StatusOK, utils.Response(true, "Busca realizada com sucesso.", res))
}

func (c *MovieController) Details(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("id"))

	res, err := c.MovieFacade.Description(id)

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, utils.Response(false, "Erro ao buscar detalhes do filme.", err))
	}

	return ctx.JSON(http.StatusOK, utils.Response(true, "Busca realizada com sucesso.", res))
}
