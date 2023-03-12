package service

import (
	"fmt"
	"github.com/XCPCBoard/common/dao"
	"github.com/XCPCBoard/common/errors"
	"github.com/XCPCBoard/common/logger"
	"github.com/XCPCBoard/user/entity"
)

//**********************************************************
//			    			CRUD						  //
//**********************************************************
/*
	这里没有create和delete的操作，因为UserService在创建时同步进行了
*/

// UpdateWebAccountWithOutNumService 更新用户网站账户(除了number相关信息)
// @param account 用户信息(需要包含主键)
func UpdateWebAccountWithOutNumService(account *entity.Accounts) *errors.MyError {

	if account.Rank != 0 || account.Total != 0 {
		return errors.CreateError(errors.ERROR.Code, "接口不支持插入rank和总过题数", account)
	}

	//注意Model里结构体
	res := dao.DBClient.Model(&entity.Accounts{}).
		Where("id = ?", account.Id).Updates(account)

	//error
	return CreatError(res, "更新用户网站账户出错",
		fmt.Sprintf("account id（user id):%v", account.Id))
}

// SelectWebAccountService 查询用户网站账户
// @param id 用户id
func SelectWebAccountService(id string, account *entity.Accounts) *errors.MyError {

	res := dao.DBClient.Model(&entity.Accounts{}).
		Find(account, id)
	//error
	return CreatError(res, "查询用户网站账号信息出错",
		fmt.Sprintf("user id:%v", id))

}

//**********************************************************
//			    			count						  //
//**********************************************************

// CountAllAccount 统计用户数量
func CountAllAccount(number *int64) error {
	res := dao.DBClient.Table(entity.AccountTableName).Count(number)
	if res.Error != nil {
		logger.L.Error("计算所有用户数量错误", res.Error, 0, "")
		return res.Error
	}
	return nil
}

// SelectMultipleAccount 查询多个用户
func SelectMultipleAccount(now int, pageSize int) (*[]entity.Accounts, error) {
	var accounts []entity.Accounts
	if res := dao.DBClient.Limit(pageSize).Offset(now).Find(&accounts); res.Error != nil {
		logger.L.Error("分页查询account错误", res.Error,
			0, fmt.Sprintf("now:%v,pageSize:%v", now, pageSize))
		return nil, res.Error
	}
	return &accounts, nil
}
