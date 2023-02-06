package api

import (
	"errors"
	"fmt"
	"github.com/XCPCBoard/api/http/middleware"
	"github.com/XCPCBoard/common/config"
	"github.com/XCPCBoard/user/service"
	"github.com/gin-gonic/gin"
)

const (
	TokenIDStr   = middleware.TokenIDStr
	TokenNameStr = middleware.TokenAccountStr
)

// CheckUserIDIsCorrect 根据用户id检查是否和token中的id一致,一致则返回true，限制用户只能修改自己的数据
func CheckUserIDIsCorrect(ctx *gin.Context, id string) bool {
	ID, ok := ctx.Get(TokenIDStr)
	tokenID := fmt.Sprintf("%v", ID)
	return ok && tokenID == id
}

// CheckUserIsAdmin 检测用户是否为管理员
// @return bool,error : 是否为管理，错误信息
func CheckUserIsAdmin(ctx *gin.Context) (bool, error) {

	user := map[string]interface{}{}
	tokenId, tokenName, err := GetUserMsgFromCtx(ctx)
	if err != nil {
		return false, err
	}
	if err := service.SelectUserServiceWithOutPassword(tokenId, &user); err != nil {
		return false, err
	}
	if account, ok := user["account"]; ok && fmt.Sprintf("%v", account) == tokenName {
		if isAdmin, ok2 := user["is_administrator"]; ok2 && fmt.Sprintf("%v", isAdmin) == "1" {
			return true, nil
		}
	}
	return false, nil
}

// CheckUserIsSuperAdmin 检测用户是否为超级管理员
func CheckUserIsSuperAdmin(ctx *gin.Context) (bool, error) {
	if tokenName, ok := ctx.Get(TokenNameStr); ok && tokenName == config.Conf.Admin.Name {
		return CheckUserIsAdmin(ctx)
	}
	return false, nil
}

// GetUserMsgFromCtx 获取用户id和name
// @return tokenID,tokenName,error
func GetUserMsgFromCtx(ctx *gin.Context) (string, string, error) {
	ID, ok := ctx.Get(TokenIDStr)
	if !ok {
		return "", "", errors.New("获取context中用户id失败")
	}
	Name, ck := ctx.Get(TokenNameStr)
	if !ck {
		return "", "", errors.New("获取context中用户Name失败")
	}
	tokenID := fmt.Sprintf("%v", ID)
	tokenName := fmt.Sprintf("%v", Name)
	return tokenID, tokenName, nil
}

////log.Warn并将err写入ctx
//func logWarnAndCtxError(ctx *gin.Context, msg string, code int, param string) {
//	logger.Logger.Warn(msg, 1, param)
//	ctxError(ctx, msg, code)
//}
//
//func logErrorAndCtxError(ctx *gin.Context, msg string, code int, myErr error, param string) {
//	logger.Logger.Error(msg, myErr, 1, param)
//	ctxError(ctx, msg, code)
//}
//
////将error写入ctx
//func ctxError(ctx *gin.Context, msg string, code int) {
//	myErr := errors2.MyError{
//		Msg: msg,
//		Code: code,
//		Data:
//	}
//	ctx.Error(&myErr)
//
//	//ctx.Error的返回值一定是非空的
//	//if err := ctx.Error(myErr); err != nil {
//	//	logger.Logger.Error("gin ctx get error", err, 0, fmt.Sprintf("ctx:%+v", ctx))
//	//	ctx.AbortWithStatusJSON(http.StatusNotAcceptable, myErr)
//	//}
//	return
//}
