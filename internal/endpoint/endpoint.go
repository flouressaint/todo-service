package endpoint

import (
	"net/http"
	"strconv"

	"github.com/flouressaint/todo-service/internal/entity"
	"github.com/flouressaint/todo-service/internal/service"
	"github.com/labstack/echo/v4"
)

type Endpoint struct {
	s *service.Service
}

func New(s *service.Service) *Endpoint {
	return &Endpoint{s: s}
}

func (e *Endpoint) CreateUser(c echo.Context) error {
	var input entity.User

	if err := c.Bind(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return err
	}

	if err := c.Validate(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return err
	}

	id, err := e.s.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return err
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (e *Endpoint) CreateTodo(c echo.Context) error {
	var input entity.Todo
	if err := c.Bind(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return err
	}

	if err := c.Validate(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return err
	}

	id, err := e.s.CreateTodo(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return err
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (e *Endpoint) GetTodos(c echo.Context) error {
	type UserId struct {
		UserId int `json:"user_id" binding:"required"`
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

	todos, err := e.s.GetTodos(input.UserId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return err
	}

	return c.JSON(http.StatusOK, todos)
}

func (e *Endpoint) DeleteTodo(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return err
	}

	err = e.s.DeleteTodo(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return err
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}

func (e *Endpoint) UpdateTodo(c echo.Context) error {
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

	err = e.s.UpdateTodo(id, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return err
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}
