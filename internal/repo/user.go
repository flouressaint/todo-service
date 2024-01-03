package repo

import (
	"github.com/flouressaint/todo-service/internal/entity"
	"gorm.io/gorm"
)

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{db}
}

func (r *UserRepo) CreateUser(user entity.User) (int, error) {
	if err := r.db.Create(&user).Error; err != nil {
		return user.Id, err
	}
	return user.Id, nil
}
