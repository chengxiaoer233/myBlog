package model

import (
	"gorm.io/gorm"
)

type Readdetail struct {
	gorm.Model
	ArticleId int    `json:"article_id"   gorm:"type:int;default:0"`
	UserName  string `json:"username"     gorm:"column:username;type:varchar(20)"`
	Count     int    `json:"count"        gorm:"type:int;not null;default:0"`
}
