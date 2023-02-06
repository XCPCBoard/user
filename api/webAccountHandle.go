package api

import (
	"encoding/json"
	"fmt"
	response "github.com/XCPCBoard/api/http"
	"github.com/XCPCBoard/common/errors"
	"github.com/XCPCBoard/common/logger"
	"github.com/XCPCBoard/user/entity"
	"github.com/XCPCBoard/user/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// UpdateWebAccountBySelfHandle 更新用户自己的爬虫网站账户
func UpdateWebAccountBySelfHandle(ctx *gin.Context) {
	_logTheCtx := fmt.Sprintf("%#v , %#v", ctx.Params, ctx.Keys)
	//获取参数
	data := &entity.Account{}
	if err := ctx.ShouldBind(data); err != nil {
		logger.Logger.Error("获取更新网站账户时，参数错误", err, 0, _logTheCtx)
		ctx.Error(errors.GetError(errors.VALID_ERROR, "参数错误"))
		return
	}

	//检查权限
	//若待更新的id和token id 不等
	if !CheckUserIDIsCorrect(ctx, strconv.Itoa(data.Id)) {
		logger.Logger.Warn("更新用户网站信息时，用户id和待更新的id不等", 0, _logTheCtx)
		ctx.Error(errors.GetError(errors.VALID_ERROR, "参数错误"))
		return
	}

	//更新
	if err := service.UpdateWebAccountService(data); err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, response.SuccessResponse("ok", nil))

}

// UpdateWebAccountByAdminHandle 管理员更新用户爬虫网站账户信息
func UpdateWebAccountByAdminHandle(ctx *gin.Context) {
	_logTheCtx := fmt.Sprintf("%#v , %#v", ctx.Params, ctx.Keys)
	//获取参数
	data := &entity.Account{}
	if err := ctx.ShouldBind(data); err != nil {
		logger.Logger.Error("获取更新网站账户时，参数错误", err, 0, _logTheCtx)
		ctx.Error(errors.GetError(errors.VALID_ERROR, "参数错误"))
		return
	}

	//检查权限
	//非管理员
	if ok, err := CheckUserIsAdmin(ctx); err != nil {
		logger.Logger.Error("查询用户权限错误", err, 0, _logTheCtx)
		ctx.Error(errors.GetError(errors.INNER_ERROR, "查询用户权限错误"))
		return
	} else if !ok {
		logger.Logger.Warn("用户非管理员", 0, _logTheCtx)
		ctx.Error(errors.GetError(errors.VALID_ERROR, "用户权限错误"))
		return
	}

	//非超管
	//管理员互相之间可以修改这些基本信息

	//更新
	if err := service.UpdateWebAccountService(data); err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, response.SuccessResponse("ok", nil))

}

// SelectWebAccountHandle 查询用户网站账户
func SelectWebAccountHandle(ctx *gin.Context) {
	//获取id
	id := ctx.PostForm("id")

	//谁都能差

	//更新
	user := entity.Account{}
	err := service.SelectWebAccountService(id, &user)
	if err != nil {
		ctx.Error(err)
		return
	} else {
		m := make(map[string]interface{})
		j, _ := json.Marshal(user)
		if err := json.Unmarshal(j, &m); err != nil {
			logger.Logger.Error("json解码错误", err, 0, user)
			ctx.Error(errors.CreateError(500, "json解码错误", err))
		}
		ctx.JSON(http.StatusOK, response.SuccessResponse("ok", m))
	}
}
