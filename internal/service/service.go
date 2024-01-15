package service

import (
	"github.com/flouressaint/todo-service/internal/entity"
	"github.com/flouressaint/todo-service/internal/repo"
	"github.com/google/uuid"
)

type Auth interface {
	ParseToken(accessToken string) (uuid.UUID, error)
}

type User interface {
	CreateUser(user entity.User) (int, error)
}

type Todo interface {
	CreateTodo(todo entity.Todo) (uuid.UUID, error)
	GetTodos(userId uuid.UUID) ([]entity.Todo, error)
	DeleteTodo(id, userId uuid.UUID) error
	UpdateTodo(id uuid.UUID, newTodo entity.Todo) error
}

type Services struct {
	Auth Auth
	User User
	Todo Todo
}

type ServicesDependencies struct {
	Repo    *repo.Repositories
	SignKey string
}

func NewServices(deps ServicesDependencies) *Services {
	return &Services{
		Auth: NewAuthService(deps.Repo.User, deps.SignKey),
		User: NewUserService(deps.Repo.User),
		Todo: NewTodoService(deps.Repo.Todo),
	}
}
