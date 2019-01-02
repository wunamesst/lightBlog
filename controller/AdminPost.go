package controller

import (
	"github.com/gin-gonic/gin"
	"lightBlog/library/helper"
	"lightBlog/model"
	. "lightBlog/global"
	"html/template"
)

//添加文章页面
func AdminPostsGet(c *gin.Context) {
	customStyle := []string{"posts.css"}

	display(c, "admin/post.html", gin.H{
		"customStyle": customStyle,
	})
}

//提交文章内容
func AdminPostsPost(c *gin.Context) {
	var postData struct {
		Title   string `form:"title" binding:"required"`
		Tag     string `form:"tag" binding:""`
		Content string `form:"content_html" binding:"required"`
		Status  int `form:"status" binding:"required"`
		Author  string `form:"author" binding:""`
	}

	if e := c.ShouldBind(&postData); e != nil {
		helper.Debug(e.Error())
		return
	}

	post := &model.BlogPost{
		Title: postData.Title,
		Content: template.HTMLEscapeString(postData.Content),
		Status: postData.Status,
		Author: postData.Author,
		CreateAt: helper.GetCurrentDate(),
	}
	helper.Debug(post)

	id, e := DB.InsertOne(post)
	if e != nil {
		helper.Debug(e.Error())
		return
	}

	helper.Debug(id)
}
