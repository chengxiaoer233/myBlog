package routes

import (
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	v1 "myBlog/api/v1"
	"myBlog/middleware"
	"myBlog/middleware/logger"
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
	r.Use(logger.Logger())
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

	// 前端展示页面,采用路由组的形式来阻止路由关系
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
			routerV1Cate.GET("/:id", v1.GetCateInfo)
			routerV1Cate.GET("", v1.GetCateS)
		}

		//  文章模块
		routerV1Article := routerV1.Group("article")
		{
			// 查询单个、查询所有、查询某个分类下的所有文章
			routerV1Article.GET("/info/:id", v1.GetOneArtInfo)
			routerV1Article.GET("", v1.GetArts)
			routerV1Article.GET("/list/:id", v1.GetCateArt)

		}

		//  评论模块
		{
			routerV1.POST("addcomment", v1.AddComment)
			routerV1.GET("comment/info/:id", v1.GetComment)
			routerV1.GET("commentfront/:id", v1.GetCommentListFront)
			routerV1.GET("commentcount/:id", v1.GetCommentCount)
		}

		// 登录控制模块
		routerV1.POST("login", v1.Login)
		routerV1.POST("loginfront", v1.LoginFront)

		// 文件上传
		routerV1.POST("upload", v1.Upload)

		// 个人页展示
		routerV1.GET("profile/:id", v1.GetProfile)

		// 展示网站总访问量
		routerV1.GET("data/total", v1.GetTotalCount)
	}

	// admin后台相关接口
	// admin接口都需要鉴权
	routerAdmin := r.Group("api/v1/admin")
	routerAdmin.Use(middleware.JwtToken())
	{
		// 用户模块
		routerAdminUser := routerAdmin.Group("user")
		{
			// 删除某个用户、修改某个用户信息、查询所有用户信息
			routerAdminUser.DELETE("/:id", v1.DeleteUser)
			routerAdminUser.PUT("/:id", v1.EditUser)
			routerAdminUser.GET("/users", v1.GetUsers)
		}

		//修改密码
		routerAdmin.PUT("/changepw/:id", v1.ChangeUserPassword)

		// 分类模块
		routerAdminCate := routerAdmin.Group("category")
		{
			// 增加分类、删除分类、修改分类、查询分类
			routerAdminCate.POST("/add", v1.AddCategory)
			routerAdminCate.DELETE("/:id", v1.DeleteCate)
			routerAdminCate.PUT("/:id", v1.EditCate)
			routerAdminCate.GET("", v1.GetCateS)
		}

		// 文章模块
		routerAdminArticle := routerAdmin.Group("article")
		{
			// 添加文章、删除文章、修改文章、查询文章
			routerAdminArticle.POST("/add", v1.AddArticle)
			routerAdminArticle.DELETE("/:id", v1.DeleteArticle)
			routerAdminArticle.PUT("/:id", v1.EditArticle)
			routerAdminArticle.GET("", v1.GetArts)
			routerAdminArticle.GET("/info/:id", v1.GetOneArtInfo)
		}

		// 评论模块
		{
			routerAdmin.GET("comment/list", v1.GetCommentList)
			routerAdmin.DELETE("delcomment/:id", v1.DeleteComment)
			routerAdmin.PUT("checkcomment/:id", v1.CheckComment)
			routerAdmin.PUT("uncheckcomment/:id", v1.UnCheckComment)
		}

		// 个人页
		routerAdminProfile := routerAdmin.Group("profile")
		{
			// 获取个人页，更新个人页
			routerAdminProfile.GET("/:id", v1.GetProfile)
			routerAdminProfile.PUT("/:id", v1.UpdateProfile)
		}

		// 数据统计模块
		routerAdminData := routerAdmin.Group("data")
		routerAdminData.GET("infos", v1.GetVisitsInfos)
		routerAdminData.GET("total", v1.GetTotalCount)
	}

	// run
	r.Run(model.ServerConf.HttpPort)
}
