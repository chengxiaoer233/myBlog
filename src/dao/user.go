package dao

import (
	"errors"
	"gorm.io/gorm"
	"myBlog/model"
	"fmt"
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

// 删除用户之前的用户信息校验,判断是否存在及合法
func (u *SqlUser)CheckUserBeforeDelete(id int,password string)(int,error){

	var user model.User
	err := Db.Where("id = ? and password = ?",id,password).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound{
		return model.ErrInner,err
	}

	// 存在且信息都正确
	if user.ID > 0 {
		return model.Success,nil
	}

	// 不存在该用户信息
	return model.ErrUserNameOrPwWrong ,errors.New(model.GetErrMsg(model.ErrUserNameOrPwWrong))
}

// CreateUser 新增用户
func (u *SqlUser) CreateUser(data *model.User) (int, error) {

	err := Db.Create(&data).Error
	if err != nil {
		return model.Error, err
	}

	return model.Success, nil
}

// GetUserInfo 根据id查询单个用户信息
func (u *SqlUser) GetUser(id int)(user *model.User,err error){
	err = Db.Where("id = ?",id).First(&user).Error
	fmt.Println("user=",user,",err=",err)
	return
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

// DeleteUser 删除用户列表，改用delete + json body的方式，需要前端携带密码
func (u *SqlUser) DeleteUser(id int)(int,error){

	// 这里采用的是硬删除的方式，gorm默认是软删除
	err := Db.Where("id = ?",id).Unscoped().Delete(&model.User{}).Error
	if err != nil{
		return model.ErrInner,err
	}

	return model.Success,errors.New("ok")
}