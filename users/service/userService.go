package service

import (
	log "github.com/sirupsen/logrus"
	"user/dao"
	"user/users/entity"
)

//CreatUserInit 新建账户
func CreatUserInit(user *entity.User) error {
	//user.CreatedAt=
	res := dao.DBClient.Create(&user)
	if res.Error != nil {
		log.Errorf("function 'InsertUserInit' failed,  %v", res.Error)
		return res.Error
	}
	return nil
}

//DeleteUser 删除用户 bool判断是否成功，error判断错误
func DeleteUser(user *entity.User) (bool, error) {

	res := dao.DBClient.Where("account", user.Account).Delete(&user)
	//error
	if res.Error != nil {
		log.Errorf("function 'DeleteUser' failed,  %v", res.Error)
		return false, res.Error
	} else if res.RowsAffected == 0 {
		//not find
		log.Errorf("the user to be deleted could not be found")
		return false, nil
	}
	return true, nil
}

//UpdateUser 更新用户 bool判断是否成功，error判断错误
func UpdateUser(user map[string]interface{}) (bool, error) {
	delete(user, "id")
	res := dao.DBClient.Where("account", user["account"]).Updates(user)

	//error
	if res.Error != nil {
		log.Errorf("function 'UpdateUser' failed,  %v", res.Error)
		return false, res.Error
	} else if res.RowsAffected == 0 {
		//not find
		log.Errorf("the user to be deleted could not be found")
		return false, nil
	}
	return true, nil
}
