package routes

import (
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	v1 "myBlog/api/v1"
	"myBlog/middleware"
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

	r := gin.New()
	// 设置信任网络 []string
	// nil 为不计算，避免性能消耗，上线应当设置
	_ = r.SetTrustedProxies(nil)

	r.HTMLRender = createMyRender()
	r.Use(middleware.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.Cors())

	r.Static("/static", "./web/front/dist/static")
	r.Static("/admin", "./web/admin/dist")
	r.StaticFile("/favicon.ico", "/web/front/dist/favicon.ico")

	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "front", nil)
	})

	r.GET("/admin", func(c *gin.Context) {
		c.HTML(200, "admin", nil)
	})

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

		// 文章模块
		routerAdminArticle := routerAdmin.Group("article")
		{
			// 添加文章、删除文章、修改文章、查询文章
			routerAdminArticle.POST("/add", v1.AddArticle)
			routerAdminArticle.DELETE("/:id", v1.DeleteArticle)
			routerAdminArticle.PUT("/:id", v1.EditArticle)
		//	routerAdminArticle.GET("/", v1.GetCateS)
		}
	}

	// run
	r.Run(model.ServerConf.HttpPort)
}
