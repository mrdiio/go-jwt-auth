package models

type Article struct {
	ID      uint   `json:"id" gorm:"primary_key"`
	Title   string `json:"title"`
	Content string `json:"content"`
	UserID  uint   `json:"user_id"`
	User    User   `json:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
