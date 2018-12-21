package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-xorm/xorm"
	_ "github.com/go-sql-driver/mysql"

	"path"
	"io/ioutil"
	"encoding/json"
	"fmt"

	"lightBlog/controller"
	"lightBlog/library/util"
	"lightBlog/global"
	"time"
	"lightBlog/model"
	"runtime"
	"path/filepath"
	"lightBlog/library/helper"
)

func main() {
	InitCommonPath()
	InitDatabase()
	//Install()
	InitServer()
}

//初始化公用变量、常量等
func InitCommonPath() {
	global.CfgPath = path.Join(GetBasePath(), "./config")
	global.ViewPath = path.Join(GetBasePath(), "./view")
}

//创建服务器
func InitServer() {
	gin.SetMode(gin.DebugMode)
	router := gin.Default()

	//访问/static时，资源指向static路径
	router.Static("/static", "./static")
	//自动恢复
	//router.Use(gin.Recovery())

	//模块加载路径
	util.SetTempates(router)
	//路由定义
	router.GET("/", controller.Index)
	router.GET("/admin", controller.AdminIndex)

	router.Run(global.Config.Host)
}

//初始化数据库连接
func InitDatabase() {
	helper.Debug(global.CfgPath)
	bt, e := ioutil.ReadFile(global.CfgPath + "/config.json")
	helper.CheckFatalError("读取配置文件错误", e)

	e = json.Unmarshal(bt, &global.Config)
	helper.CheckFatalError("配置文件格式错误", e)

	//连接数据库
	global.DB, e = xorm.NewEngine("mysql",
		fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8",
			global.Config.MySQL.Username,
			global.Config.MySQL.Password,
			global.Config.MySQL.Host,
			global.Config.MySQL.Port,
			global.Config.MySQL.Database))
	helper.CheckFatalError("连接数据库错误", e)

	global.DB.DatabaseTZ = time.Local
	//global.DB.SetMapper(core.SnakeMapper{})
	global.DB.ShowSQL(true)
	global.DB.ShowExecTime(true)

	e = global.DB.Ping()
	helper.CheckFatalError("数据库连接丢失", e)
}

//测试安装
func Install() {
	var tables = []interface{}{
		new(model.BlogAdmin),
	}

	global.DB.DropTables(tables...)
	e := global.DB.Sync2(tables...)
	helper.CheckFatalError("创建数据库失败", e)
}

//项目路径
func GetBasePath() string {
	_, filename, _, _ := runtime.Caller(0)
	basePath := path.Dir(filename)

	return basePath
}
