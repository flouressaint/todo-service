package app

import (
	"log"

	"github.com/flouressaint/todo-service/config"
	"github.com/flouressaint/todo-service/internal/endpoint"
	"github.com/flouressaint/todo-service/internal/repository"
	"github.com/flouressaint/todo-service/internal/service"
	"github.com/labstack/echo/v4"
)

type App struct {
	e    *endpoint.Endpoint
	s    *service.Service
	r    *repository.Repository
	echo *echo.Echo
}

func New() (*App, error) {
	a := &App{}

	// config
	conf, err := config.New(".")
	if err != nil {
		log.Fatal("? Failed to load environment variables", err)
	}

	a.r = repository.New(conf)
	a.s = service.New(a.r)
	a.e = endpoint.New(a.s)

	a.echo = echo.New()
	a.echo.POST("/user", a.e.CreateUser)
	a.echo.GET("/todo", a.e.GetTodos)
	a.echo.POST("/todo", a.e.CreateTodo)
	a.echo.DELETE("/todo/:id", a.e.DeleteTodo)
	a.echo.PUT("/todo/:id", a.e.UpdateTodo)
	return a, nil
}

func (a *App) Run() error {
	return a.echo.Start(":8080")
}
