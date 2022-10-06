package v1

import (
	"github.com/gin-gonic/gin"
	"myBlog/model"
	"net/http"
	"fmt"
	"strconv"
)

// 文章刷新页面302重定向
func ArticleDetail302(c *gin.Context){

	var id int
	var err error
	if id, err = strconv.Atoi(c.Param("id")); err != nil {
		c.JSON(http.StatusOK, resp2Client(model.ErrInner, err.Error(), nil))
		return
	}
	fmt.Println("host=",c.Request.Host,",url=",c.Request.URL)
	c.Redirect(302, c.Request.Host + fmt.Sprintf("/v1/article/info/%d",id))
}

// 获取分类页面302重定向
func GetCateArt302(c *gin.Context){

	var id int
	var err error
	if id, err = strconv.Atoi(c.Param("id")); err != nil {
		c.JSON(http.StatusOK, resp2Client(model.ErrInner, err.Error(), nil))
		return
	}

	fmt.Println("host=",c.Request.Host,",url=",c.Request.URL)
	c.Redirect(302, c.Request.Host + fmt.Sprintf("/api/v1/article/list/%d",id))
}