package app

import (
	"log"

	"github.com/flouressaint/todo-service/config"
	"github.com/flouressaint/todo-service/internal/controller"
	"github.com/flouressaint/todo-service/internal/repo"
	"github.com/flouressaint/todo-service/internal/service"
	"github.com/flouressaint/todo-service/pkg/validator"
	"github.com/labstack/echo/v4"
)

func Run() {
	// config
	cfg, err := config.New(".")
	if err != nil {
		log.Fatal("? Failed to load environment variables", err)
	}

	repositories := repo.NewRepositories(cfg)

	deps := service.ServicesDependencies{
		Repo: repositories,
	}
	services := service.NewServices(deps)

	handler := echo.New()
	// setup handler validator as lib validator
	handler.Validator = validator.NewCustomValidator()
	controller.NewRouter(handler, services)

	handler.Start(":" + cfg.ServerPort)
}
