package service

import (
	"fmt"
	"github.com/XCPCBoard/common/dao"
	"github.com/XCPCBoard/common/errors"
	"github.com/XCPCBoard/common/logger"
	"github.com/XCPCBoard/user/entity"
	"sync"
)

//**********************************************************
//			    			CRUD						  //
//**********************************************************
/*
	这里没有create和delete的操作，因为UserService在创建时同步进行了
*/

// UpdateWebAccountService 更新用户网站账户 bool判断是否成功，error判断错误
// @param account 用户信息（注意不要包含主键）
// @param id 用户id
func UpdateWebAccountService(account *entity.Account) *errors.MyError {

	////检查主键是否为0
	//if account.Id == 0 {
	//	//err := errors.New(fmt.Sprintf("account's id should not be 0"))
	//	logger.Logger.Error("账户id不能为0", nil, 0, fmt.Sprintf("account :%+v", account))
	//	return err
	//}

	//注意Model里结构体
	res := dao.DBClient.Model(&entity.Account{}).
		Where("id = ?", account.Id).Updates(account)

	//error
	return CreatError(res, "更新用户网站账户出错",
		fmt.Sprintf("account id（user id):%v", account.Id))
}

// SelectWebAccountService 查询用户网站账户
// @param id 用户id
func SelectWebAccountService(id string, account *entity.Account) *errors.MyError {

	res := dao.DBClient.Model(&entity.Account{}).
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
		logger.Logger.Error("计算所有用户数量错误", res.Error, 0, "")
		return res.Error
	}
	return nil
}

// SelectMultipleAccount 查询多个用户
func SelectMultipleAccount(now int, pageSize int) (*[]entity.Account, error) {
	var accounts []entity.Account
	if res := dao.DBClient.Limit(pageSize).Offset(now).Find(&accounts); res.Error != nil {
		logger.Logger.Error("分页查询account错误", res.Error,
			0, fmt.Sprintf("now:%v,pageSize:%v", now, pageSize))
		return nil, res.Error
	}
	return &accounts, nil
}

//**********************************************************
//			    			handle						  //
//**********************************************************

// handleMultipleAccountData 处理用户数据
func handleMultipleAccountData(pageSize int, now int, handler func(user *entity.Account) *errors.MyError) *errors.MyError {
	var err *errors.MyError
	var wg sync.WaitGroup

	//获取数据
	accounts, e := SelectMultipleAccount(now, pageSize)
	if e != nil {
		resErr := errors.NewError(500, "分页查询account错误")
		resErr.Data = e.Error()
		return resErr
	}

	//并发处理
	for _, account := range *accounts {
		wg.Add(1)

		go func(ac *entity.Account) {
			//捕获可能在调用逻辑中发生的panic & wg.Down
			defer func() {
				if e := recover(); e != nil {
					logger.Logger.Error("handler中发生了panic", nil, 0, nil)
				}
				defer wg.Done()
			}()

			// 取最后一个报错的handler调用逻辑，并最终向外返回
			e := handler(ac)
			if e != nil {
				logger.Logger.Error(e.Msg, e, 0,
					fmt.Sprintf("处理account错误 data:%#v", e.Data))
				err = e
			}
		}(&account)
	}
	wg.Wait()
	return err
}

// HandleAllAccount	分批处理全部用户
//
//	@param	pageSize	每批循环处理的用户数，pageSize=10时每批处理10个用户，pageSize<1时直接处理全部用户
//	@param	handler		回调函数
func HandleAllAccount(pageSize int, handler func(user *entity.Account) *errors.MyError) *errors.MyError {
	//获取用户数量
	var sum int64
	if err := CountAllAccount(&sum); err != nil {
		e := errors.NewError(500, "计算所有用户数量错误")
		return errors.GetError(e, err.Error())
	}
	if pageSize < 1 {
		pageSize = int(sum)
	}

	//分页查询
	pageNum := int(sum) / pageSize
	total := pageNum * pageSize

	for i := 0; i < total; i += pageSize {
		if err := handleMultipleAccountData(pageSize, i, handler); err != nil {
			return err
		}
	}

	//不整除情况
	if int(sum)%pageSize != 0 {
		if err := handleMultipleAccountData(total, int(sum)-total, handler); err != nil {
			return err
		}
	}
	return nil
}

// HandleOneAccount 处理一个用户
//
//	@param	id		用户id
//	@param	handler	回调函数
func HandleOneAccount(id string, handler func(user *entity.Account) *errors.MyError) *errors.MyError {
	account := entity.Account{}
	if err := SelectWebAccountService(id, &account); err != nil {
		logger.Logger.Err(err, 0)
		return err
	}
	if err := handler(&account); err != nil {
		logger.Logger.Err(err, 0)
		return err
	}
	return nil
}
