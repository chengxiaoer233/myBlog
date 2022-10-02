package main

import (
	"myBlog/dao"
	"myBlog/routes"
)

func main() {

	dao.Init()

	routes.InitRouters()
}
