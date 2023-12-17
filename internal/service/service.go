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
