package response

type LogbookRequest struct {
	UserID     uint   `json:"user_id"`
	Judul      string `json:"judul"`
	Topik      string `json:"topik"`
	Pembimbing string `json:"pembimbing"`
}

func (LogbookRequest) TableName() string {
	return "logbooks"
}
