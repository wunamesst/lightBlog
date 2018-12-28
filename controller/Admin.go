package controller

import "github.com/gin-gonic/gin"

func AdminIndex(c *gin.Context) {
	display(c, "admin/index.html", nil)
}