package controller

import (
	_ "github.com/flouressaint/todo-service/docs"
	"github.com/flouressaint/todo-service/internal/service"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func NewRouter(handler *echo.Echo, services *service.Services) {
	handler.GET("/health", func(c echo.Context) error { return c.NoContent(200) })
	handler.GET("/swagger/*", echoSwagger.WrapHandler)

	newUserRoutes(handler.Group("/user"), services.User)

	authMiddleware := &AuthMiddleware{services.Auth}
	api := handler.Group("/api", authMiddleware.UserIdentity)
	{
		newTodoRoutes(api.Group("/todo"), services.Todo)
	}

}
