package controller

import (
	"user/users/service"
)

//UpdateWebAccountController 更新用户网站账户 bool判断是否成功，error判断错误
//@param account 用户信息（注意不要包含主键）
//@param id 用户id
func UpdateWebAccountController(id uint, account map[string]interface{}) (bool, error) {
	//检查是否包含主键
	if _, ok := account["id"]; ok {
		return false, nil
	}
	return service.UpdateWebAccountService(id, account)
}

//SelectWebAccountController 查询用户网站账户
//@param id 用户id
func SelectWebAccountController(id uint) (map[string]interface{}, error) {
	return service.SelectWebAccountService(id)
}
