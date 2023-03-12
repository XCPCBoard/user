package service

import (
	"fmt"
	"github.com/XCPCBoard/common/dao"
	"github.com/XCPCBoard/common/errors"
	"github.com/XCPCBoard/common/logger"
	"github.com/XCPCBoard/user/entity"
	"gorm.io/gorm"
)

type UserService struct {
}

//**********************************************************
//			    			CRUD						  //
//**********************************************************

// CreateUserInitService 新建账户
// 同时会创建Account表中的数据
func CreateUserInitService(_account string, _keyword string, _email string, emailCode string) *errors.MyError {

	if len(_keyword) < 6 || len(_keyword) > 30 {
		e := errors.CreateError(errors.ERROR.Code, "密码长度错误", len(_keyword))
		logger.L.Err(e, 0)
		return e
	}

	user := &entity.Users{
		Account:         _account,
		Keyword:         _keyword,
		Email:           _email,
		IsAdministrator: 0,
	}
	var err2 *errors.MyError
	user.Keyword, err2 = GetPassWord(user.Keyword)
	if err2 != nil {
		logger.L.Err(err2, 0)
		return err2
	}

	account := &entity.Accounts{}

	err := dao.DBClient.Transaction(func(tx *gorm.DB) error {

		if creatUser := tx.Omit("id").Create(user); creatUser.Error != nil {
			return creatUser.Error
		}
		//查询刚刚插入的数据
		userX := &entity.Users{}
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
		logger.L.Error("Create Users and account failed", err, 0, fmt.Sprintf("user:%#v", user))
		return errors.GetError(errors.VALID_ERROR, "创建用户错误:"+err.Error())
	}
	return nil
}

// DeleteUserService 删除用户
// 同时会删除Account表中的数据
func DeleteUserService(user *entity.Users) *errors.MyError {

	//begin Transaction
	tx := dao.DBClient.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if tx.Error != nil {
		logger.L.Error("delete user failed", tx.Error, 0, fmt.Sprintf("user id:%d", user.Id))
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
	res = tx.Delete(&entity.Accounts{}, user.Id)
	if err := CreatError(res, "delete web account error",
		fmt.Sprintf("user id :%d", user.Id)); err != nil {
		tx.Rollback()
		return err
	}
	if com := tx.Commit(); com.Error != nil {
		logger.L.Error("删除用户时，提交事务错误", com.Error, 0, fmt.Sprintf("id:%d", user.Id))
	}
	return nil
}

// UpdateUserService 更新用户
// @param user 用户信息（注意不要包含主键）
// @param id 用户id
func UpdateUserService(id string, user *entity.Users) *errors.MyError {

	//强制为0
	user.Id = 0
	res := dao.DBClient.Model(&entity.Users{}).
		Where("id = ?", id).Updates(user)
	//error
	return CreatError(res, "Update user error", fmt.Sprintf("{user id:%v }", id))
}

// SelectUserServiceWithOutPassword 根据id查询用户（不包含密码)
// @param id 用户id
// @param user 用户数据指针
func SelectUserServiceWithOutPassword(id string, user map[string]interface{}) *errors.MyError {

	res := dao.DBClient.Model(&entity.Users{}).Find(user, id)
	//error
	if err := res.Error; err != nil {
		logger.L.Error("查询用户数据错误", res.Error, 0, fmt.Sprintf("user id:%v", id))
		rep := errors.INNER_ERROR
		rep.Data = "系统错误:" + "查询用户数据错误"
		return rep
	}

	//删掉密码
	delete(user, "keyword")
	return nil
}

//**********************************************************
//			    		其他条件查询						//
//**********************************************************

// SelectUserServiceByEmail 根据email查询用户（不包含密码)
// @param email 用户email
// @param user 用户数据指针
func SelectUserServiceByEmail(email string, user *entity.Users) *errors.MyError {

	res := dao.DBClient.Model(&entity.Users{}).Where("email = ?", email).Find(user)
	//error
	if err := res.Error; err != nil {
		logger.L.Error("根据email查询用户出错", res.Error, 0, fmt.Sprintf("user email:%v", email))
		rep := errors.INNER_ERROR
		rep.Data = "系统错误:" + "根据email查询用户"
		return rep
	}
	//不必删掉密码，还要对比
	return nil
}

// SelectUserServiceByName 模糊查询用户姓名
func SelectUserServiceByName(name string, users *[]entity.Users) *errors.MyError {
	res := dao.DBClient.Model(&entity.Users{}).Where("name LIKE ?", fmt.Sprintf("%v%%", name)).Find(users)
	if err := res.Error; err != nil {
		logger.L.Error("模糊查询用户姓名出错", res.Error, 0, fmt.Sprintf("user name:%v", name))
		rep := errors.INNER_ERROR
		rep.Data = "系统错误:" + "模糊查询用户姓名出错"
		return rep
	}
	for i := 0; i < len(*users); i++ {
		(*users)[i].Keyword = ""
	}

	return nil
}

// SelectUserServiceByAccount 模糊查询用户账户名
func SelectUserServiceByAccount(account string, users *[]entity.Users) *errors.MyError {
	res := dao.DBClient.Model(&entity.Users{}).Where("account LIKE ?", fmt.Sprintf("%v%%", account)).Find(users)
	if err := res.Error; err != nil {
		logger.L.Error("模糊查询用户账户名出错", res.Error, 0, fmt.Sprintf("user account:%v", account))
		rep := errors.INNER_ERROR
		rep.Data = "系统错误:" + "模糊查询用户账户名出错"
		return rep
	}
	for i := 0; i < len(*users); i++ {
		(*users)[i].Keyword = ""
	}
	return nil

}

//**********************************************************
//			    			Auth
//**********************************************************

// SelectUserIsAdmin 查询用户是否为管理员，是则返回true
func SelectUserIsAdmin(id string) (bool, error) {
	user := make(map[string]interface{})
	if err := SelectUserServiceWithOutPassword(id, user); err != nil {
		return false, err
	}
	if isAdmin, ok2 := user["is_administrator"]; !ok2 || fmt.Sprintf("%v", isAdmin) != "1" {
		if !ok2 {
			logger.L.Error("请求的用户id找不到is_administrator字段", nil, 0, fmt.Sprintf("id:%v", id))
		}
		return false, nil
	}
	//无需进log
	return true, nil

}

//**********************************************************
//			    			select all
//**********************************************************

func CountAllUser(number *int64) error {
	res := dao.DBClient.Table(entity.UserTableName).Count(number)
	if res.Error != nil {
		logger.L.Error("计算所有用户数量错误", res.Error, 0, "")
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
		logger.L.Error("查询用户是否存在错误", res.Error,
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
		logger.L.Error("查询用户是否存在错误", res.Error,
			0, fmt.Sprintf("email:%v", email))
		return false, res.Error
	}
	return temp != 0, nil
}
