package dao

import (
	"errors"
	"myBlog/model"
)

var UserInterface SqlUser

type SqlUser struct {
}

// CheckUser 查询用户是否存在
func (u *SqlUser) CheckUser(name string) (int, error) {

	var user []model.User
	err := Db.Select("id").Where("userName=?", name).Find(&user).Error
	if err != nil {
		return model.ErrInner, err
	}

	if len(user) > 0 {
		return model.ErrUserNameUsed, errors.New(model.GetErrMsg(model.ErrUserNameUsed))
	}

	return model.Success, nil
}

// CreateUser 新增用户
func (u *SqlUser) CreateUser(data *model.User) (int, error) {

	err := Db.Create(&data).Error
	if err != nil {
		return model.Error, err
	}

	return model.Success, nil
}

// GetUsers 查询用户列表
func (u *SqlUser) GetUsers(userName string, pageSize int, pageNum int) (users []model.User, total int64, code int, err error) {

	// 默认是success,后面如果有错误，会直接被替换掉
	code = model.Success
	if userName != "" {

		err = Db.Select("id,userName,role,created_at").Where("userName like ?", userName+"%").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users).Error
		if err != nil {
			code = model.ErrInner
			return
		}

		err = Db.Model(&users).Where("userName like ?", userName+"%").Count(&total).Error
		if err != nil {
			code = model.ErrInner
			return
		}

		return
	}

	err = Db.Select("id,userName,role,created_at").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users).Error
	if err != nil {
		code = model.ErrInner
		return
	}

	err = Db.Model(&users).Count(&total).Error
	if err != nil {
		code = model.ErrInner
		return
	}

	return
}
