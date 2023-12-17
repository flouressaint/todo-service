package service

import (
	"github.com/flouressaint/todo-service/internal/entity"
	"github.com/flouressaint/todo-service/internal/repository"
)

type Service struct {
	r *repository.Repository
}

func New(r *repository.Repository) *Service {
	return &Service{r: r}
}

func (s *Service) CreateUser(user entity.User) (int, error) {
	return s.r.CreateUser(user)
}

func (s *Service) CreateTodo(todo entity.Todo) (int, error) {
	return s.r.CreateTodo(todo)
}

func (s *Service) GetTodos(userId int) ([]entity.Todo, error) {
	return s.r.GetTodos(userId)
}

func (s *Service) DeleteTodo(id int) error {
	todo, err := s.r.GetTodo(id)
	if err != nil {
		return err
	}
	return s.r.DeleteTodo(todo)
}

func (s *Service) UpdateTodo(id int, newTodo entity.Todo) error {
	todo, err := s.r.GetTodo(id)
	if err != nil {
		return err
	}
	return s.r.UpdateTodo(todo, newTodo)
}
