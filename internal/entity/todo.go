package entity

import "github.com/google/uuid"

type Todo struct {
	Id        uuid.UUID `json:"id" gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	UserId    uuid.UUID `json:"user_id" gorm:"type:uuid;default:gen_random_uuid();not null"`
	Title     string    `json:"title" validate:"required" gorm:"type:varchar(255);not null"`
	Completed bool      `json:"completed" gorm:"type:bool;"`
}
