package model

import "gorm.io/gorm"

type User struct {
	// （1）validate规则校验，防止用户输入无效数据,其实一般在用户接入层做校验较多，从源头上防止用户输入无效数据
	gorm.Model
	UserName string `json:"userName" gorm:"type:varchar(20);not null"  validate:"required,min=4,max=12"  label:"用户名"`
	Password string `json:"password" gorm:"type:varchar(500);nou null" validate:"required,min=6,max=120" label:"密码"`
	Role     int    `json:"role"     gorm:"type int;default 2"         validate:"required,gte=2"         label:"角色码"`
}
