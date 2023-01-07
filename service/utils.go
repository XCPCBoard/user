package service

import (
	"crypto/md5"
	"errors"
	"fmt"
	"github.com/XCPCBoard/common/config"
	"github.com/XCPCBoard/user/entity"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

//Paginate 分页查询
func Paginate(page int, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page == 0 {
			page = 1
		}
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}
		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

//CreatError 生成错误信息，检查是否未找到数据
//@param res sql结果
//@param msg 未找到数据时的错误信息
func CreatError(res *gorm.DB, msg string) error {
	if res.Error != nil {
		log.Errorf(res.Error.Error())
		return res.Error
	} else if res.RowsAffected == 0 {
		err := errors.New(msg)
		log.Errorf(err.Error())
		return err
	}
	return nil
}

//GetPassWord 获取加密后的密码
func GetPassWord(keyword string, user *entity.User) string {
	res := []byte(keyword + user.CreatedAt.String() + config.Conf.Secret)
	return fmt.Sprintf("%x", md5.Sum(res))
}
