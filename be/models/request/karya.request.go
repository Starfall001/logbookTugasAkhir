package request

import "time"

type KaryaRequest struct {
	TanggalSubmit  time.Time `json:"tanggal_submit" form:"tanggal_submit" validate:"required"`
	TanggalPublish time.Time `json:"tanggal_publish" form:"tanggal_publish" validate:"required"`
	Link           string    `json:"link" form:"link" validate:"required,email"`
}
