package entity

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        uint           `json:"id" gorm:"primary_key"`
	Name      string         `json:"name" `
	Username  string         `json:"username"`
	Email     string         `json:"email"`
	Password  string         `json:"-" gorm:"column:password"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index,column:deleted_at"`
}
