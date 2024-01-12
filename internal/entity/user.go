package entity

import "time"

type User struct {
	//Id        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
<<<<<<< Updated upstream
	Id        int    `gorm:"primaryKey"`
	Username  string `json:"username" validate:"required" gorm:"unique;type:varchar(255);not null"`
=======
	Id        int    `json:"id" gorm:"primaryKey"`
	Username  string `json:"username" validate:"required" gorm:"type:varchar(255);not null"`
>>>>>>> Stashed changes
	Password  string `json:"password" validate:"required" gorm:"type:varchar(255);not null"`
	CreatedAt time.Time
}
