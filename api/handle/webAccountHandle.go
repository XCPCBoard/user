package handle

import (
	"fmt"
	response "github.com/XCPCBoard/api/http"
	"github.com/XCPCBoard/common/errors"
	"github.com/XCPCBoard/common/logger"
	"github.com/XCPCBoard/user/entity"
	"github.com/XCPCBoard/user/service"
	"github.com/XCPCBoard/user/service/rank"
	"github.com/gin-gonic/gin"
	"strconv"
)

// UpdateWebAccountBySelfHandle 更新用户自己的爬虫网站账户
func UpdateWebAccountBySelfHandle(ctx *gin.Context) {
	_logTheCtx := fmt.Sprintf("%#v , %#v", ctx.Params, ctx.Keys)
	//获取参数
	data := &entity.Accounts{}
	if err := ctx.ShouldBind(data); err != nil {
		logger.L.Error("获取更新网站账户时，参数错误", err, 0, _logTheCtx)
		ctx.Error(errors.GetError(errors.VALID_ERROR, "参数错误"))
		return
	}

	//检查权限
	//若待更新的id和token id 不等
	errToken := CheckUserIDIsCorrect(ctx, strconv.Itoa(data.Id))
	if errToken != nil {
		return
	}
	//更新
	if err := service.UpdateWebAccountWithOutNumService(data); err != nil {
		ctx.Error(err)
		return
	}

	response.SuccessResponseAddToken(ctx, nil)

}

// SelectWebAccountHandle 查询用户网站账户
func SelectWebAccountHandle(ctx *gin.Context) {
	//获取id
	id := ctx.Query("id")

	//谁都能查

	//更新
	user := entity.Accounts{}
	err := service.SelectWebAccountService(id, &user)
	if err != nil {
		ctx.Error(err)
		return
	}
	m := make(map[string]interface{})
	if err := StrToMap(ctx, user, &m); err != nil {
		logger.L.Err(err, 0)
		return
	}

	response.SuccessResponseAddToken(ctx, m)
}

// SelectAllUserAccountDataPaging 分页获取所有用户的网站所有刷题信息
func SelectAllUserAccountDataPaging(ctx *gin.Context) {
	pageSizeStr := ctx.Query("pageSize")
	pageNumStr := ctx.Query("pageNum")
	pageNum, err := strconv.Atoi(pageNumStr)
	if err != nil {
		e := errors.CreateError(errors.ERROR.Code, "page信息不为整数", pageNumStr)
		logger.L.Err(e, 0)
		ctx.Error(e)
		return
	}
	pageSize, err2 := strconv.Atoi(pageSizeStr)
	if err2 != nil {
		e := errors.CreateError(errors.ERROR.Code, "page信息不为整数", pageSizeStr)
		logger.L.Err(e, 0)
		ctx.Error(e)
		return
	}

	data, err3 := rank.SelectAccountPagingOrderByTotal(pageNum, pageSize)
	if err3 != nil {
		logger.L.Err(err3, 0)
		ctx.Error(err3)
		return
	}
	response.SuccessResponseAddToken(ctx, data)

}
