package controller

import (
	"net/http"
	"strconv"

	"github.com/flouressaint/todo-service/internal/entity"
	"github.com/flouressaint/todo-service/internal/service"
	"github.com/labstack/echo/v4"
)

type todoRoutes struct {
	todoService service.Todo
}
type UserId struct {
	UserId int `json:"user_id" validate:"required"`
}

func newTodoRoutes(g *echo.Group, todoService service.Todo) {
	r := &todoRoutes{
		todoService: todoService,
	}

	g.POST("", r.CreateTodo)
	g.GET("", r.GetTodos)
	g.DELETE("/:id", r.DeleteTodo)
	g.PUT("/:id", r.UpdateTodo)
}

func (r *todoRoutes) CreateTodo(c echo.Context) error {
	var input entity.Todo
	if err := c.Bind(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return err
	}
	if err := c.Validate(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return err
	}

	id, err := r.todoService.CreateTodo(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return err
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (r *todoRoutes) GetTodos(c echo.Context) error {
	var input UserId
	if err := c.Bind(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return err
	}
	if err := c.Validate(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return err
	}

	todos, err := r.todoService.GetTodos(input.UserId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return err
	}

	return c.JSON(http.StatusOK, todos)
}

func (r *todoRoutes) DeleteTodo(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return err
	}

	var input UserId
	if err := c.Bind(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return err
	}
	if err := c.Validate(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return err
	}

	err = r.todoService.DeleteTodo(id, input.UserId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return err
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}

func (r *todoRoutes) UpdateTodo(c echo.Context) error {
	var input entity.Todo
	if err := c.Bind(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return err
	}
	if err := c.Validate(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return err
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return err
	}

	err = r.todoService.UpdateTodo(id, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return err
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}
