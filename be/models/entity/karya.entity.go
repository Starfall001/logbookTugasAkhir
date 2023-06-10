package entity

import (
	"gorm.io/gorm"
	"logbook_ta/models/response"
	"time"
)

type Karya struct {
	ID             uint                  `json:"id" gorm:"primary_key"`
	UserID         uint                  `json:"user_id"`
	User           response.UserResponse `json:"user"`
	TanggalSubmit  time.Time             `json:"tanggal_submit" gorm:"type:timestamp;default:now()"`
	TanggalPublish time.Time             `json:"tanggal_publish" gorm:"type:timestamp;default:now()"`
	Link           string                `json:"link"`
	CreatedAt      time.Time             `json:"created_at"`
	UpdatedAt      time.Time             `json:"updated_at"`
	DeletedAt      gorm.DeletedAt        `json:"-" gorm:"index,column:deleted_at"`
}
