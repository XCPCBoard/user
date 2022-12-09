package util

import (
	"errors"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

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
