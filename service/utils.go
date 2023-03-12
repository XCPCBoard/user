package service

import (
	"crypto/md5"
	"fmt"
	"github.com/XCPCBoard/common/config"
	"github.com/XCPCBoard/common/errors"
	"github.com/XCPCBoard/common/logger"
	"github.com/XCPCBoard/user/entity"
	"golang.org/x/crypto/bcrypt"
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
		logger.L.Error(msg, res.Error, 1, para)
		rep := errors.INNER_ERROR
		rep.Data = "系统错误:" + msg
		return rep
	} else if res.RowsAffected == 0 {
		logger.L.Warn("can not find data:"+msg, 1, para)
		rep := errors.NOT_FOUND
		rep.Data = "未获取到数据:" + msg
		return rep
	}
	return nil
}

// GetPassWord 生成密码
func GetPassWord(keyword string) (string, *errors.MyError) {
	res := []byte(keyword + config.Conf.Secret)
	temp := fmt.Sprintf("%x", md5.Sum(res))
	logger.L.Debug("[ck]", 0, temp)

	hash, err := bcrypt.GenerateFromPassword([]byte(temp), bcrypt.DefaultCost)
	logger.L.Debug("[ck2]", 0, hash)
	logger.L.Debug("[ck3]", 0, len(hash))

	if err != nil {
		return "", errors.CreateError(errors.INNER_ERROR.Code, "获取加密密码失败", err)
	}
	return string(hash), nil
}

// ComparePasswords 判断密码是否一直
func ComparePasswords(keyword string, user *entity.Users) bool {
	res := []byte(keyword + config.Conf.Secret)

	temp := fmt.Sprintf("%x", md5.Sum(res))
	logger.L.Debug("[ck]", 0, temp)
	logger.L.Debug("[ck2]", 0, user.Keyword)
	logger.L.Debug("[ck3]", 0, user.Keyword)
	err := bcrypt.CompareHashAndPassword([]byte(user.Keyword), []byte(temp))
	return err == nil
}
