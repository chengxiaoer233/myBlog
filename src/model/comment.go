package model

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	UserId    uint   `json:"user_id"`
	ArticleId uint   `json:"article_id"`
	Title     string `json:"article_title"`
	Username  string `json:"username"`
	Content   string `json:"content" gorm:"type:varchar(500);not null;"`
	Status    int8   `json:"status"  gorm:"type:tinyint;default:2" `
}
