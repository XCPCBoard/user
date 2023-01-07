package api

import (
	"errors"
	"fmt"
	"github.com/XCPCBoard/common/config"
	"github.com/XCPCBoard/user/service"
	"github.com/gin-gonic/gin"
)

const (
	TokenIDStr   = "xcpc_user_id"
	TokenNameStr = "xcpc_user_name"
)

//CheckUserIDIsCorrect 根据用户id检查是否和token中的id一致,一致则返回true，限制用户只能修改自己的数据
func CheckUserIDIsCorrect(ctx *gin.Context, id string) bool {
	ID, ok := ctx.Get(TokenIDStr)
	tokenID := fmt.Sprintf("%v", ID)
	return ok && tokenID == id
}

//CheckUserIsAdmin 检测用户是否为管理员
//@return bool,error : 是否为管理，错误信息
func CheckUserIsAdmin(ctx *gin.Context) (bool, error) {

	user := map[string]interface{}{}
	tokenId, tokenName, err := GetUserMsgFromCtx(ctx)
	if err != nil {
		return false, err
	}
	if err := service.SelectUserService(tokenId, &user); err != nil {
		return false, err
	}
	if account, ok := user["account"]; ok && fmt.Sprintf("%v", account) == tokenName {
		if isAdmin, ok2 := user["is_administrator"]; ok2 && fmt.Sprintf("%v", isAdmin) == "1" {
			return true, nil
		}
	}
	return false, nil
}

//CheckUserIsSuperAdmin 检测用户是否为超级管理员
func CheckUserIsSuperAdmin(ctx *gin.Context) (bool, error) {
	if tokenName, ok := ctx.Get(TokenNameStr); ok && tokenName == config.Conf.Admin.Name {
		return CheckUserIsAdmin(ctx)
	}
	return false, nil
}

//GetUserMsgFromCtx 获取用户id和name
//@return tokenID,tokenName,error
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
