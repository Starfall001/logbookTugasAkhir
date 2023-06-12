package request

type KaryaRequest struct {
	UserID         uint   `json:"user_id" form:"user_id"`
	TanggalSubmit  string `json:"tanggal_submit" form:"tanggal_submit"`
	TanggalPublish string `json:"tanggal_publish" form:"tanggal_publish"`
	Link           string `json:"link" form:"link" `
}

type KaryaUpdateRequest struct {
	TanggalSubmit  string `json:"tanggal_submit" form:"tanggal_submit"`
	TanggalPublish string `json:"tanggal_publish" form:"tanggal_publish"`
	Link           string `json:"link" form:"link" `
}

func (KaryaRequest) TableName() string {
	return "karyas"
}
