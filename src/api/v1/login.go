package v1

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"myBlog/dao"
	"myBlog/middleware"
	"myBlog/model"
	"myBlog/utlis"
	"net/http"
)

const (
	admin = "admin"
	front = "front"
)

// 后台登录
func Login(c *gin.Context) {
	login(c, "admin")
}

func LoginFront(c *gin.Context) {
	login(c, "front")
}

func login(c *gin.Context, src string) {

	var user model.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusOK, resp2Client(model.ErrInner, err.Error(), nil))
		return
	}

	// 根据用户名查询单个用户
	resp, err := dao.UserInterface.GetUsersByName(user.UserName)
	if err != nil && err != gorm.ErrRecordNotFound {
		c.JSON(http.StatusOK, resp2Client(model.ErrInner, err.Error(), nil))
		return

	} else if err == gorm.ErrRecordNotFound {
		c.JSON(http.StatusOK, resp2Client(model.ErrUserNotExists, model.GetErrMsg(model.ErrUserNotExists), nil))
		return
	}

	// 用户存在，需要判断秘钥是否正确
	code, err := utlis.CompareHashAndPassword([]byte(resp[0].Password), []byte(user.Password))
	if code == model.ErrInner {
		c.JSON(http.StatusOK, resp2Client(code, err.Error(), nil))
		return
	}
	if err != nil {
		c.JSON(http.StatusOK, resp2Client(model.ErrTokenWrong, model.GetErrMsg(model.ErrTokenWrong), nil))
		return
	}

	// 根据是前端还是后台接口，返回对应的结果
	if src == front {
		c.JSON(http.StatusOK, gin.H{
			"status":  model.Success,
			"data":    user.UserName,
			"id":      resp[0].ID,
			"message": model.GetErrMsg(model.Success),
		})
		return
	}

	// 判断用户权限是否满足，不满足的直接返回
	if resp[0].Role != 1 { // 1为管理员，2位普通用户
		c.JSON(http.StatusOK, resp2Client(model.ErrUserNoRight, model.GetErrMsg(model.ErrUserNoRight), nil))
		return
	}

	// 鉴权通过，签发token
	token, err := middleware.NewJWT().SetToken(c, &user)
	if err != nil {
		c.JSON(http.StatusOK, resp2Client(model.ErrInner, err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  model.Success,
		"data":    user.UserName,
		"id":      user.ID,
		"message": model.GetErrMsg(model.Success),
		"token":   token,
	})
}
