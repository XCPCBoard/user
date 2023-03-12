package service

import (
	"fmt"
	"github.com/XCPCBoard/common/errors"
	"github.com/XCPCBoard/common/logger"
	"github.com/XCPCBoard/user/entity"
)

// UserLoginService 用户登录服务
func UserLoginService(keyword, email string, user *entity.Users) *errors.MyError {
	//搜素用户
	if ok, err := UserEmailIsExist(email); err != nil {
		e := errors.CreateError(errors.ERROR.Code, "查询用户是否存在时出错", email)
		logger.L.Err(e, 0)
		return e
	} else if !ok {
		e := errors.CreateError(errors.ERROR.Code, "用户不存在:", email)
		logger.L.Err(e, 0)
		return e
	}

	//找到用户数据
	if err := SelectUserServiceByEmail(email, user); err != nil {
		logger.L.Err(err, 0)
		return err
	}

	//对比密码
	if ComparePasswords(keyword, user) == false {

		//info 级别
		logger.L.Info("用户密码错误", 0,
			fmt.Sprintf("email:%v,keyword（encodered):%v,user:%+v", email, keyword, user))
		return errors.GetError(errors.VALID_ERROR, "密码错误")
	}
	return nil
}

// UserRegister 用户注册
func UserRegister(account, _keyword, _email, emailCode string) *errors.MyError {

	//验证邮箱
	if err := VerifyEmailCode(RegisterOption, _email, emailCode); err != nil {
		logger.L.Error(err.Msg, err, 0, fmt.Sprintf("email:%v", _email))
		return err
	}

	if err := CreateUserInitService(account, _keyword, _email, emailCode); err != nil {
		logger.L.Err(err, 0)
		return err
	}

	return nil

}

// UserResetPasswordCheck 检查用户重置密码
func UserResetPasswordCheck(keyword, email, emailCode string) (*entity.Users, *errors.MyError) {
	//验证密码格式
	if len(keyword) < 6 || len(keyword) > 30 {
		return nil, errors.GetError(errors.VALID_ERROR, "密码长度有误")
	}

	//验证邮箱验证码
	if err := VerifyEmailCode(RegisterOption, email, emailCode); err != nil {
		logger.L.Error(err.Msg, err, 0, fmt.Sprintf("email:%v", email))
		return nil, err
	}
	//搜索用户
	user := entity.Users{}
	if err := SelectUserServiceByEmail(email, &user); err != nil {
		return nil, err
	}
	return &user, nil
}
