package v1

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"myBlog/dao"
	"myBlog/model"
	"myBlog/utlis"
	"myBlog/utlis/validator"
	"net/http"
	"strconv"
)

// 通用返回接口
type resp struct {
	Status int         `json:"status"`
	Msg    string      `json:"msg"`
	Data   interface{} `json:"data"`
}

func resp2Client(code int, msg string, data interface{}) resp {

	return resp{
		Status: code,
		Msg:    msg,
		Data:   data,
	}
}

// AddUser 添加用户
func AddUser(c *gin.Context) {

	var data model.User
	err := c.ShouldBindJSON(&data)
	if err != nil {
		c.JSON(http.StatusOK, resp2Client(model.Error, err.Error(), nil))
		return
	}

	// 新建用户时，validator 参数校验，不用每个字符这样去校验
	// 这里也可以自己解析json参数校验，实际项目中自己校验的较多，为了代码的可读性
	msg, err := validator.Validate(&data)
	if err != nil {
		c.JSON(
			http.StatusOK, gin.H{
				"status":  model.ErrInner,
				"message": msg,
			},
		)
		c.Abort()
		return
	}

	// 检查用户是否已经存在
	if code, err := dao.UserInterface.CheckUser(data.UserName); code == model.ErrUserNameUsed {
		c.JSON(http.StatusOK, resp2Client(code, err.Error(), nil))
		return
	}

	// 数据库密码加密
	data.Password, err = utlis.GenerateFromPassword(data.Password)
	if err != nil {
		c.JSON(http.StatusOK, resp2Client(model.ErrInner, err.Error(), nil))
		return
	}

	// 插入数据
	code, err := dao.UserInterface.CreateUser(&data)
	if err != nil {
		c.JSON(http.StatusOK, resp2Client(code, err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, resp2Client(model.Success, model.GetErrMsg(model.Success), nil))
	return
}

// DeleteUser 删除用户
// 这里采用delete + json方式，需要用户携带密码才可以删除
func DeleteUser(c *gin.Context) {

	var id int
	var err error
	if id, err = strconv.Atoi(c.Param("id")); err != nil {
		c.JSON(http.StatusOK, resp2Client(model.ErrInner, err.Error(), nil))
		return
	}

	// 先跑通，这里使用的是带query参数而不会post + json + body
	/*// 定义内部变量
	type Req struct {
		Id       int    `json:"id"`
		PassWord string `json:"passWord"`
	}

	// json解析
	var req Req
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusOK, resp2Client(model.ErrInner, err.Error(), nil))
		return
	}*/

	// 查询单个用户
	_, err = dao.UserInterface.GetUserById(id)
	if err != nil && err != gorm.ErrRecordNotFound {
		c.JSON(http.StatusOK, resp2Client(model.ErrInner, err.Error(), nil))
		return

	} else if err == gorm.ErrRecordNotFound {
		c.JSON(http.StatusOK, resp2Client(model.ErrUserNotExists, model.GetErrMsg(model.ErrUserNotExists), nil))
		return
	}

	// 严格来讲，这里是需要校验用户秘钥的，暂时放过
	/*// 用户存在，需要判断秘钥是否正确
	code, err := utlis.CompareHashAndPassword([]byte(data.Password), []byte(req.PassWord))
	if code == model.ErrInner {
		c.JSON(http.StatusOK, resp2Client(code, err.Error(), nil))
		return
	}
	if err != nil {
		c.JSON(http.StatusOK, resp2Client(model.ErrTokenWrong, model.GetErrMsg(model.ErrTokenWrong), nil))
		return
	}*/

	// 直接删除用户
	code, err := dao.UserInterface.DeleteUser(id)
	c.JSON(http.StatusOK, resp2Client(code, err.Error(), nil))
	return
}

// EditUser 编辑用户,输入为json格式,需要做严格校验 用户id + 用户名
func EditUser(c *gin.Context) {

	var id int
	var err error
	if id, err = strconv.Atoi(c.Param("id")); err != nil {
		c.JSON(http.StatusOK, resp2Client(model.ErrInner, err.Error(), nil))
		return
	}

	var user model.User
	err = c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusOK, resp2Client(model.ErrInner, err.Error(), nil))
		return
	}

	// 根据userId查找用户相关信息
	_, err = dao.UserInterface.GetUserById(id)
	if err != nil && err != gorm.ErrRecordNotFound { // 其他错误
		c.JSON(http.StatusOK, resp2Client(model.ErrInner, err.Error(), nil))
		return
	} else if err == gorm.ErrRecordNotFound { //没找到，直接insert
		code, err := dao.UserInterface.EditUser(&user)
		if err != nil {
			c.JSON(http.StatusOK, resp2Client(model.ErrInner, err.Error(), nil))
			return
		}

		c.JSON(http.StatusOK, resp2Client(code, model.GetErrMsg(code), nil))
		return
	}

	/*	// 已经找到对应的用户信息
		if len(users) > 1 { // 找到多个用户，直接返回
			c.JSON(http.StatusOK, resp2Client(model.ErrUserNameUsed, model.GetErrMsg(model.ErrUserNameUsed), nil))
			return
		} else if len(users) == 1 { // 找到一个用户信息
			id := users[0].ID

			if id == user.ID { // 是当前用户，直接更新
				code, err := dao.UserInterface.EditUser(&user)
				if err != nil {
					c.JSON(http.StatusOK, resp2Client(model.ErrInner, err.Error(), nil))
					return
				}

				c.JSON(http.StatusOK, resp2Client(code, model.GetErrMsg(code), nil))
				return

			} else { // 其他用户，报错
				c.JSON(http.StatusOK, resp2Client(model.ErrUserNameUsed, model.GetErrMsg(model.ErrUserNameUsed), nil))
				return
			}
		}*/

	// 严格来讲这里需要校验用户密码之类的，暂时放过，直接更新用户信息
	user.ID = uint(id)
	code, err := dao.UserInterface.EditUser(&user)
	if err != nil {
		c.JSON(http.StatusOK, resp2Client(model.ErrInner, err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, resp2Client(code, model.GetErrMsg(code), nil))
	return
}

// GetUserInfo 查询单个用户
func GetOneUserInfo(c *gin.Context) {

	var id int
	var err error
	if id, err = strconv.Atoi(c.Param("id")); err != nil {
		c.JSON(http.StatusOK, resp2Client(model.ErrInner, err.Error(), nil))
		return
	}

	data, err := dao.UserInterface.GetUserById(id)
	if err != nil && err != gorm.ErrRecordNotFound {
		c.JSON(http.StatusOK, resp2Client(model.ErrInner, err.Error(), nil))
		return

	} else if err == gorm.ErrRecordNotFound {
		c.JSON(http.StatusOK, resp2Client(model.ErrUserNotExists, model.GetErrMsg(model.ErrUserNotExists), nil))
		return
	}

	// 信息返回，不用返回用户密码
	data.Password = "********"

	c.JSON(http.StatusOK, resp2Client(model.Success, model.GetErrMsg(model.Success), data))
	return
}

// admin后台相关接口
// GetUsers 查询所有用户列表
func GetUsers(c *gin.Context) {

	pageSize, err := strconv.Atoi(c.DefaultQuery("pagesize", "10"))
	if err != nil {
		c.JSON(http.StatusOK, resp2Client(model.ErrInner, err.Error(), nil))
		return
	}

	pageNum, err := strconv.Atoi(c.DefaultQuery("pagenum", "1"))
	if err != nil {
		c.JSON(http.StatusOK, resp2Client(model.ErrInner, err.Error(), nil))
		return
	}

	userName := c.Query("userName")

	switch {
	case pageSize >= 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}

	if pageNum == 0 {
		pageNum = 1
	}

	data, total, code, err := dao.UserInterface.GetUsers(userName, pageSize, pageNum)
	if code != model.Success {
		c.JSON(http.StatusOK, resp2Client(model.ErrInner, err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"data":   data,
		"total":  total,
		"msg":    model.GetErrMsg(code),
	})
}

// ChangeUserPassword 修改密码
func ChangeUserPassword(c *gin.Context) {

	var id int
	var err error
	if id, err = strconv.Atoi(c.Param("id")); err != nil {
		c.JSON(http.StatusOK, resp2Client(model.ErrInner, err.Error(), nil))
		return
	}

	var user model.User
	err = c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusOK, resp2Client(model.ErrInner, err.Error(), nil))
		return
	}

	_, err = dao.UserInterface.GetUserById(id)
	if err != nil && err != gorm.ErrRecordNotFound {
		c.JSON(http.StatusOK, resp2Client(model.ErrInner, err.Error(), nil))
		return

	} else if err == gorm.ErrRecordNotFound {
		c.JSON(http.StatusOK, resp2Client(model.ErrUserNotExists, model.GetErrMsg(model.ErrUserNotExists), nil))
		return
	}

	// 严格来讲，这里需要验证用户的密码，暂时就不校验了，直接更新

	// 数据库密码加密
	user.Password, err = utlis.GenerateFromPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusOK, resp2Client(model.ErrInner, err.Error(), nil))
		return
	}

	// 严格来讲这里需要校验用户密码之类的，暂时放过，直接更新用户信息
	user.ID = uint(id)
	code, err := dao.UserInterface.EditPassWord(&user)
	if err != nil {
		c.JSON(http.StatusOK, resp2Client(model.ErrInner, err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, resp2Client(code, model.GetErrMsg(code), nil))
	return
}
