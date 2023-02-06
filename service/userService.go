package service

import (
	"fmt"
	"github.com/XCPCBoard/common/dao"
	"github.com/XCPCBoard/common/errors"
	"github.com/XCPCBoard/common/logger"
	"github.com/XCPCBoard/user/entity"
	"gorm.io/gorm"
	"time"
)

type UserService struct {
}

//**********************************************************
//			    			CRUD						  //
//**********************************************************

// CreateUserInitService 新建账户
// 同时会创建Account表中的数据
func CreateUserInitService(_account string, _keyword string, _email string) *errors.MyError {

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
		//网站id=user id
		account.Id = userX.Id
		if acc := tx.Create(account); acc.Error != nil {
			tx.Rollback()
			return acc.Error
		}
		return nil
	})

	if err != nil {
		logger.Logger.Error("Create User and account failed", err, 0, fmt.Sprintf("user:%#v", user))
		return errors.GetError(errors.VALID_ERROR, "创建用户错误:"+err.Error())
	}
	return nil
}

// DeleteUserService 删除用户
// 同时会删除Account表中的数据
func DeleteUserService(user *entity.User) *errors.MyError {

	//begin Transaction
	tx := dao.DBClient.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if tx.Error != nil {
		logger.Logger.Error("delete user failed", tx.Error, 0, fmt.Sprintf("user id:%d", user.Id))
		return errors.GetError(errors.INNER_ERROR, "删除用户错误")
	}

	//delete where id = ?
	res := tx.Delete(user)

	//error
	if err := CreatError(res, "delete user error",
		fmt.Sprintf("user id:%d", user.Id)); err != nil {
		tx.Rollback()
		return err
	}

	//delete webAccount
	res = tx.Delete(&entity.Account{}, user.Id)
	if err := CreatError(res, "delete web account error",
		fmt.Sprintf("user id :%d", user.Id)); err != nil {
		tx.Rollback()
		return err
	}
	if com := tx.Commit(); com.Error != nil {
		logger.Logger.Error("删除用户时，提交事务错误", com.Error, 0, fmt.Sprintf("id:%d", user.Id))
	}
	return nil
}

// UpdateUserService 更新用户
// @param user 用户信息（注意不要包含主键）
// @param id 用户id
func UpdateUserService(id string, user *entity.User) *errors.MyError {

	//强制为0
	user.Id = 0
	res := dao.DBClient.Model(&entity.User{}).
		Where("id = ?", id).Updates(user)
	//error
	return CreatError(res, "Update user error", fmt.Sprintf("{user id:%v }", id))
}

// SelectUserServiceWithOutPassword 根据id查询用户（不包含密码)
// @param id 用户id
// @param user 用户数据指针
func SelectUserServiceWithOutPassword(id string, user *map[string]interface{}) *errors.MyError {

	res := dao.DBClient.Model(&entity.User{}).Find(user, id)
	//error
	err := CreatError(res, "select User Service error", fmt.Sprintf("user id:%v", id))

	//删掉密码
	if err == nil {
		delete(*user, "keyword")
	}

	return err
}

// SelectUserServiceByEmail 根据email查询用户（不包含密码)
// @param email 用户email
// @param user 用户数据指针
func SelectUserServiceByEmail(email string, user *entity.User) *errors.MyError {

	res := dao.DBClient.Model(&entity.User{}).Where("email = ?", email).Find(user)
	//error
	err := CreatError(res, "select User Service by email error", fmt.Sprintf("user email:%v", email))

	//不必删掉密码，还要对比
	return err
}

//**********************************************************
//			    			Auth
//**********************************************************

// SelectUserIsAdmin 查询用户是否为管理员，是则返回true
func SelectUserIsAdmin(id string) (bool, error) {
	user := map[string]interface{}{}
	if err := SelectUserServiceWithOutPassword(id, &user); err != nil {
		return false, err
	}
	if isAdmin, ok2 := user["is_administrator"]; !ok2 || fmt.Sprintf("%v", isAdmin) != "1" {
		if !ok2 {
			logger.Logger.Error("请求的用户id找不到is_administrator字段", nil, 0, fmt.Sprintf("id:%v", id))
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

func CountAllUser(number *int64) error {
	res := dao.DBClient.Table(entity.UserTableName).Count(number)
	if res.Error != nil {
		logger.Logger.Error("计算所有用户数量错误", res.Error, 0, "")
		return res.Error
	}
	return nil
}

//**********************************************************//
//			    			验证用户是否存在					//
//**********************************************************//

func UserAccountIsExist(account string) (bool, error) {
	//提前用where是因为gorm通过参数防止sql注入
	temp := 0
	res := dao.DBClient.Where("account=?", account).
		Raw(fmt.Sprintf("select 1 from %v limit 1", entity.UserTableName)).Scan(&temp)
	if res.Error != nil {
		logger.Logger.Error("查询用户是否存在错误", res.Error,
			0, fmt.Sprintf("account:%v", account))
		return false, res.Error
	}
	return temp == 0, nil
}

func UserEmailIsExist(email string) (bool, error) {
	//提前用where是因为gorm通过参数防止sql注入
	temp := 0
	res := dao.DBClient.Where("email=?", email).
		Raw(fmt.Sprintf("select 1 from %v limit 1", entity.UserTableName)).Scan(&temp)
	if res.Error != nil {
		logger.Logger.Error("查询用户是否存在错误", res.Error,
			0, fmt.Sprintf("email:%v", email))
		return false, res.Error
	}
	return temp != 0, nil
}
