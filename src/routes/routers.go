package routes

import (
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	v1 "myBlog/api/v1"
	"myBlog/model"
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

	// 采用路由组的形式来阻止路由关系
	// 前端展示页面
	routerV1 := r.Group("api/v1")
	{
		// 用户模块
		routerV1 = routerV1.Group("user")
		{
			routerV1.POST("add", v1.AddUser)
			routerV1.DELETE("/",v1.DeleteUser)
			routerV1.GET("/:id", v1.GetUserInfo)
			routerV1.GET("/users", v1.GetUsers)
		}

		// 分类模块
		routerV1 = routerV1.Group("category")
		{
			routerV1.POST("/", v1.GetCate)
			routerV1.GET("/:id", v1.GetCateInfo)
		}

		//  文章模块
		routerV1 = routerV1.Group("article")
		{
			routerV1.GET("/", v1.GetArt)
			routerV1.GET("/list/:id", v1.GetCateArt)
			routerV1.GET("/info/:id", v1.GetArtInfo)
		}
	}

	// run
	r.Run(model.ServerConf.HttpPort)
}
