package response

type UserResponse struct {
	ID       uint   `json:"id"`
	Name     string `json:"name" `
	Username string `json:"username" `
	Email    string `json:"email"`
	// CreatedAt time.Time `json:"created_at"`
	// UpdatedAt time.Time `json:"updated_at"`
	// DeletedAt gorm.DeletedAt `json:"-"`
}
