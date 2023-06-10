package request

import (
	"time"
)

type KaryaRequest struct {
	TanggalSubmit  time.Time `json:"tanggal_submit" form:"tanggal_submit" gorm:"type:timestamp;default:now()"`
	TanggalPublish time.Time `json:"tanggal_publish" form:"tanggal_publish" gorm:"type:timestamp;default:now()"`
	Link           string    `json:"link" form:"link" `
}
