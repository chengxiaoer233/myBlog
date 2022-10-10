package v1

import (
	"github.com/gin-gonic/gin"
	"myBlog/dao"
	"myBlog/model"
	"net/http"
	"strconv"
)

// AddComment 添加评论
func AddComment(c *gin.Context) {

	var data model.Comment
	err := c.ShouldBindJSON(&data)
	if err != nil {
		c.JSON(http.StatusOK, resp2Client(model.Error, err.Error(), nil))
		return
	}

	err = dao.CommentInterface.AddComment(data)
	if err != nil {
		c.JSON(http.StatusOK, resp2Client(model.ErrInner, err.Error(), nil))
		return
	}
	c.JSON(http.StatusOK, resp2Client(model.Success, model.GetErrMsg(model.Success), nil))
	return
}

// GetComment 获取谋篇文章的评论
func GetComment(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusOK, resp2Client(model.ErrInner, err.Error(), nil))
		return
	}

	data, err := dao.CommentInterface.GetOneComment(id)
	if err != nil {
		c.JSON(http.StatusOK, resp2Client(model.ErrInner, err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  model.Success,
		"data":    data,
		"message": model.GetErrMsg(model.Success),
	})
	return
}

// GetCommentListFront 前端展示页面显示评论列表
func GetCommentListFront(c *gin.Context) {

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

	id, err := strconv.Atoi(c.Param("id"))
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

	data, total, err := dao.CommentInterface.GetCommentListFront(id, pageSize, pageNum)
	if err != nil {
		c.JSON(http.StatusOK, resp2Client(model.ErrInner, err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  model.Success,
		"data":    data,
		"total":   total,
		"message": model.GetErrMsg(model.Success),
	})
	return
}

// GetCommentCount 获取谋篇文章的数量
func GetCommentCount(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusOK, resp2Client(model.ErrInner, err.Error(), nil))
		return
	}

	total, err := dao.CommentInterface.GetCommentCount(id)
	if err != nil {
		c.JSON(http.StatusOK, resp2Client(model.ErrInner, err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  model.Success,
		"total":   total,
		"message": model.GetErrMsg(model.Success),
	})
	return
}

// GetCommentList 后台接口查询所有评论列表
func GetCommentList(c *gin.Context) {

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

	data, total, err := dao.CommentInterface.GetCommentLists(pageSize, pageNum)
	if err != nil {
		c.JSON(http.StatusOK, resp2Client(model.ErrInner, err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  model.Success,
		"data":    data,
		"total":   total,
		"message": model.GetErrMsg(model.Success),
	})
	return
}

// DeleteComment 删除评论
func DeleteComment(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusOK, resp2Client(model.ErrInner, err.Error(), nil))
		return
	}

	err = dao.CommentInterface.DeleteComment(id)
	if err != nil {
		c.JSON(http.StatusOK, resp2Client(model.ErrInner, err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  model.Success,
		"message": model.GetErrMsg(model.Success),
	})
	return
}

// CheckComment 通过评论
func CheckComment(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusOK, resp2Client(model.ErrInner, err.Error(), nil))
		return
	}

	var data model.Comment
	err = c.ShouldBindJSON(&data)
	if err != nil {
		c.JSON(http.StatusOK, resp2Client(model.ErrInner, err.Error(), nil))
		return
	}

	err = dao.CommentInterface.CheckComment(id, &data)
	if err != nil {
		c.JSON(http.StatusOK, resp2Client(model.ErrInner, err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  model.Success,
		"message": model.GetErrMsg(model.Success),
	})
	return
}

// CheckComment 通过评论
func UnCheckComment(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusOK, resp2Client(model.ErrInner, err.Error(), nil))
		return
	}

	var data model.Comment
	err = c.ShouldBindJSON(&data)
	if err != nil {
		c.JSON(http.StatusOK, resp2Client(model.ErrInner, err.Error(), nil))
		return
	}

	err = dao.CommentInterface.UncheckComment(id, &data)
	if err != nil {
		c.JSON(http.StatusOK, resp2Client(model.ErrInner, err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  model.Success,
		"message": model.GetErrMsg(model.Success),
	})
	return
}
