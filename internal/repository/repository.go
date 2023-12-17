package repository

import (
	"fmt"
	"github.com/flouressaint/todo-service/config"
	"github.com/flouressaint/todo-service/internal/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
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
	err = db.AutoMigrate(&entity.User{})
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
