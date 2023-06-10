package request

import "logbook_ta/models/response"

type KaryaRequest struct {
	UserID         uint `json:"user_id" form:"user_id"`
	User           response.UserResponse
	TanggalSubmit  string `json:"tanggal_submit" form:"tanggal_submit"`
	TanggalPublish string `json:"tanggal_publish" form:"tanggal_publish"`
	Link           string `json:"link" form:"link" `
}

func (KaryaRequest) TableName() string {
	return "karyas"
}
