package rank

import (
	"fmt"
	"github.com/XCPCBoard/common/errors"
	"github.com/XCPCBoard/common/logger"
	"github.com/XCPCBoard/user/entity"
	"github.com/XCPCBoard/user/service"
	"sync"
)

//**********************************************************
//			    			manage						  //
//**********************************************************

// manageMultipleAccountData 处理用户数据
func manageMultipleAccountData(pageSize int, now int, mp map[string]interface{},
	handler func(user *entity.Accounts, mp map[string]interface{}) *errors.MyError) *errors.MyError {
	var err *errors.MyError
	var wg sync.WaitGroup

	//获取数据
	accounts, e := service.SelectMultipleAccount(now, pageSize)
	if e != nil {
		resErr := errors.NewError(errors.INNER_ERROR.Code, "分页查询account错误")
		resErr.Data = e.Error()
		return resErr
	}

	//并发处理
	for _, account := range *accounts {
		wg.Add(1)

		go func(ac *entity.Accounts) {
			//捕获可能在调用逻辑中发生的panic & wg.Down
			defer func() {
				if e := recover(); e != nil {
					logger.L.Error("handler中发生了panic", nil, 0, nil)
				}
				defer wg.Done()
			}()

			// 取最后一个报错的handler调用逻辑，并最终向外返回	，mp用来传参
			e := handler(ac, mp)
			if e != nil {
				logger.L.Error(e.Msg, e, 0,
					fmt.Sprintf("处理account错误 data:%#v", e.Data))
				err = e
			}
		}(&account)
	}
	wg.Wait()
	return err
}

// ManageAllAccount	分批处理全部用户
//
//	@param	pageSize	每批循环处理的用户数，pageSize=10时每批处理10个用户，pageSize<1时直接处理全部用户
//	@param	handler		回调函数
func ManageAllAccount(pageSize int, mp map[string]interface{}, handler func(user *entity.Accounts, mp map[string]interface{}) *errors.MyError) *errors.MyError {
	//获取用户数量
	var sum int64
	if err := service.CountAllAccount(&sum); err != nil {
		e := errors.NewError(errors.INNER_ERROR.Code, "计算所有用户数量错误")
		return errors.GetError(e, err.Error())
	}
	if pageSize < 1 {
		pageSize = int(sum)
	}

	//分页查询
	pageNum := int(sum) / pageSize
	total := pageNum * pageSize

	for i := 0; i < total; i += pageSize {
		if err := manageMultipleAccountData(pageSize, i, mp, handler); err != nil {
			return err
		}
	}

	//不整除情况
	if int(sum)%pageSize != 0 {
		if err := manageMultipleAccountData(total, int(sum)-total, mp, handler); err != nil {
			return err
		}
	}
	return nil
}

// ManageOneAccount 处理一个用户
//
//	@param	id		用户id
//	@param	handler	回调函数
func ManageOneAccount(id string, mp map[string]interface{},
	handler func(user *entity.Accounts, mp map[string]interface{}) *errors.MyError) *errors.MyError {
	account := entity.Accounts{}
	if err := service.SelectWebAccountService(id, &account); err != nil {
		logger.L.Err(err, 0)
		return err
	}
	if err := handler(&account, mp); err != nil {
		logger.L.Err(err, 0)
		return err
	}
	return nil
}
