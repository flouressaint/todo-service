package controller

import (
	"net/http"

	"github.com/flouressaint/todo-service/internal/entity"
	"github.com/flouressaint/todo-service/internal/service"
	"github.com/labstack/echo/v4"
)

type userRoutes struct {
	userService service.User
}

func newUserRoutes(g *echo.Group, userService service.User) {
	r := &userRoutes{
		userService: userService,
	}

	g.POST("", r.CreateUser)
}

func (r *userRoutes) CreateUser(c echo.Context) error {
	var input entity.User

	if err := c.Bind(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return err
	}

	if err := c.Validate(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return err
	}

	id, err := r.userService.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return err
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}
