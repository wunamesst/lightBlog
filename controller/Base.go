package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//渲染模板
func display(c *gin.Context, templateName string, data interface{}) {
	c.HTML(http.StatusOK, templateName, data)
}