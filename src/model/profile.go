package model

import "gorm.io/gorm"

type Profile struct {
	gorm.Model
	Name      string `json:"name"      gorm:"type:varchar(20)"`
	Desc      string `json:"desc"      gorm:"type:varchar(200)"`
	QqChat    string `json:"qqChat"    gorm:"type:varchar(200)"`
	Wechat    string `json:"weChat"    gorm:"type:varchar(100)"`
	Weibo     string `json:"weiBo"     gorm:"type:varchar(200)"`
	Bili      string `json:"bili"      gorm:"type:varchar(200)"`
	Email     string `json:"email"     gorm:"type:varchar(200)"`
	Img       string `json:"img"       gorm:"type:varchar(200)" `
	Avatar    string `json:"avatar"    gorm:"type:varchar(200)" `
	IcpRecord string `json:"icpRecord" gorm:"type:varchar(200)" `
}
