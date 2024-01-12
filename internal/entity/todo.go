package entity

type Todo struct {
	Id        int    `json:"id" gorm:"primaryKey"`
	UserId    int    `json:"user_id" validate:"required" gorm:"type:int;not null"`
	Title     string `json:"title" validate:"required" gorm:"type:varchar(255);not null"`
	Completed bool   `json:"completed" gorm:"type:bool;"`
}
