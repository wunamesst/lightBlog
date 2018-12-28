package model

type BlogPost struct {
	Id int `xorm:"int(11) pk autoincr comment('主键ID')"`
	Title string `xorm:"varchar(512) notnull default('') comment('标题')"`
	Content string `xorm:"text comment('博客内容')"`
	Status int `xorm:"tinyint(2) default(1) comment('状态，0-禁用，1-启用')"`
	Author string `xorm:"varchar(255) notnull default('') comment('作者')"`
	CreateAt string `xorm:"datetime created comment('创建时间')"`
	UpdateAt string `xorm:"datetime updated comment('更新时间')"`
}