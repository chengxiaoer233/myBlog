package model

import "gorm.io/gorm"

type Article struct {
	Category Category `gorm:"foreignkey:Cid"`
	gorm.Model
	Title        string `json:"title"   gorm:"type:varchar(100);not null"`
	Cid          int    `json:"cid"     gorm:"type:int;not null"`
	Desc         string `json:"desc"    gorm:"type:varchar(200)"`
	Content      string `json:"content" gorm:"type:longtext"`
	Img          string `json:"img"     gorm:"type:varchar(100)"`
	CommentCount int    `json:"comment_count" gorm:"type:int;not null;default:0"`
	ReadCount    int    `json:"read_count"    gorm:"type:int;not null;default:0"`
}
