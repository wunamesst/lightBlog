package util

import (
	"log"
	"fmt"
	"crypto/md5"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CheckFatalError(title string, e error) {
	if nil != e {
		log.Fatal(title, e)
	}
}

//调试输出
func Debug(args ...interface{}) {
	fmt.Println(args...)
}

//MD5加密
func Md5(str string) string {
	data := md5.New()
	data.Write([]byte(str))
	cipherStr := data.Sum(nil)
	return hex.EncodeToString(cipherStr)
}

//渲染模板
func Display(c *gin.Context, templateName string, data interface{}) {
	c.HTML(http.StatusOK, templateName, data)
}
