package models

type Tag struct {
	Model
	UserID uint   `json:"user_id" gorm:"index"`
	Name   string `json:"name"`
}
