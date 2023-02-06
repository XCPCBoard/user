package service

import (
	"crypto/md5"
	"fmt"
	"github.com/XCPCBoard/common/config"
	"github.com/XCPCBoard/common/errors"
	"github.com/XCPCBoard/common/logger"
	"github.com/XCPCBoard/user/entity"
	"gorm.io/gorm"
)

// Paginate 分页查询
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

// CreatError 生成错误信息
// @param res sql结果
// @param msg 未找到数据时的错误信息
func CreatError(res *gorm.DB, msg string, para string) *errors.MyError {
	if res.Error != nil {
		logger.Logger.Error(msg, res.Error, 1, para)
		rep := errors.INNER_ERROR
		rep.Data = "系统错误:" + msg
		return rep
	} else if res.RowsAffected == 0 {
		logger.Logger.Warn("can not find data:"+msg, 1, para)
		rep := errors.VALID_ERROR
		rep.Data = "参数错误:" + msg
		return rep
	}
	return nil
}

// GetPassWord 获取加密后的密码
func GetPassWord(keyword string, user *entity.User) string {
	res := []byte(keyword + user.CreatedAt.String() + config.Conf.Secret)
	return fmt.Sprintf("%x", md5.Sum(res))
}
