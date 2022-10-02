package routes

import (
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	"myBlog/model"
	"net/http"
)

// 引入前端渲染
func createMyRender() multitemplate.Renderer {
	p := multitemplate.NewRenderer()
	p.AddFromFiles("admin", "web/admin/dist/index.html")
	p.AddFromFiles("front", "web/front/dist/index.html")
	return p
}

func InitRouters() {

	// setMode
	gin.SetMode(model.ServerConf.AppMode)

	// 采用默认的gin 中间件
	r := gin.Default()

	// gin handle
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	})

	// run
	r.Run(model.ServerConf.HttpPort)
}
