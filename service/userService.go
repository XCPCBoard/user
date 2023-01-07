package service

import (
	"errors"
	"fmt"
	"github.com/XCPCBoard/common/dao"
	"github.com/XCPCBoard/user/entity"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"time"
)

type UserService struct {
}

//**********************************************************
//			    			CRUD
//**********************************************************

//CreateUserInitService 新建账户
//同时会创建Account表中的数据
func CreateUserInitService(_account string, _keyword string, _email string) error {

	user := &entity.User{
		Account: _account,
		Keyword: _keyword,
		Email:   _email,
	}
	//获取当前时间，为密码加盐
	user.CreatedAt = time.Now()
	user.Keyword = GetPassWord(user.Keyword, user)
	account := &entity.Account{}

	err := dao.DBClient.Transaction(func(tx *gorm.DB) error {

		if creatUser := tx.Omit("id").Create(user); creatUser.Error != nil {
			return creatUser.Error
		}
		//查询刚刚插入的数据
		userX := &entity.User{}
		check := tx.Where("account", user.Account).Find(userX)
		if check.Error != nil {
			tx.Rollback()
			return check.Error
		}
		account.Id = userX.Id
		if acc := tx.Create(account); acc.Error != nil {
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

	//delete where id = ?
	res := tx.Delete(user)

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
func UpdateUserService(id string, user *entity.User) error {

	res := dao.DBClient.Model(&entity.User{}).
		Where("id = ?", id).Updates(user)

	//error
	return CreatError(res, fmt.Sprintf("the user to be deleted could not be found:%v", id))
}

//SelectUserService 根据id查询用户（不包含密码)
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

//SelectUserServiceByEmail 根据email查询用户（不包含密码)
//@param email 用户email
//@param user 用户数据指针
func SelectUserServiceByEmail(email string, user *entity.User) error {

	res := dao.DBClient.Model(&entity.User{}).Where("email = ?", email).Find(user)
	//error
	err := CreatError(res, fmt.Sprintf("the user to be deleted could not be found by email:%v", email))

	//不必删掉密码，还要对比
	return err
}

//**********************************************************
//			    			Auth
//**********************************************************

//SelectUserIsAdmin 查询用户是否为管理员，是则返回true
func SelectUserIsAdmin(id string) (bool, error) {
	user := map[string]interface{}{}
	if err := SelectUserService(id, &user); err != nil {
		return false, err
	}
	if isAdmin, ok2 := user["is_administrator"]; !ok2 || fmt.Sprintf("%v", isAdmin) != "1" {
		if !ok2 {
			log.Println("请求的用户id找不到is_administrator字段，id:%v", ok2)
		}
		return false, nil
	}
	//无需进log
	return true, nil

}

//**********************************************************
//			    			select all
//**********************************************************

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
