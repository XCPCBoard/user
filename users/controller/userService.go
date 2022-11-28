package controller

import (
	"user/users/entity"
	"user/users/service"
)

//CreatUser 新建账户
func CreatUser(account string, keyword string, email string) (bool, error) {
	user := entity.User{
		Account: account,
		Keyword: keyword,
		Email:   email,
	}

	err := service.CreatUserInit(&user)
	if err != nil {
		return false, err
	}
	return true, nil
}

func Delete(id uint, account string) (bool, error) {
	user := entity.User{
		Id:      id,
		Account: account,
	}
	return service.DeleteUser(&user)
}
