package controller

import (
	"github.com/flouressaint/todo-service/internal/service"
	"github.com/labstack/echo/v4"
)

func NewRouter(handler *echo.Echo, services *service.Services) {
	newUserRoutes(handler.Group("/user"), services.User)
	newTodoRoutes(handler.Group("/todo"), services.Todo)
}
