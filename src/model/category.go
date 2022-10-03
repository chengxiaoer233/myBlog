package model

type Category struct {
	ID   uint   `json:"id"   gorm:"primary_key;auto_increment" `
	Name string `json:"name" gorm:"type:varchar(20);not null"`
}
