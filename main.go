package main

import (
	"runtime"
	"path"
	"github.com/gin-gonic/gin"
	"fmt"
	"lightBlog/controller"
)

var CfgPath,UtilPath,ViewPath string

func main() {
	Init()

	CreateServ()
}

//初始化公用变量、常量等
func Init() {
	_, filename, _, _ := runtime.Caller(0)
	BasePath := path.Dir(filename)

	CfgPath  = path.Join(BasePath, "/config")
	UtilPath = path.Join(BasePath, "/util")
	ViewPath = path.Join(BasePath, "/view")

	fmt.Println(ViewPath)

}

//创建服务器
func CreateServ() {
	server := gin.Default()

	server.Static("/static", "./static")
	server.LoadHTMLGlob(ViewPath + "/*")

	server.GET("/", controller.Index)

	server.Run(":8080")
}