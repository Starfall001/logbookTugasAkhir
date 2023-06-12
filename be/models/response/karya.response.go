package response

import (
	"time"
)

type KaryaResponse struct {
	ID uint `json:"id"`

	TanggalSubmit  time.Time `json:"tanggal_submit"`
	TanggalPublish time.Time `json:"tanggal_publish"`
	Link           string    `json:"link"`
}

func (KaryaResponse) TableName() string {
	return "karyas"
}
