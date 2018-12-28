package controller

import (
	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	display(c, "frontend/index.html", nil)
}