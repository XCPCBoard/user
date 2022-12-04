package service

import (
	log "github.com/sirupsen/logrus"
	"user/dao"
	"user/users/entity"
)

/*
	这里没有create和delete的操作，因为UserService在创建时同步进行了
*/

//UpdateWebAccountService 更新用户网站账户 bool判断是否成功，error判断错误
//@param account 用户信息（注意不要包含主键）
//@param id 用户id
func UpdateWebAccountService(id uint, account map[string]interface{}) (bool, error) {
	//防止包含主键导致数据库更新错误
	if _, ok := account["id"]; ok {
		account["id"] = id
	}
	//注意Model里结构体
	res := dao.DBClient.Model(&entity.Account{}).
		Where("id", id).Updates(account)

	//error
	if res.Error != nil {
		log.Errorf("function 'UpdateWebAccount' failed,  %v", res.Error)
		return false, res.Error
	} else if res.RowsAffected == 0 {
		//not find
		log.Errorf("the user to be deleted could not be found")
		return false, nil
	}
	return true, nil
}

//SelectWebAccountService 查询用户网站账户
//@param id 用户id
func SelectWebAccountService(id uint) (map[string]interface{}, error) {

	account := map[string]interface{}{}
	res := dao.DBClient.Model(&entity.Account{}).
		First(&account, id)
	//error
	if res.Error != nil {
		log.Errorf("function 'SelectWebAccountService' failed,  %v", res.Error)
		return nil, res.Error
	}
	return account, nil

}
