package repo

import (
	"fmt"
	"log"

	"github.com/flouressaint/todo-service/config"
	"github.com/flouressaint/todo-service/internal/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User interface {
	CreateUser(user entity.User) (int, error)
}

type Todo interface {
	CreateTodo(todo entity.Todo) (int, error)
	GetTodos(userId int) ([]entity.Todo, error)
	GetTodo(id int) (entity.Todo, error)
	DeleteTodo(user entity.Todo) error
	UpdateTodo(todo, newTodo entity.Todo) error
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
	fmt.Println("? Connected Successfully to the Database")

	// migrate the database
	err = db.AutoMigrate(&entity.User{}, &entity.Todo{})
	if err != nil {
		log.Fatal("Failed to migrate the Database")
	}
	fmt.Println("? Migration complete")
	return &Repositories{
		User: NewUserRepo(db),
		Todo: NewTodoRepo(db),
	}
}

type Repository struct {
	db *gorm.DB
}

func New(config config.Config) *Repository {
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
	fmt.Println("? Connected Successfully to the Database")

	// migrate the database
	err = db.AutoMigrate(&entity.User{}, &entity.Todo{})
	if err != nil {
		log.Fatal("Failed to migrate the Database")
	}
	fmt.Println("? Migration complete")

	return &Repository{db: db}
}
