package repo

import (
	"fmt"
	"log"

	"github.com/flouressaint/todo-service/internal/entity"
	"gorm.io/gorm"
)

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	// migrate the database
	err := db.AutoMigrate(&entity.User{})
	if err != nil {
		log.Fatal("Failed to migrate User in the Database")
	}
	log.Println("Migration User complete")
	return &UserRepo{db}
}

func (r *UserRepo) CreateUser(user entity.User) (int, error) {
	if err := r.db.Create(&user).Error; err != nil {
		return user.Id, fmt.Errorf("user with name %s already exists", user.Username)
	}
	return user.Id, nil
}
