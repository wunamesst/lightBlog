package helper

import (
	"log"
	"fmt"
	"crypto/md5"
	"encoding/hex"
	"time"
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

func GetCurrentDate() string {
	return time.Now().Format("2006-01-02 15:04:05")
}
