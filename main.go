package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-xorm/xorm"
	_ "github.com/go-sql-driver/mysql"

	"runtime"
	"path"
	"io/ioutil"
	"encoding/json"
	"fmt"

	"lightBlog/controller"
	"lightBlog/library/util"
	"lightBlog/global"
	"time"
	"lightBlog/model"
)

func main() {
	InitCommonPath()
	InitDatabase()
	//Install()
	InitServer()
}

//初始化公用变量、常量等
func InitCommonPath() {
	_, filename, _, _ := runtime.Caller(0)
	BasePath := path.Dir(filename)

	global.CfgPath = path.Join(BasePath, "/config")
	global.ViewPath = path.Join(BasePath, "/view")
}

//创建服务器
func InitServer() {
	server := gin.Default()

	//访问/static时，资源指向static路径
	server.Static("/static", "./static")
	//模块加载路径
	server.LoadHTMLGlob(global.ViewPath + "/*")
	//自动恢复
	server.Use(gin.Recovery())

	//路由定义
	server.GET("/", controller.Index)

	server.Run(global.Config.Host)
}

//初始化数据库连接
func InitDatabase() {
	util.Debug(global.CfgPath)
	bt, e := ioutil.ReadFile(global.CfgPath + "/config.json")
	util.CheckFatalError("读取配置文件错误", e)

	e = json.Unmarshal(bt, &global.Config)
	util.CheckFatalError("配置文件格式错误", e)

	//连接数据库
	util.Debug(global.Config.MySQL.Host)
	global.DB, e = xorm.NewEngine("mysql",
		fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8",
			global.Config.MySQL.Username,
			global.Config.MySQL.Password,
			global.Config.MySQL.Host,
			global.Config.MySQL.Port,
			global.Config.MySQL.Database))
	util.CheckFatalError("连接数据库错误", e)

	global.DB.DatabaseTZ = time.Local
	//global.DB.SetMapper(core.SnakeMapper{})
	global.DB.ShowSQL(true)
	global.DB.ShowExecTime(true)

	e = global.DB.Ping()
	util.CheckFatalError("数据库连接丢失", e)
}

//测试安装
func Install() {
	var tables = []interface{}{
		new(model.BlogAdmin),
	}

	global.DB.DropTables(tables...)
	e := global.DB.Sync2(tables...)
	util.CheckFatalError("创建数据库失败", e)
}