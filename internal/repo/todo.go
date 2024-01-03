package repo

import (
	"github.com/flouressaint/todo-service/internal/entity"
	"gorm.io/gorm"
)

type TodoRepo struct {
	db *gorm.DB
}

func NewTodoRepo(db *gorm.DB) *TodoRepo {
	return &TodoRepo{db}
}

func (r *TodoRepo) CreateTodo(todo entity.Todo) (int, error) {
	if err := r.db.Create(&todo).Error; err != nil {
		return todo.Id, err
	}
	return todo.Id, nil
}

func (r *TodoRepo) GetTodo(id int) (entity.Todo, error) {
	var todo entity.Todo
	if err := r.db.First(&todo, id).Error; err != nil {
		return todo, err
	}
	return todo, nil
}

func (r *TodoRepo) GetTodos(userId int) ([]entity.Todo, error) {
	var todos []entity.Todo
	if err := r.db.Find(&todos).Where("userId = ?", userId).Error; err != nil {
		return todos, err
	}
	return todos, nil
}

func (r *TodoRepo) DeleteTodo(todo entity.Todo) error {
	if err := r.db.Delete(&todo).Error; err != nil {
		return err
	}
	return nil
}

func (r *TodoRepo) UpdateTodo(todo, newTodo entity.Todo) error {
	if err := r.db.Model(&todo).Updates(newTodo).Error; err != nil {
		return err
	}
	return nil
}
