package app

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/flouressaint/todo-service/config"
	"github.com/flouressaint/todo-service/internal/controller"
	"github.com/flouressaint/todo-service/internal/repo"
	"github.com/flouressaint/todo-service/internal/service"
	"github.com/flouressaint/todo-service/pkg/httpserver"
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

	// HTTP server
	log.Printf("Starting http server...")
	log.Printf("Server port: %s", cfg.ServerPort)
	httpServer := httpserver.New(handler, httpserver.Port(cfg.ServerPort))

	// Waiting signal
	log.Printf("Configuring graceful shutdown...")
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		log.Printf("app - Run - signal: " + s.String())
	case err = <-httpServer.Notify():
		log.Fatal("app - Run - httpServer.Notify: %w", err)
	}

	// Graceful shutdown
	log.Printf("Shutting down...")
	err = httpServer.Shutdown()
	if err != nil {
		log.Fatal("app - Run - httpServer.Shutdown: %w", err)
	}
}
