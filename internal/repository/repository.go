package repository

import (
	"fmt"
	"log"

	"github.com/flouressaint/todo-service/config"
	"github.com/flouressaint/todo-service/internal/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

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

func (r *Repository) CreateUser(user entity.User) (int, error) {
	if err := r.db.Create(&user).Error; err != nil {
		return user.Id, err
	}
	return user.Id, nil
}

func (r *Repository) CreateTodo(todo entity.Todo) (int, error) {
	if err := r.db.Create(&todo).Error; err != nil {
		return todo.Id, err
	}
	return todo.Id, nil
}

func (r *Repository) GetTodo(id int) (entity.Todo, error) {
	var todo entity.Todo
	if err := r.db.First(&todo, id).Error; err != nil {
		return todo, err
	}
	return todo, nil
}

func (r *Repository) GetTodos(userId int) ([]entity.Todo, error) {
	var todos []entity.Todo
	if err := r.db.Find(&todos).Where("userId = ?", userId).Error; err != nil {
		return todos, err
	}
	return todos, nil
}

func (r *Repository) DeleteTodo(todo entity.Todo) error {
	if err := r.db.Delete(&todo).Error; err != nil {
		return err
	}
	return nil
}

func (r *Repository) UpdateTodo(todo, newTodo entity.Todo) error {
	if err := r.db.Model(&todo).Updates(newTodo).Error; err != nil {
		return err
	}
	return nil
}
