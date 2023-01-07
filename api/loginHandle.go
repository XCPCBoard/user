package api

import (
	"github.com/XCPCBoard/api/errors"
	response "github.com/XCPCBoard/api/http"
	"github.com/XCPCBoard/api/http/token"
	"github.com/XCPCBoard/user/entity"
	"github.com/XCPCBoard/user/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

//UserLoginHandle 用户登录handle
func UserLoginHandle(ctx *gin.Context) {
	//获取参数
	keyword := ctx.PostForm("keyword")
	email := ctx.PostForm("email")
	user := entity.User{}

	//搜素用户
	if err := service.SelectUserServiceByEmail(email, &user); err != nil {
		ctx.Error(errors.NewError(http.StatusForbidden, "搜索用户失败："+err.Error()))
		return
	}
	//对比密码
	keyword = service.GetPassWord(keyword, &user)
	if keyword != user.Keyword {
		ctx.Error(errors.NewError(http.StatusForbidden, "密码错误"))
		return
	}

	//获取token
	tk, err := token.GenerateToken(user.Account, strconv.Itoa(user.Id))
	if err != nil {
		ctx.Error(errors.NewError(http.StatusInternalServerError, "token生成错误"))
		return
	}

	ctx.JSON(http.StatusOK, response.SuccessResponse("ok", map[string]interface{}{
		"token": tk,
	}))

}
