package controller

import (
	"github.com/gin-gonic/gin"
	"lightBlog/library/helper"
	"lightBlog/model"
	. "lightBlog/global"
)

func AdminPostsGet(c *gin.Context) {
	customStyle := []string{"posts.css"}

	display(c, "admin/post.html", gin.H{
		"customStyle": customStyle,
	})
}

func AdminPostsPost(c *gin.Context) {
	var postData struct {
		Title   string `form:"title" binding:"required"`
		Tag     string `form:"tag" binding:""`
		Content string `form:"content" binding:"required"`
		Status  int `form:"status" binding:"required"`
		Author  string `form:"author" binding:""`
	}

	if e := c.ShouldBind(&postData); e != nil {
		helper.Debug(e.Error())
		return
	}

	post := &model.BlogPost{
		Title: postData.Title,
		Content: postData.Content,
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
