package controller

import (
	log "github.com/sirupsen/logrus"
	"strconv"
	"user/users/entity"
	"user/users/service"
)

//CreatUserController 新建账户
//@param account 账号
//@param keyword 密码
//@param email 邮箱
func CreatUserController(account string, keyword string, email string) error {
	user := entity.User{
		Account: account,
		Keyword: keyword,
		Email:   email,
	}

	return service.CreatUserInitService(&user)

}

//DeleteUserController 删除用户
//@param id 用户id
//@param account 用户账户
func DeleteUserController(id string, account string) error {
	uid, err := strconv.Atoi(id)
	if err != nil {
		log.Errorf(err.Error())
		return err
	}
	user := entity.User{
		Id:      uid,
		Account: account,
	}
	return service.DeleteUserService(&user)
}

//UpdateUserController 更新用户
//@param id 用户id
func UpdateUserController(user map[string]interface{}) error {

	return service.UpdateUserService(user)
}

//SelectUserController 查询用户
//@param id 用户id
//@return map[string]interface{} 用户信息（除去了密码）
func SelectUserController(id string) (map[string]interface{}, error) {

	user := map[string]interface{}{}
	err := service.SelectUserService(id, &user)
	if err != nil {
		return nil, err
	}
	return user, nil
}
