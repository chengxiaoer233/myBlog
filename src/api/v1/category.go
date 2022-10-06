package v1

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"myBlog/dao"
	"myBlog/model"
	"net/http"
	"strconv"
)

// AddCategory 添加分类
func AddCategory(c *gin.Context) {

	var data model.Category
	err := c.ShouldBindJSON(&data)
	if err != nil {
		c.JSON(http.StatusOK, resp2Client(model.Error, err.Error(), nil))
		return
	}

	// 检查分类是否已经存在
	if code, err := dao.CategoryInterface.CheckCategoryByName(data.Name); code == model.ErrCateNameUsed {
		c.JSON(http.StatusOK, resp2Client(code, err.Error(), nil))
		return
	}

	// 插入数据
	code, err := dao.CategoryInterface.CreateCategory(&data)
	if err != nil {
		c.JSON(http.StatusOK, resp2Client(code, err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, resp2Client(model.Success, model.GetErrMsg(model.Success), nil))
	return
}

// DeleteCate 删除分类
func DeleteCate(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusOK, resp2Client(model.ErrInner, err.Error(), nil))
		return
	}

	// 需要查询分类是否存在
	code, err := dao.CategoryInterface.CheckCategoryById(id)
	if code == model.ErrInner { //内部错误，直接返回
		c.JSON(http.StatusOK, resp2Client(model.ErrInner, err.Error(), nil))
		return
	} else if code == model.Success { //分类没找到,直接返回
		c.JSON(http.StatusOK, resp2Client(model.ErrCateNotExists, model.GetErrMsg(model.ErrCateNotExists), nil))
		return
	}

	// 删除已经存在的分类
	err = dao.CategoryInterface.DeleteCategory(id)
	if err != nil {
		c.JSON(http.StatusOK, resp2Client(model.ErrInner, err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, resp2Client(model.Success, model.GetErrMsg(model.Success), nil))
	return
}

// EditCate 编辑分类名
func EditCate(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusOK, resp2Client(model.ErrInner, err.Error(), nil))
		return
	}

	var data model.Category
	err = c.ShouldBindJSON(&data)
	if err != nil {
		c.JSON(http.StatusOK, resp2Client(model.ErrInner, err.Error(), nil))
		return
	}

	// 需要查询分类是否存在
	code, err := dao.CategoryInterface.CheckCategoryById(id)
	if code == model.ErrInner { //内部错误，直接返回
		c.JSON(http.StatusOK, resp2Client(model.ErrInner, err.Error(), nil))
		return
	} else if code == model.Success { //分类没找到,直接返回
		c.JSON(http.StatusOK, resp2Client(model.ErrCateNotExists, model.GetErrMsg(model.ErrCateNotExists), nil))
		return
	}

	// 更新分类
	code, err = dao.CategoryInterface.EditCategory(id, &data)
	if err != nil {
		c.JSON(http.StatusOK, resp2Client(code, model.GetErrMsg(code), nil))
		return
	}

	c.JSON(http.StatusOK, resp2Client(model.Success, model.GetErrMsg(model.Success), nil))
	return
}

// GetCateInfo 查询一个分类信息
func GetCateInfo(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusOK, resp2Client(model.ErrInner, err.Error(), nil))
		return
	}

	data, err := dao.CategoryInterface.GetOneCateGory(id)
	if err != nil && err != gorm.ErrRecordNotFound {
		c.JSON(http.StatusOK, resp2Client(model.ErrInner, err.Error(), nil))
		return
	} else if err == gorm.ErrRecordNotFound {
		c.JSON(http.StatusOK, resp2Client(model.ErrCateNotExists, model.GetErrMsg(model.ErrCateNotExists), nil))
		return
	}

	c.JSON(http.StatusOK, resp2Client(model.Success, model.GetErrMsg(model.Success), data))
}

// GetCate 查询所有分类列表
func GetCateS(c *gin.Context) {

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

	switch {
	case pageSize >= 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}

	if pageNum == 0 {
		pageNum = 1
	}

	data, total, err := dao.CategoryInterface.GetAllCateS(pageSize, pageNum)
	if err != nil && err != gorm.ErrRecordNotFound {
		c.JSON(http.StatusOK, resp2Client(model.ErrInner, err.Error(), nil))
		return
	} else if err == gorm.ErrRecordNotFound {
		c.JSON(http.StatusOK, resp2Client(model.ErrCateNotExists, model.GetErrMsg(model.ErrCateNotExists), nil))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": model.Success,
		"data":   data,
		"total":  total,
		"msg":    model.GetErrMsg(model.Success),
	})

	return
}
