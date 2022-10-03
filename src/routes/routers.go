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
		routerV1User := routerV1.Group("user")
		{
			// 增加用户、查询单个用户、查询所有用户
			routerV1User.POST("add", v1.AddUser)
			routerV1User.GET("/:id", v1.GetOneUserInfo)
			routerV1User.GET("/users", v1.GetUsers)
		}

		// 分类模块
		routerV1Cate := routerV1.Group("category")
		{
			// 查询单个分类、查询所有分类
			routerV1Cate.GET("/", v1.GetCateS)
			routerV1Cate.GET("/:id", v1.GetCateInfo)
		}
/*
		//  文章模块
		routerV1 = routerV1.Group("article")
		{
			routerV1.GET("/", v1.GetArt)
			routerV1.GET("/list/:id", v1.GetCateArt)
			routerV1.GET("/info/:id", v1.GetArtInfo)
		}*/
	}

	// admin后台相关接口
	routerAdmin := r.Group("api/v1/admin")
	{
		// 用户模块
		routerAdminUser := routerAdmin.Group("user")
		{
			// 删除某个用户、修改某个用户信息、查询所有用户信息
			routerAdminUser.DELETE("/", v1.DeleteUser)
			routerAdminUser.PUT("/", v1.EditUser)
			routerAdminUser.GET("/users", v1.GetUsers)
		}

		// 分类模块
		routerAdminCate := routerAdmin.Group("category")
		{
			// 增加分类、删除分类、修改分类、查询分类
			routerAdminCate.POST("/add", v1.AddCategory)
			routerAdminCate.DELETE("/:id", v1.DeleteCate)
			routerAdminCate.PUT("/:id", v1.EditCate)
			routerAdminCate.GET("/", v1.GetCateS)
		}
	}

	// run
	r.Run(model.ServerConf.HttpPort)
}
