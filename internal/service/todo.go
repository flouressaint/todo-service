package service

import (
	"github.com/flouressaint/todo-service/internal/entity"
	"github.com/flouressaint/todo-service/internal/repo"
)

type TodoService struct {
	todoRepo repo.Todo
}

func NewTodoService(todoRepo repo.Todo) *TodoService {
	return &TodoService{
		todoRepo: todoRepo,
	}
}

func (s *TodoService) CreateTodo(todo entity.Todo) (int, error) {
	return s.todoRepo.CreateTodo(todo)
}

func (s *TodoService) GetTodos(userId int) ([]entity.Todo, error) {
	return s.todoRepo.GetTodos(userId)
}

func (s *TodoService) DeleteTodo(id, userId int) error {
	return s.todoRepo.DeleteTodo(id, userId)
}

func (s *TodoService) UpdateTodo(id int, newTodo entity.Todo) error {
	return s.todoRepo.UpdateTodo(id, newTodo)
}
