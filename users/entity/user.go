package entity

import (
	"errors"
	"gorm.io/gorm"
	"time"
	"user/config"
)

const userTableName = "users"

//User 用户变量
type User struct {
	Id int //用户id

	Account         string //账号
	Keyword         string //密码
	Email           string //邮箱
	IsAdministrator bool   //管理员标签(=1时为管理员)

	Name        string //姓名
	ImagePath   string //头像的路径
	Signature   string //个性签名
	PhoneNumber string //手机号
	QQNumber    string //qq号

	CreatedAt time.Time //创建时间
	UpdatedAt time.Time //更新时间
}

//TableName GORM框架会自动检索结构体名复数形式的表名来对齐
////若表名不是结构体的复数，则利用下面的代码来对齐表名
//func (User) TableName() string {
//	return userTableName
//}

//BeforeDelete 钩子函数
func (u *User) BeforeDelete(tx *gorm.DB) (err error) {
	if u.Account == config.Conf.AdminName { //超级管理员不能被删除
		return errors.New("admin user not allowed to delete")
	}
	return nil
}

func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	if u.Account == config.Conf.AdminName { //超管不能被更新
		return errors.New("admin user not allowed to update")
	}
	return
}
