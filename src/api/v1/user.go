package v1

import (
	"github.com/gin-gonic/gin"
	"myBlog/dao"
	"myBlog/model"
	"myBlog/utlis"
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
		c.JSON(http.StatusOK, resp2Client(model.Error, "ShouldBindJSON error", nil))
		return
	}

	// 检查用户是否已经存在
	if code, err := dao.UserInterface.CheckUser(data.UserName); code == model.ErrUserNameUsed {
		c.JSON(http.StatusOK, resp2Client(code, err.Error(), nil))
		return
	}

	// 数据库密码加密
	data.Password,err = utlis.GenerateFromPassword(data.Password)
	if err != nil{
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

// GetUserInfo 查询单个用户
func GetUserInfo(c *gin.Context) {
}

// GetUsers 查询所有用户列表
func GetUsers(c *gin.Context) {

	pageSize, err := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	if err != nil {
		c.JSON(http.StatusOK, resp2Client(model.ErrInner, "pageSize strconv.Atoi error", nil))
		return
	}

	pageNum, err := strconv.Atoi(c.DefaultQuery("pageNum", "1"))
	if err != nil {
		c.JSON(http.StatusOK, resp2Client(model.ErrInner, "pageNum strconv.Atoi error", nil))
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

// EditUser 编辑用户
func EditUser(c *gin.Context) {
}

// ChangeUserPassword 修改密码
func ChangeUserPassword(c *gin.Context) {
}

// DeleteUser 删除用户
func DeleteUser(c *gin.Context) {
}
