package v1

import (
	"github.com/gin-gonic/gin"
	"myBlog/dao"
	"myBlog/model"
	"net/http"
	"fmt"
	"strconv"
)

// AddArticle 添加文章
func AddArticle(c *gin.Context) {

	var data model.Article
	err := c.ShouldBindJSON(&data)
	if err != nil {
		c.JSON(http.StatusOK,resp2Client(model.ErrInner,err.Error(),nil))
		return
	}

	err = dao.ArticleInterface.CreateArt(&data)
	if err != nil {
		c.JSON(http.StatusOK,resp2Client(model.ErrInner,err.Error(),nil))
		return
	}

	fmt.Printf("%+v=",data)
	c.JSON(http.StatusOK,resp2Client(model.Success,model.GetErrMsg(model.Success),data))
	return
}

// DeleteArt 删除文章
func DeleteArticle(c *gin.Context) {

	id,err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusOK, resp2Client(model.ErrInner, err.Error(), nil))
		return
	}

	// 需要查询文章是否存在
	code,err := dao.ArticleInterface.CheckArticleById(id)
	if code == model.ErrInner { //内部错误，直接返回
		c.JSON(http.StatusOK, resp2Client(model.ErrInner, err.Error(), nil))
		return
	}else if code == model.Success{ //文章没找到,直接返回
		c.JSON(http.StatusOK, resp2Client(model.ErrArtNotExists, model.GetErrMsg(model.ErrArtNotExists), nil))
		return
	}

	// 删除已经存在的文章
	err = dao.ArticleInterface.DeleteArt(id)
	if err != nil{
		c.JSON(http.StatusOK, resp2Client(model.ErrInner, err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, resp2Client(model.Success, model.GetErrMsg(model.Success), nil))
	return
}

// EditArt 编辑文章
func EditArticle(c *gin.Context) {

	id,err := strconv.Atoi(c.Param("id"))
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
	code,err := dao.ArticleInterface.CheckArticleById(id)
	if code == model.ErrInner { //内部错误，直接返回
		c.JSON(http.StatusOK, resp2Client(model.ErrInner, err.Error(), nil))
		return
	}else if code == model.Success{ //文章没找到,直接返回
		c.JSON(http.StatusOK, resp2Client(model.ErrArtNotExists, model.GetErrMsg(model.ErrArtNotExists), nil))
		return
	}

	// 更新文章
	err = dao.ArticleInterface.EditArticle(id,&data)
	if err != nil{
		c.JSON(http.StatusOK,resp2Client(model.ErrInner,err.Error(),nil))
		return
	}

	c.JSON(http.StatusOK, resp2Client(model.Success, model.GetErrMsg(model.Success), nil))
	return
}

// GetCateArt 查询分类下的所有文章
func GetCateArt(c *gin.Context) {

}

// GetArtInfo 查询单个文章信息
func GetArtInfo(c *gin.Context) {

}

// GetArt 查询文章列表
func GetArt(c *gin.Context) {

}




