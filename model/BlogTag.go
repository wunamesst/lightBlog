package model

type BlogTag struct {
	Id int `xorm:"int(11) pk autoincr comment('主键ID')"`
	Tag string `xorm:"varchar(512) notnull default('') comment('标签名')"`
	Status int `xorm:"tinyint(2) default(1) comment('状态，0-禁用，1-启用')"`
	CreateAt string `xorm:"datetime created comment('创建时间')"`
	UpdateAt string `xorm:"datetime updated comment('更新时间')"`
}