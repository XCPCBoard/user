package service

import (
	"errors"
	"fmt"
	"github.com/XCPCBoard/user/dao"
	"github.com/XCPCBoard/user/entity"
	"github.com/XCPCBoard/user/util"
	log "github.com/sirupsen/logrus"
)

/*
	这里没有create和delete的操作，因为UserService在创建时同步进行了
*/

//UpdateWebAccountService 更新用户网站账户 bool判断是否成功，error判断错误
//@param account 用户信息（注意不要包含主键）
//@param id 用户id
func UpdateWebAccountService(account map[string]interface{}) error {

	//检查是否包含主键
	if _, ok := account["id"]; !ok {
		err := errors.New(fmt.Sprintf("can't find account's id:%v", account["id"]))
		log.Errorf(err.Error())
		return err
	}
	//注意Model里结构体
	res := dao.DBClient.Model(&entity.Account{}).
		Where("id = ?", account["id"]).Updates(account)

	//error
	return util.CreatError(res,
		fmt.Sprintf("the account to be deleted could not be found:%v", account["id"]))
}

//SelectWebAccountService 查询用户网站账户
//@param id 用户id
func SelectWebAccountService(id string, account *map[string]interface{}) error {

	res := dao.DBClient.Model(&entity.Account{}).
		Find(account, id)
	//error
	return util.CreatError(res,
		fmt.Sprintf("the account to be deleted could not be found:%v", id))

}
