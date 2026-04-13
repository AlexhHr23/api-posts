package models

type Post struct {
	ID        uint `json:"id"`
	UserID    uint `json:"user_id"`
	Title     uint `json:"title"`
	Content   uint `json:"content"`
	CreatedAt uint `json:"created_at"`
	UpdatedAt uint `json:"updated_at"`
}
