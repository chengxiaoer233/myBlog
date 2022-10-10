package v1

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"myBlog/dao"
	"myBlog/model"
	"net/http"
	"strconv"
)

// AddArticle 添加文章
func AddArticle(c *gin.Context) {

	var data model.Article
	err := c.ShouldBindJSON(&data)
	if err != nil {
		c.JSON(http.StatusOK, resp2Client(model.ErrInner, err.Error(), nil))
		return
	}

	err = dao.ArticleInterface.CreateArt(&data)
	if err != nil {
		c.JSON(http.StatusOK, resp2Client(model.ErrInner, err.Error(), nil))
		return
	}

	// 阅读记录表也需要插入数据
	var read model.Readdetail
	read.ArticleId = int(data.ID)
	read.Count = 1
	err = dao.DataInterface.CreateRecord(&read)
	if err != nil {
		c.JSON(http.StatusOK, resp2Client(model.ErrInner, err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, resp2Client(model.Success, model.GetErrMsg(model.Success), data))
	return
}

// DeleteArt 删除文章
func DeleteArticle(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusOK, resp2Client(model.ErrInner, err.Error(), nil))
		return
	}

	// 需要查询文章是否存在
	code, err := dao.ArticleInterface.CheckArticleById(id)
	if code == model.ErrInner { //内部错误，直接返回
		c.JSON(http.StatusOK, resp2Client(model.ErrInner, err.Error(), nil))
		return
	} else if code == model.Success { //文章没找到,直接返回
		c.JSON(http.StatusOK, resp2Client(model.ErrArtNotExists, model.GetErrMsg(model.ErrArtNotExists), nil))
		return
	}

	// 删除已经存在的文章
	err = dao.ArticleInterface.DeleteArt(id)
	if err != nil {
		c.JSON(http.StatusOK, resp2Client(model.ErrInner, err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, resp2Client(model.Success, model.GetErrMsg(model.Success), nil))
	return
}

// EditArt 编辑文章
func EditArticle(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusOK, resp2Client(model.ErrInner, err.Error(), nil))
		return
	}

	var data model.Article
	err = c.ShouldBindJSON(&data)
	if err != nil {
		c.JSON(http.StatusOK, resp2Client(model.ErrInner, err.Error(), nil))
		return
	}

	// 需要查询文章是否存在
	code, err := dao.ArticleInterface.CheckArticleById(id)
	if code == model.ErrInner { //内部错误，直接返回
		c.JSON(http.StatusOK, resp2Client(model.ErrInner, err.Error(), nil))
		return
	} else if code == model.Success { //文章没找到,直接返回
		c.JSON(http.StatusOK, resp2Client(model.ErrArtNotExists, model.GetErrMsg(model.ErrArtNotExists), nil))
		return
	}

	// 更新文章
	err = dao.ArticleInterface.EditArticle(id, &data)
	if err != nil {
		c.JSON(http.StatusOK, resp2Client(model.ErrInner, err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, resp2Client(model.Success, model.GetErrMsg(model.Success), nil))
	return
}

// GetOneArtInfo 查询单个文章
func GetOneArtInfo(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusOK, resp2Client(model.ErrInner, err.Error(), nil))
		return
	}

	// 需要查询文章是否存在
	code, err := dao.ArticleInterface.CheckArticleById(id)
	if code == model.ErrInner { //内部错误，直接返回
		c.JSON(http.StatusOK, resp2Client(model.ErrInner, err.Error(), nil))
		return
	} else if code == model.Success { //文章没找到,直接返回
		c.JSON(http.StatusOK, resp2Client(model.ErrArtNotExists, model.GetErrMsg(model.ErrArtNotExists), nil))
		return
	}

	// 直接返回查询结果
	resp, err := dao.ArticleInterface.GetArtInfo(id)
	if err != nil && err != gorm.ErrRecordNotFound {
		c.JSON(http.StatusOK, resp2Client(model.ErrInner, err.Error(), nil))
		return
	} else if err == gorm.ErrRecordNotFound {
		c.JSON(http.StatusOK, resp2Client(model.ErrArtNotExists, model.GetErrMsg(model.ErrArtNotExists), nil))
		return
	}

	c.JSON(http.StatusOK, resp2Client(model.Success, model.GetErrMsg(model.Success), resp))
}

// GetArts 查询所有文章列表
func GetArts(c *gin.Context) {

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

	title := c.DefaultQuery("title", "") //文章搜索

	switch {
	case pageSize >= 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}

	if pageNum == 0 {
		pageNum = 1
	}

	// 查询所有的文章
	if title == "" {

		resp, total, err := dao.ArticleInterface.GetAllArts(pageSize, pageNum)
		if err != nil && err != gorm.ErrRecordNotFound {
			c.JSON(http.StatusOK, resp2Client(model.ErrInner, err.Error(), nil))
			return
		} else if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusOK, resp2Client(model.ErrCateNotExists, model.GetErrMsg(model.ErrCateNotExists), nil))
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status": model.Success,
			"data":   resp,
			"total":  total,
			"msg":    model.GetErrMsg(model.Success),
		})

		return
	}

	// 带参数搜索（这里需要支持模糊搜索）
	resp, total, err := dao.ArticleInterface.SearchArticle(title, pageSize, pageNum)
	if err != nil && err != gorm.ErrRecordNotFound {
		c.JSON(http.StatusOK, resp2Client(model.ErrInner, err.Error(), nil))
		return
	} else if err == gorm.ErrRecordNotFound {
		c.JSON(http.StatusOK, resp2Client(model.ErrCateNotExists, model.GetErrMsg(model.ErrCateNotExists), nil))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": model.Success,
		"data":   resp,
		"total":  total,
		"msg":    model.GetErrMsg(model.Success),
	})

	return
}

// GetCateArt 查询分类下的所有文章
func GetCateArt(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusOK, resp2Client(model.ErrInner, err.Error(), nil))
		return
	}

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

	resp, total, err := dao.ArticleInterface.GetCateArts(id, pageSize, pageNum)
	if err != nil && err != gorm.ErrRecordNotFound {
		c.JSON(http.StatusOK, resp2Client(model.ErrInner, err.Error(), nil))
		return
	} else if err == gorm.ErrRecordNotFound {
		c.JSON(http.StatusOK, resp2Client(model.ErrCateNotExists, model.GetErrMsg(model.ErrCateNotExists), nil))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": model.Success,
		"data":   resp,
		"total":  total,
		"msg":    model.GetErrMsg(model.Success),
	})

	return
}
