package service

import (
	"errors"
	"fmt"
	"github.com/XCPCBoard/common/dao"
	"github.com/XCPCBoard/user/entity"
	log "github.com/sirupsen/logrus"
)

/*
	这里没有create和delete的操作，因为UserService在创建时同步进行了
*/

//UpdateWebAccountService 更新用户网站账户 bool判断是否成功，error判断错误
//@param account 用户信息（注意不要包含主键）
//@param id 用户id
func UpdateWebAccountService(account *entity.Account) error {

	//检查主键是否为0
	if account.Id == 0 {
		err := errors.New(fmt.Sprintf("account's id should not be 0"))
		log.Errorf(err.Error())
		return err
	}
	//注意Model里结构体
	res := dao.DBClient.Model(&entity.Account{}).
		Where("id = ?", account.Id).Updates(account)

	//error
	return CreatError(res,
		fmt.Sprintf("the account to be deleted could not be found:%v", account.Id))
}

//SelectWebAccountService 查询用户网站账户
//@param id 用户id
func SelectWebAccountService(id string, account *map[string]interface{}) error {

	res := dao.DBClient.Model(&entity.Account{}).
		Find(account, id)
	//error
	return CreatError(res,
		fmt.Sprintf("the account to be deleted could not be found:%v", id))

}
