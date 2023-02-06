package entity

import (
	"errors"
	"github.com/XCPCBoard/common/config"
	"gorm.io/gorm"
	"time"
)

// User 用户变量
type User struct {
	Id              int       `gorm:"primaryKey;column:id" form:"id"`                  //type:*int         comment:用户id        version:2023-00-01 16:57
	Account         string    `gorm:"column:account" form:"account"`                   //type:string       comment:账号          version:2023-00-01 16:57
	Keyword         string    `gorm:"column:keyword" form:"keyword"`                   //type:string       comment:密码          version:2023-00-01 16:57
	Email           string    `gorm:"column:email" form:"email"`                       //type:string       comment:邮箱          version:2023-00-01 16:57
	IsAdministrator string    `gorm:"column:is_administrator" form:"is_administrator"` //type:string       comment:管理员标签    version:2023-00-01 16:57
	Name            string    `gorm:"column:name" form:"name"`                         //type:string       comment:姓名          version:2023-00-01 16:57
	ImagePath       string    `gorm:"column:image_path" form:"image_path"`             //type:string       comment:头像路径      version:2023-00-01 16:57
	Signature       string    `gorm:"column:signature" form:"signature"`               //type:string       comment:个性签名      version:2023-00-01 16:57
	PhoneNumber     string    `gorm:"column:phone_number" form:"phone_number"`         //type:string       comment:手机号        version:2023-00-01 16:57
	QQNumber        string    `gorm:"column:qq_number" form:"qq_number"`               //type:string       comment:qq号          version:2023-00-01 16:57
	CreatedAt       time.Time `gorm:"column:created_at" form:"created_at"`             //type:*time.Time   comment:创建时间      version:2023-00-01 16:57
	UpdatedAt       time.Time `gorm:"column:updated_at" form:"updated_at"`             //type:*time.Time   comment:更新时间      version:2023-00-01 16:57
}

// TableName GORM框架会自动检索结构体名复数形式的表名来对齐
// //若表名不是结构体的复数，则利用下面的代码来对齐表名
func (User) TableName() string {
	return UserTableName
}

// BeforeDelete 钩子函数
func (u *User) BeforeDelete(tx *gorm.DB) (err error) {
	if u.Account == config.Conf.Admin.Name { //超级管理员不能被删除
		return errors.New("admin user not allowed to delete")
	}
	return nil
}

func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	if u.Account == config.Conf.Admin.Name { //超管不能被更新
		return errors.New("admin user not allowed to update")
	}
	return
}
