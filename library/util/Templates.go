package util

import (
	"github.com/gin-gonic/gin"
	"path/filepath"
	"lightBlog/global"
	"html/template"
)

func SetTempates(engine *gin.Engine) {

	funcMap := template.FuncMap{

	}
	engine.SetFuncMap(funcMap)

	//设置模板定界符，默认是{{ }}
	engine.Delims("{{", "}}")
	engine.LoadHTMLGlob(filepath.Join(global.ViewPath, "./**/*"))
}