package models

type Article struct {
	Model
	UserID uint   `json:"user_id" gorm:"index"`
	Title  string `json:"title"`
	Body   string `json:"body" gorm:"type:text"`
}
