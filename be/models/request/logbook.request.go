package request

type LogbookRequest struct {
	UserID     uint   `json:"user_id" form:"user_id" validate:"required"`
	Judul      string `json:"judul"  form:"judul" validate:"required"`
	Topik      string `json:"topik" form:"topik" validate:"required"`
	Pembimbing string `json:"pembimbing" form:"pembimbing" validate:"required"`
}

type LogbookUpdateRequest struct {
	//UserID     uint   `json:"user_id" form:"user_id"`
	Judul      string `json:"judul"  form:"judul" validate:"required"`
	Topik      string `json:"topik" form:"topik" validate:"required"`
	Pembimbing string `json:"pembimbing" form:"pembimbing" validate:"required"`
}

func (LogbookRequest) TableName() string {
	return "logbooks"
}
