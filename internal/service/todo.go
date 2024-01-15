package service

import (
	"github.com/flouressaint/todo-service/internal/entity"
	"github.com/flouressaint/todo-service/internal/repo"
	"github.com/google/uuid"
)

type TodoService struct {
	todoRepo repo.Todo
}

func NewTodoService(todoRepo repo.Todo) *TodoService {
	return &TodoService{
		todoRepo: todoRepo,
	}
}

func (s *TodoService) CreateTodo(todo entity.Todo) (uuid.UUID, error) {
	return s.todoRepo.CreateTodo(todo)
}

func (s *TodoService) GetTodos(userId uuid.UUID) ([]entity.Todo, error) {
	return s.todoRepo.GetTodos(userId)
}

func (s *TodoService) DeleteTodo(id, userId uuid.UUID) error {
	return s.todoRepo.DeleteTodo(id, userId)
}

func (s *TodoService) UpdateTodo(id uuid.UUID, newTodo entity.Todo) error {
	return s.todoRepo.UpdateTodo(id, newTodo)
}
