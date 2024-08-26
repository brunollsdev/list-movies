package server

import (
	"github.com/brunollsdev/like-movies/application/controllers"
	"github.com/brunollsdev/like-movies/application/facades"
	"github.com/brunollsdev/like-movies/application/middleware"
	"github.com/brunollsdev/like-movies/core/domain"
	_ "github.com/brunollsdev/like-movies/docs"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	"gorm.io/gorm"
)

type Server struct {
	Db   *gorm.DB
	User *domain.User
}

func (server *Server) Start() {
	e := echo.New()

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	server.moviesAPI(e)
	server.userAPI(e)

	e.Logger.Fatal(e.Start(":5000"))
}

func (server *Server) moviesAPI(e *echo.Echo) {
	gp := e.Group("api/v1/movies")

	gp.Use(middleware.AuthMiddleware(server.Db))

	ctrl := controllers.NewMovieController()

	ctrl.MovieFacade = facades.NewMoviesFacade()
	ctrl.Db = server.Db

	gp.POST("/favorite", ctrl.Favorite)
	gp.GET("/list", ctrl.List)
	gp.DELETE("/:id/remove", ctrl.Remove)
	gp.PUT("/:id/update", ctrl.Update)
	gp.GET("/search", ctrl.Search)
	gp.GET("/:id/details", ctrl.Details)
}

func (server *Server) userAPI(e *echo.Echo) {
	gp := e.Group("api/v1/user")

	ctrl := controllers.NewUserController()
	ctrl.Db = server.Db

	gp.POST("/create", ctrl.Register)
}
