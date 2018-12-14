package global

import (
	"github.com/go-xorm/xorm"
)

//全局路径
var (
	CfgPath, ViewPath string
)

//数据库对象
var DB *xorm.Engine

//配置
var Config struct{
	//Host
	Host string
	//数据库
	MySQL struct {
		Host string
		Port int
		Username string
		Password string
		Database string
	}
	//Redis
	Redis struct {
		Host string
		Port int
		Password string
		Database int
	}
}