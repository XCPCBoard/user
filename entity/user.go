package entity

import (
	"errors"
	"github.com/XCPCBoard/common/config"
	"gorm.io/gorm"
	"time"
)

// Users 用户变量
type Users struct {
	Id              int       `gorm:"column:id" json:"Id" form:"Id"`                                         //type:int         comment:用户id        version:2023-02-10 22:49
	Account         string    `gorm:"column:account" json:"Account" form:"Account"`                          //type:string      comment:账号          version:2023-02-10 22:49
	Keyword         string    `gorm:"column:keyword" json:"Keyword" form:"Keyword"`                          //type:string      comment:密码          version:2023-02-10 22:49
	Email           string    `gorm:"column:email" json:"Email" form:"Email"`                                //type:string      comment:邮箱          version:2023-02-10 22:49
	IsAdministrator int       `gorm:"column:is_administrator" json:"IsAdministrator" form:"IsAdministrator"` //type:int         comment:管理员标签    version:2023-02-10 22:49
	Name            string    `gorm:"column:name" json:"Name" form:"Name"`                                   //type:string      comment:姓名          version:2023-02-10 22:49
	ImagePath       string    `gorm:"column:image_path" json:"ImagePath" form:"ImagePath"`                   //type:string      comment:头像路径      version:2023-02-10 22:49
	Signature       string    `gorm:"column:signature" json:"Signature" form:"Signature"`                    //type:string      comment:个性签名      version:2023-02-10 22:49
	PhoneNumber     string    `gorm:"column:phone_number" json:"PhoneNumber" form:"PhoneNumber"`             //type:string      comment:手机号        version:2023-02-10 22:49
	QQNumber        string    `gorm:"column:qq_number" json:"QQNumber" form:"QQNumber"`                      //type:string      comment:qq号          version:2023-02-10 22:49
	CreatedAt       time.Time `gorm:"column:created_at" json:"CreatedAt" form:"CreatedAt"`                   //type:time.Time   comment:创建时间      version:2023-02-10 22:49
	UpdatedAt       time.Time `gorm:"column:updated_at" json:"UpdatedAt" form:"UpdatedAt"`                   //type:time.Time   comment:更新时间      version:2023-02-10 22:49
}

// TableName GORM框架会自动检索结构体名复数形式的表名来对齐
// //若表名不是结构体的复数，则利用下面的代码来对齐表名
func (Users) TableName() string {
	return UserTableName
}

// BeforeDelete 钩子函数
func (u *Users) BeforeDelete(tx *gorm.DB) (err error) {
	if u.Account == config.Conf.Admin.Name { //超级管理员不能被删除
		return errors.New("admin user not allowed to delete")
	}
	return nil
}

func (u *Users) BeforeUpdate(tx *gorm.DB) (err error) {
	if u.Account == config.Conf.Admin.Name { //超管不能被更新
		return errors.New("admin user not allowed to update")
	}
	return
}
