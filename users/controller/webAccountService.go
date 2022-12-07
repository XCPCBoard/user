package controller

import (
	"user/users/service"
)

//UpdateWebAccountController 更新用户网站账户 bool判断是否成功，error判断错误
//@param account 用户信息（注意不要包含主键）
//@param id 用户id
func UpdateWebAccountController(id string, account map[string]interface{}) error {

	return service.UpdateWebAccountService(account)
}

//SelectWebAccountController 查询用户网站账户
//@param id 用户id
func SelectWebAccountController(id string) (map[string]interface{}, error) {
	user := map[string]interface{}{}
	err := service.SelectWebAccountService(id, &user)
	if err != nil {
		return nil, err
	}
	return user, nil
}
