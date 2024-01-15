package repo

import (
	"fmt"
	"log"

	"github.com/flouressaint/todo-service/config"
	"github.com/flouressaint/todo-service/internal/entity"
	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User interface {
	CreateUser(user entity.User) (int, error)
}

type Todo interface {
	CreateTodo(todo entity.Todo) (uuid.UUID, error)
	GetTodos(userId uuid.UUID) ([]entity.Todo, error)
	DeleteTodo(id, userId uuid.UUID) error
	UpdateTodo(id uuid.UUID, newTodo entity.Todo) error
}

type Repositories struct {
	User
	Todo
}

func NewRepositories(config config.Config) *Repositories {
	// connect to the database
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		config.DBHost,
		config.DBUsername,
		config.DBPassword,
		config.DBName,
		config.DBPort,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the Database")
	}
	log.Println("Connected Successfully to the Database")

	return &Repositories{
		User: NewUserRepo(db),
		Todo: NewTodoRepo(db),
	}
}
