package entity

import (
	"gorm.io/gorm"
	"time"
)

type Logbook struct {
	ID         uint           `json:"id" gorm:"primary_key"`
	UserID     uint           `json:"user_id"`
	User       User           `json:"user"`
	Judul      string         `json:"judul"`
	Topik      string         `json:"topik"`
	Pembimbing string         `json:"pembimbing"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `json:"-" gorm:"index,column:deleted_at"`
}
