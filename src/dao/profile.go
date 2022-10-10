package dao

import "myBlog/model"

var ProfileInterface SqlProfile

type SqlProfile struct {
}

// CheckUser 查询主页id是否存在
func (p *SqlProfile) GetProfileById(id int) (profile model.Profile, err error) {
	err = Db.Where("id = ?", id).Find(&profile).Error
	return
}

// UpdateUser 更新主页信息
func (p *SqlProfile) UpdateProfileById(profile model.Profile) (err error) {
	err = Db.Model(&model.Profile{}).Where("id = ?", profile.ID).Updates(&profile).Error
	return
}
