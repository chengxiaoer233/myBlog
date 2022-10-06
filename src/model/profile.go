package model

type Profile struct {
	ID        int    `gorm:"primaryKey" json:"id"`
	Name      string `json:"name"      gorm:"column:name;type:varchar(20)"`
	Desc      string `json:"desc"      gorm:"column:desc;type:varchar(200)"`
	QqChat    string `json:"qq_chat"   gorm:"column:qqchat;type:varchar(200)"`
	Wechat    string `json:"wechat"    gorm:"column:wechat;type:varchar(100)"`
	Weibo     string `json:"weibo"     gorm:"column:weibo;type:varchar(200)"`
	Bili      string `json:"bili"      gorm:"column:bili;type:varchar(200)"`
	Email     string `json:"email"     gorm:"column:email;type:varchar(200)"`
	Img       string `json:"img"       gorm:"column:img;type:varchar(200)" `
	Avatar    string `json:"avatar"    gorm:"column:avatar;type:varchar(200)" `
	IcpRecord string `json:"icp_record" gorm:"column:icp_record;type:varchar(200)" `
}
