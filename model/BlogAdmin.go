package model

type BlogAdmin struct{
	Id int `xorm:"int(11) pk autoincr comment('主键ID')"`
	Username string `xorm:"varchar(255) notnull default('') comment('用户名') index"`
	Password string `xorm:"char(32) notnull default('') comment('密码')"`
	Status int `xorm:"tinyint(2) default(1) comment('状态，0-禁用，1-启用')"`
	Group int `xorm:"int(11) notnull default(0) comment('用户组')"`
	CreateAt string `xorm:"datetime created comment('创建时间')"`
	UpdateAt string `xorm:"datetime updated comment('更新时间')"`
}