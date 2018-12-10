package util

import (
	"log"
	"fmt"
	"crypto/md5"
	"encoding/hex"
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

func Md5(str string) string {
	data := md5.New()
	data.Write([]byte(str))
	cipherStr := data.Sum(nil)
	return hex.EncodeToString(cipherStr)
}