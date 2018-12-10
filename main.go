package main

import (
	"runtime"
	"path"
	"github.com/gin-gonic/gin"
	"lightBlog/controller"
	"lightBlog/library/util"
)

var (
	CfgPath, ViewPath string
)

func main() {
	Init()

	CreateServer()
}

//初始化公用变量、常量等
func Init() {
	_, filename, _, _ := runtime.Caller(0)
	BasePath := path.Dir(filename)

	CfgPath = path.Join(BasePath, "/config/")
	ViewPath = path.Join(BasePath, "/view/")

	util.Debug(ViewPath)
}

//创建服务器
func CreateServer() {
	server := gin.Default()

	//访问/static时，资源指向static路径
	server.Static("/static", "./static")
	//模块加载路径
	server.LoadHTMLGlob(ViewPath + "/*")

	//路由定义
	server.GET("/", controller.Index)

	server.Run(":8080")
}
