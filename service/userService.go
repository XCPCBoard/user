package service

import (
	"crypto/md5"
	"errors"
	"fmt"
	"github.com/XCPCBoard/user/dao"
	"github.com/XCPCBoard/user/entity"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"time"
)

const salt = "19491001"

//CreateUserInitService 新建账户
//同时会创建Account表中的数据
func CreateUserInitService(_account string, _keyword string, _email string) error {

	user := entity.User{
		Account: _account,
		Keyword: _keyword,
		Email:   _email,
	}
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
		log.Errorf("Create User Init failed,  %v", err)
		return err
	}
	return nil
}

//DeleteUserService 删除用户
//同时会删除Account表中的数据
func DeleteUserService(user *entity.User) error {

	//begin Transaction
	tx := dao.DBClient.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if tx.Error != nil {
		log.Errorf(" delete user failed,  %v", tx.Error)
		return tx.Error
	}

	//delete where id = ? and account = ?
	res := tx.Where("account", user.Account).Delete(user)

	//error
	if res.Error != nil {
		log.Errorf("delete user failed,  %v", res.Error)
		tx.Rollback()
		return res.Error
	} else if res.RowsAffected == 0 {
		//not find
		err := errors.New(fmt.Sprintf("the user to be deleted could not be found:%v", user.Id))
		log.Errorf(err.Error())
		tx.Rollback()
		return err
	}

	//delete webAccount
	res = tx.Delete(&entity.Account{}, user.Id)
	if res.Error != nil {
		log.Errorf("delete account failed,  %v", res.Error)
		tx.Rollback()
		return res.Error
	} else if res.RowsAffected == 0 {
		//not find
		err := errors.New(fmt.Sprintf("the account to be deleted could not be found:%v", user.Id))
		log.Errorf(err.Error())
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

//UpdateUserService 更新用户
//@param user 用户信息（注意不要包含主键）
//@param id 用户id
func UpdateUserService(user map[string]string) error {
	//检查是否包含主键
	if _, ok := user["id"]; !ok {
		err := errors.New("can't find user's id")
		log.Errorf(err.Error())
		return err
	}
	res := dao.DBClient.Model(&entity.User{}).
		Where("id = ?", user["id"]).Updates(user)

	//error
	return CreatError(res, fmt.Sprintf("the user to be deleted could not be found:%v", user["id"]))
}

//SelectUserService 查询用户
//@param id 用户id
//@param user 用户数据指针
func SelectUserService(id string, user *map[string]interface{}) error {

	res := dao.DBClient.Model(&entity.User{}).Find(user, id)
	//error
	err := CreatError(res, fmt.Sprintf("the user to be deleted could not be found:%v", id))

	//删掉密码
	if err == nil {
		delete(*user, "keyword")
	}

	return err

}

func SelectUserByPage() error {
	//var number int
	//if err := CountAllUser(&number); err != nil {
	//	return err
	//}
	//pageNumber := number / 100
	//go func() {
	//
	//}()
	//dao.DBClient.Scopes(Paginate()).Find(&users)
	return nil
}

func CountAllUser(number *int) error {
	res := dao.DBClient.Raw(fmt.Sprintf("select count(*) from %v", entity.UserTableName)).Scan(number)
	if res.Error != nil {
		log.Errorf(res.Error.Error())
		return res.Error
	}
	return nil
}
