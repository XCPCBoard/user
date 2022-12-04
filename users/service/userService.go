package service

import (
	"crypto/md5"
	"fmt"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"time"
	"user/dao"
	"user/users/entity"
)

const salt = "19491001"

//CreatUserInitService 新建账户
//同时会创建Account表中的数据
func CreatUserInitService(user *entity.User) error {

	//获取当前时间，为密码加盐
	user.CreatedAt = time.Now()
	keyword := []byte(user.Keyword + user.CreatedAt.String() + salt)
	user.Keyword = fmt.Sprintf("%x", md5.Sum(keyword))
	account := entity.Account{}

	err := dao.DBClient.Transaction(func(tx *gorm.DB) error {

		if creatUser := tx.Create(user); creatUser.Error != nil {
			return creatUser.Error
		}
		//查询刚刚插入的数据
		userX := entity.User{}
		check := tx.Where("account", user.Account).Find(&userX)
		if check.Error != nil {
			tx.Rollback()
			return check.Error
		}
		account.Id = userX.Id
		if acc := tx.Create(&account); acc.Error != nil {
			tx.Rollback()
			return acc.Error
		}
		return nil
	})

	if err != nil {
		log.Errorf("function 'CreatUserInitService' failed,  %v", err)
		return err
	}
	return nil
}

//DeleteUserService 删除用户 bool判断是否成功，error判断错误
//同时会删除Account表中的数据
func DeleteUserService(user *entity.User) (bool, error) {

	//begin Transaction
	tx := dao.DBClient.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if tx.Error != nil {
		log.Errorf("function 'DeleteUserService' Transaction failed,  %v", tx.Error)
		return false, tx.Error
	}

	//delete where id = ? and account = ?
	res := tx.Where("account", user.Account).Delete(user)

	//error
	if res.Error != nil {
		log.Errorf("function 'DeleteUserService' delete user failed,  %v", res.Error)
		tx.Rollback()
		return false, res.Error
	} else if res.RowsAffected == 0 {
		//not find
		log.Errorf("the user to be deleted could not be found")
		tx.Rollback()
		return false, nil
	}

	//delete webAccount
	res = tx.Delete(&entity.Account{}, user.Id)
	if res.Error != nil {
		log.Errorf("function 'DeleteUserService' delete account failed,  %v", res.Error)
		tx.Rollback()
		return false, res.Error
	} else if res.RowsAffected == 0 {
		//not find
		log.Errorf("the account to be deleted could not be found")
		tx.Rollback()
		return false, nil
	}

	return true, tx.Commit().Error
}

//UpdateUserService 更新用户bool判断是否成功，error判断错误
//@param user 用户信息（注意不要包含主键）
//@param id 用户id
func UpdateUserService(id uint, user map[string]interface{}) (bool, error) {

	//防止包含主键导致数据库更新错误
	if _, ok := user["id"]; ok {
		user["id"] = id
	}
	res := dao.DBClient.Model(&entity.User{}).
		Where("id", id).Updates(user)

	//error
	if res.Error != nil {
		log.Errorf("function 'UpdateUserService' failed,  %v", res.Error)
		return false, res.Error
	} else if res.RowsAffected == 0 {
		//not find
		log.Errorf("the user to be deleted could not be found")
		return false, nil
	}
	return true, nil
}

//SelectUserService 查询用户
//@param id 用户id
//@param user 用户数据指针
func SelectUserService(id uint, user *map[string]interface{}) error {

	res := dao.DBClient.Model(&entity.User{}).First(user, id)
	//error
	if res.Error != nil {
		log.Errorf("function 'SelectUserService' failed,  %v", res.Error)
		return res.Error
	}

	//删掉密码
	delete(*user, "keyword")

	return nil

}
