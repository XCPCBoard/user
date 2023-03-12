package handle

import (
	"fmt"
	response "github.com/XCPCBoard/api/http"
	"github.com/XCPCBoard/api/http/token"
	"github.com/XCPCBoard/common/auth"
	"github.com/XCPCBoard/common/errors"
	"github.com/XCPCBoard/common/logger"
	"github.com/XCPCBoard/user/entity"
	"github.com/XCPCBoard/user/service"
	"github.com/gin-gonic/gin"
	"strconv"
)

// RegisterUser 注册用户
func RegisterUser(ctx *gin.Context) {

	email := ctx.PostForm("email")
	emailCode := ctx.PostForm("email_code")
	account := ctx.PostForm("account")
	//验证密码格式
	keyword := ctx.PostForm("keyword")

	if err := service.UserRegister(account, keyword, email, emailCode); err != nil {
		logger.L.Err(err, 0)
		ctx.Error(err)
		return
	}

	response.SuccessResponse(ctx, nil)
}

// UserLoginHandle 用户登录handle
func UserLoginHandle(ctx *gin.Context) {
	//获取参数

	keyword := ctx.PostForm("keyword")
	email := ctx.PostForm("email")
	user := entity.Users{}

	if err := service.UserLoginService(keyword, email, &user); err != nil {
		logger.L.Err(err, 0)
		ctx.Error(err)
		return
	}

	//获取token
	tk, err := token.GenerateToken(user.Account, strconv.Itoa(user.Id))
	if err != nil {
		e := errors.CreateError(errors.INNER_ERROR.Code, "生成token错误:", user)
		logger.L.Err(e, 0)
		ctx.Error(e)
		return

	}

	response.SuccessResponse(ctx, map[string]interface{}{
		"token": tk,
	})

}

// ResetPassword 重置密码
func ResetPassword(ctx *gin.Context) {

	//验证邮箱
	email := ctx.PostForm("email")
	emailCode := ctx.PostForm("email_code")
	keyword := ctx.PostForm("keyword")

	user, err := service.UserResetPasswordCheck(keyword, email, emailCode)
	if err != nil {

	}
	uid := strconv.Itoa(user.Id)
	//权限，只能自己修改自己
	errToken := CheckUserIDIsCorrect(ctx, uid)
	if errToken != nil {
		return
	}
	user.Keyword, err = service.GetPassWord(user.Keyword)
	if err != nil {
		ctx.Error(err)
		return
	}
	if err := service.UpdateUserService(uid, user); err != nil {
		ctx.Error(err)
		return
	}

	response.SuccessResponse(ctx, nil)

}

// 验证用户图形验证码
func checkPictureAuth(ctx *gin.Context) *errors.MyError {
	//用户输入的验证码
	authCode := ctx.PostForm("auth_code")
	//用户验证码id
	authId := ctx.PostForm("auth_id")
	//验证图片码
	if !auth.VerifyCaptcha(authId, authCode) {
		logger.L.Info("用户图片验证码输入错误", 0,
			fmt.Sprintf("auth_id:%v,auth_code:%v", authId, authCode))
		return errors.GetError(errors.VALID_ERROR, "验证码错误")
	}
	return nil
}
