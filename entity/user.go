package entity

import (
	"errors"
	"github.com/XCPCBoard/common/config"
	"gorm.io/gorm"
	"time"
)

//User 用户变量
type User struct {
	Id              int       `gorm:"primaryKey;column:id" json:"Id"`                 //type:*int         comment:用户id        version:2023-00-01 16:57
	Account         string    `gorm:"column:account" json:"Account"`                  //type:string       comment:账号          version:2023-00-01 16:57
	Keyword         string    `gorm:"column:keyword" json:"Keyword"`                  //type:string       comment:密码          version:2023-00-01 16:57
	Email           string    `gorm:"column:email" json:"Email"`                      //type:string       comment:邮箱          version:2023-00-01 16:57
	IsAdministrator string    `gorm:"column:is_administrator" json:"IsAdministrator"` //type:string       comment:管理员标签    version:2023-00-01 16:57
	Name            string    `gorm:"column:name" json:"Name"`                        //type:string       comment:姓名          version:2023-00-01 16:57
	ImagePath       string    `gorm:"column:image_path" json:"ImagePath"`             //type:string       comment:头像路径      version:2023-00-01 16:57
	Signature       string    `gorm:"column:signature" json:"Signature"`              //type:string       comment:个性签名      version:2023-00-01 16:57
	PhoneNumber     string    `gorm:"column:phone_number" json:"PhoneNumber"`         //type:string       comment:手机号        version:2023-00-01 16:57
	QQNumber        string    `gorm:"column:qq_number" json:"QqNumber"`               //type:string       comment:qq号          version:2023-00-01 16:57
	CreatedAt       time.Time `gorm:"column:created_at" json:"CreatedAt"`             //type:*time.Time   comment:创建时间      version:2023-00-01 16:57
	UpdatedAt       time.Time `gorm:"column:updated_at" json:"UpdatedAt"`             //type:*time.Time   comment:更新时间      version:2023-00-01 16:57
}

//TableName GORM框架会自动检索结构体名复数形式的表名来对齐
////若表名不是结构体的复数，则利用下面的代码来对齐表名
func (User) TableName() string {
	return UserTableName
}

//BeforeDelete 钩子函数
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
