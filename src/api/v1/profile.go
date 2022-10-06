package v1

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"myBlog/dao"
	"myBlog/model"
	"net/http"
	"strconv"
)

func GetProfile(c*gin.Context){

	var id int
	var err error
	if id, err = strconv.Atoi(c.Param("id")); err != nil {
		c.JSON(http.StatusOK, resp2Client(model.ErrInner, err.Error(), nil))
		return
	}

	data, err := dao.ProfileInterface.GetProfileById(id)
	if err != nil && err != gorm.ErrRecordNotFound {
		c.JSON(http.StatusOK, resp2Client(model.ErrInner, err.Error(), nil))
		return

	} else if err == gorm.ErrRecordNotFound {
		c.JSON(http.StatusOK, resp2Client(model.ErrProfileNotExists, model.GetErrMsg(model.ErrProfileNotExists), nil))
		return
	}

	c.JSON(http.StatusOK, resp2Client(model.Success, model.GetErrMsg(model.Success), data))
	return
}

func UpdateProfile(c*gin.Context){

	var profile model.Profile
	err := c.ShouldBindJSON(&profile)
	if err != nil {
		c.JSON(http.StatusOK,resp2Client(model.ErrInner,err.Error(),nil))
		return
	}

	_, err = dao.ProfileInterface.GetProfileById(profile.ID)
	if err != nil && err != gorm.ErrRecordNotFound {
		c.JSON(http.StatusOK, resp2Client(model.ErrInner, err.Error(), nil))
		return

	} else if err == gorm.ErrRecordNotFound {
		c.JSON(http.StatusOK, resp2Client(model.ErrProfileNotExists, model.GetErrMsg(model.ErrProfileNotExists), nil))
		return
	}

	err = dao.ProfileInterface.UpdateProfileById(profile)
	if err != nil {
		c.JSON(http.StatusOK, resp2Client(model.ErrInner, err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, resp2Client(model.Success, model.GetErrMsg(model.Success), nil))
	return
}
