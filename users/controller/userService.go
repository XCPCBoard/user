package controller

import (
	"user/users/entity"
	"user/users/service"
)

//CreatUserController 新建账户
//@param account 账号
//@param keyword 密码
//@param email 邮箱
func CreatUserController(account string, keyword string, email string) (bool, error) {
	user := entity.User{
		Account: account,
		Keyword: keyword,
		Email:   email,
	}

	err := service.CreatUserInitService(&user)
	if err != nil {
		return false, err
	}
	return true, nil
}

//DeleteUserController 删除用户
//@param id 用户id
//@param account 用户账户
func DeleteUserController(id uint, account string) (bool, error) {
	user := entity.User{
		Id:      id,
		Account: account,
	}
	return service.DeleteUserService(&user)
}

//UpdateUserController 更新用户
//@param id 用户id
func UpdateUserController(id uint, user map[string]interface{}) (bool, error) {
	//检查是否包含主键
	if _, ok := user["id"]; ok {
		return false, nil
	}
	return service.UpdateUserService(id, user)
}

//SelectUserController 查询用户
//@param id 用户id
//@return map[string]interface{} 用户信息（除去了密码）
func SelectUserController(id uint) (map[string]interface{}, error) {
	user := map[string]interface{}{}
	err := service.SelectUserService(id, &user)
	if err != nil {
		return nil, err
	}
	return user, nil
}
