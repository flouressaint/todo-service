package repo

import (
	"fmt"
	"log"

	"github.com/flouressaint/todo-service/internal/entity"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TodoRepo struct {
	db *gorm.DB
}

func NewTodoRepo(db *gorm.DB) *TodoRepo {
	// migrate the database
	err := db.AutoMigrate(&entity.Todo{})
	if err != nil {
		log.Fatal("Failed to migrate Todo in the Database")
	}
	log.Println("Migration Todo complete")
	return &TodoRepo{db}
}

func (r *TodoRepo) CreateTodo(todo entity.Todo) (uuid.UUID, error) {
	if err := r.db.Create(&todo).Error; err != nil {
		return todo.Id, err
	}
	return todo.Id, nil
}

func (r *TodoRepo) GetTodos(userId uuid.UUID) ([]entity.Todo, error) {
	var todos []entity.Todo
	if err := r.db.Where("user_id = ?", userId).Find(&todos).Error; err != nil {
		return todos, err
	}
	return todos, nil
}

func (r *TodoRepo) DeleteTodo(id, userId uuid.UUID) error {
	var todo entity.Todo
	if err := r.db.First(&todo, id).Error; err != nil {
		return err
	}
	if todo.UserId != userId {
		return fmt.Errorf("todo does not belong to user")
	}
	if err := r.db.Delete(&todo).Error; err != nil {
		return err
	}
	return nil
}

func (r *TodoRepo) UpdateTodo(id uuid.UUID, newTodo entity.Todo) error {
	var todo entity.Todo
	if err := r.db.First(&todo, id).Error; err != nil {
		return err
	}
	if todo.UserId != newTodo.UserId {
		return fmt.Errorf("todo does not belong to user")
	}
	if err := r.db.Model(&todo).Updates(newTodo).Error; err != nil {
		return err
	}
	return nil
}
