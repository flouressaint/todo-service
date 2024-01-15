package controller

import (
	"net/http"

	"github.com/flouressaint/todo-service/internal/entity"
	"github.com/flouressaint/todo-service/internal/service"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type todoRoutes struct {
	todoService service.Todo
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

type todoInput struct {
	Title     string `json:"title" validate:"required"`
	Completed bool   `json:"completed"`
}

// @Summary Create todo
// @Description Create todo
// @Tags todos
// @Accept json
// @Produce json
// @Param input body todoInput true "input"
// @Success 201 {object} controller.todoRoutes.CreateTodo.response
// @Failure 400 {object} echo.HTTPError
// @Failure 500 {object} echo.HTTPError
// @Security BearerAuth
// @Router /api/todo [post]
func (r *todoRoutes) CreateTodo(c echo.Context) error {
	var input todoInput
	if err := c.Bind(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return err
	}
	if err := c.Validate(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return err
	}
	data := entity.Todo{
		UserId:    c.Get("userId").(uuid.UUID),
		Title:     input.Title,
		Completed: input.Completed,
	}

	id, err := r.todoService.CreateTodo(data)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return err
	}

	type response struct {
		Id uuid.UUID `json:"id"`
	}

	return c.JSON(http.StatusCreated, response{
		Id: id,
	})
}

// @Summary Get todos
// @Description Get todos
// @Tags todos
// @Accept json
// @Produce json
// @Success 200 {array} entity.Todo
// @Failure 500 {object} echo.HTTPError
// @Security BearerAuth
// @Router /api/todo [get]
func (r *todoRoutes) GetTodos(c echo.Context) error {
	userId := c.Get("userId").(uuid.UUID)
	todos, err := r.todoService.GetTodos(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return err
	}

	return c.JSON(http.StatusOK, todos)
}

// @Summary Delete todo
// @Description Delete todo
// @Tags todos
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} controller.todoRoutes.DeleteTodo.response
// @Failure 400 {object} echo.HTTPError
// @Failure 500 {object} echo.HTTPError
// @Security BearerAuth
// @Router /api/todo/{id} [delete]
func (r *todoRoutes) DeleteTodo(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return err
	}
	userId := c.Get("userId").(uuid.UUID)

	err = r.todoService.DeleteTodo(id, userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return err
	}

	type response struct {
		Status string `json:"status"`
	}

	return c.JSON(http.StatusOK, response{
		Status: "ok",
	})
}

// @Summary Update todo
// @Description Update todo
// @Tags todos
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param input body todoInput true "input"
// @Success 200 {object} controller.todoRoutes.UpdateTodo.response
// @Failure 400 {object} echo.HTTPError
// @Failure 500 {object} echo.HTTPError
// @Security BearerAuth
// @Router /api/todo/{id} [put]
func (r *todoRoutes) UpdateTodo(c echo.Context) error {
	var input todoInput
	if err := c.Bind(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return err
	}
	if err := c.Validate(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return err
	}

	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return err
	}

	data := entity.Todo{
		UserId:    c.Get("userId").(uuid.UUID),
		Title:     input.Title,
		Completed: input.Completed,
	}

	err = r.todoService.UpdateTodo(id, data)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return err
	}

	type response struct {
		Status string `json:"status"`
	}

	return c.JSON(http.StatusOK, response{
		Status: "ok",
	})
}
