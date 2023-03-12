package handle

import (
	"encoding/json"
	"fmt"
	"github.com/XCPCBoard/api/http/middleware"
	"github.com/XCPCBoard/common/config"
	"github.com/XCPCBoard/common/errors"
	"github.com/XCPCBoard/common/logger"
	"github.com/XCPCBoard/user/service"
	"github.com/gin-gonic/gin"
)

const (
	TokenIDStr   = middleware.TokenIDStr
	TokenNameStr = middleware.TokenAccountStr
)

// CheckUserIDIsCorrect 根据用户id检查是否和token中的id一致,一致则返回true，限制用户只能修改自己的数据
func CheckUserIDIsCorrect(ctx *gin.Context, id string) *errors.MyError {
	ID, ok := ctx.Get(TokenIDStr)
	tokenID := fmt.Sprintf("%v", ID)

	if !(ok && tokenID == id) {
		err := errors.CreateError(errors.ERROR.Code, "用户Id与token中的id不符", fmt.Sprintf("id:%v,tokenid:%v", id, tokenID))
		logger.L.Err(err, 0)
		ctx.Error(errors.GetError(errors.VALID_ERROR, "用户非本人"))
		return err
	}
	return nil
}

// CheckUserIsAdmin 检测用户是否为管理员
// @return bool,error : 是否为管理，错误信息
func CheckUserIsAdmin(ctx *gin.Context) (bool, error) {

	user := map[string]interface{}{}
	tokenId, tokenName, err := GetUserMsgFromCtx(ctx)
	if err != nil {
		return false, err
	}
	if err := service.SelectUserServiceWithOutPassword(tokenId, user); err != nil {
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
func CheckUserIsSuperAdmin(ctx *gin.Context) *errors.MyError {
	e := errors.CreateError(errors.ERROR.Code, "用户非超管", nil)

	if tokenName, ok := ctx.Get(TokenNameStr); ok && tokenName == config.Conf.Admin.Name {
		if ok2, err := CheckUserIsAdmin(ctx); ok2 && err == nil {
			return nil
		} else if err != nil {
			e.Data = err
			e.Msg = "验证用户失败"
		}
	}
	logger.L.Error(e.Msg, e, 0, ctx.Keys)
	ctx.Error(e)
	return e
}

// GetUserMsgFromCtx 获取用户id和name
// @return tokenID,tokenName,error
func GetUserMsgFromCtx(ctx *gin.Context) (string, string, *errors.MyError) {
	ID, ok := ctx.Get(TokenIDStr)
	if !ok {
		return "", "", errors.CreateError(errors.ERROR.Code, "获取context中用户id失败", ctx.Keys)
	}
	Name, ck := ctx.Get(TokenNameStr)
	if !ck {
		return "", "", errors.CreateError(errors.ERROR.Code, "获取context中用户name失败", ctx.Keys)
	}
	tokenID := fmt.Sprintf("%v", ID)
	tokenName := fmt.Sprintf("%v", Name)
	return tokenID, tokenName, nil
}

// StrToMap 格式转换,自带log和ctx.error
func StrToMap(ctx *gin.Context, data interface{}, m *map[string]interface{}) *errors.MyError {
	j, _ := json.Marshal(data)
	if err := json.Unmarshal(j, &m); err != nil {
		e := errors.CreateError(errors.INNER_ERROR.Code, "json解码错误", err)
		logger.L.Error("json解码错误", err, 0, data)
		ctx.Error(e)
		return e
	}
	return nil
}
