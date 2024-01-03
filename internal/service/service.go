package service

import (
	"github.com/flouressaint/todo-service/internal/entity"
	"github.com/flouressaint/todo-service/internal/repo"
)

type User interface {
	CreateUser(user entity.User) (int, error)
}

type Todo interface {
	CreateTodo(todo entity.Todo) (int, error)
	GetTodos(userId int) ([]entity.Todo, error)
	DeleteTodo(id, userId int) error
	UpdateTodo(id int, newTodo entity.Todo) error
}

type Services struct {
	User User
	Todo Todo
}

type ServicesDependencies struct {
	Repo *repo.Repositories
}

func NewServices(deps ServicesDependencies) *Services {
	return &Services{
		User: NewUserService(deps.Repo.User),
		Todo: NewTodoService(deps.Repo.Todo),
	}
}
