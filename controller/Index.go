package controller

import (
	"github.com/gin-gonic/gin"
	"lightBlog/library/util"
)

func Index(c *gin.Context) {
	util.Display(c, "index.html", nil)
}