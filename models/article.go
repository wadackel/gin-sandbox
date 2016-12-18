package models

type Article struct {
	Model
	UserID uint   `gorm:"index"`
	Title  string `json:"title"`
	Body   string `json:"body" gorm:"type:text"`
}
