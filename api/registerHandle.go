package api

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
	"net/http"
	"strconv"
)

var (
	//验证码存储
	emailKey = "XCPCBoard_email_verification_code_"

	//检测验证码，错误次数不能超过5次
	emailCheck = "XCPCBoard_email_code_check_"

	registerOption       = "register_option"
	changePasswordOption = "change_password_option"
)

// RegisterUser 注册用户
func RegisterUser(ctx *gin.Context) {

	//if err:=changePasswordOption;err!=nil{
	//	ctx.Error(err)
	//	return
	//}

	email := ctx.PostForm("email")
	emailCode := ctx.PostForm("email_code")
	account := ctx.PostForm("account")

	//验证密码格式
	keyword := ctx.PostForm("keyword")
	if len(keyword) < 6 && len(keyword) > 30 {
		ctx.Error(errors.GetError(errors.VALID_ERROR, "密码长度有误"))
		return
	}

	//验证邮箱
	if err := verifyEmailCode(registerOption, email, emailCode); err != nil {
		logger.Logger.Error(err.Msg, err, 0, fmt.Sprintf("email:%v", email))
		ctx.Error(err)
		return
	}

	//创建用户，这里前端要保证提前查询过email和account
	if err := service.CreateUserInitService(account, keyword, email); err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, response.SuccessResponse("ok", nil))
}

// UserLoginHandle 用户登录handle
func UserLoginHandle(ctx *gin.Context) {
	//获取参数

	keyword := ctx.PostForm("keyword")
	email := ctx.PostForm("email")
	user := entity.User{}

	//搜素用户
	if ok, err := service.UserEmailIsExist(email); err != nil {
		ctx.Error(errors.GetError(errors.INNER_ERROR, "查询用户是否存在时出错"))
		return
	} else if !ok {
		ctx.Error(errors.GetError(errors.LOGIN_UNKNOWN, "用户不存在:"+email))
		return
	}

	//找到用户数据
	if err := service.SelectUserServiceByEmail(email, &user); err != nil {
		//SelectUserServiceByEmail里写过log了
		ctx.Error(err)
		return
	}
	//对比密码
	keyword = service.GetPassWord(keyword, &user)
	if keyword != user.Keyword {
		//info 级别
		logger.Logger.Info("用户密码错误", 0,
			fmt.Sprintf("email:%v,keyword（encodered):%v", email, keyword))
		ctx.Error(errors.GetError(errors.VALID_ERROR, "密码错误"))
		return
	}

	//获取token
	tk, err := token.GenerateToken(user.Account, strconv.Itoa(user.Id))
	if err != nil {
		logger.Logger.Error("生成token错误", err, 0, fmt.Sprintf("user:%+v", user))
		ctx.Error(errors.GetError(errors.INNER_ERROR, "生成token错误"))
		return
	}

	ctx.JSON(http.StatusOK, response.SuccessResponse("ok", map[string]interface{}{
		"token": tk,
	}))

}

// ResetPassword 重置密码
func ResetPassword(ctx *gin.Context) {

	//if err:=changePasswordOption;err!=nil{
	//	ctx.Error(err)
	//	return
	//}

	//验证邮箱
	email := ctx.PostForm("email")
	emailCode := ctx.PostForm("email_code")
	keyword := ctx.PostForm("keyword")

	//验证密码格式
	if len(keyword) < 6 && len(keyword) > 30 {
		ctx.Error(errors.GetError(errors.VALID_ERROR, "密码长度有误"))
		return
	}

	//验证邮箱验证码
	if err := verifyEmailCode(registerOption, email, emailCode); err != nil {
		logger.Logger.Error(err.Msg, err, 0, fmt.Sprintf("email:%v", email))
		ctx.Error(err)
		return
	}
	//搜索用户
	user := entity.User{}
	if err := service.SelectUserServiceByEmail(email, &user); err != nil {
		ctx.Error(err)
		return
	}

	uid := strconv.Itoa(user.Id)
	//权限，只能自己修改自己
	if !CheckUserIDIsCorrect(ctx, uid) { //查看是否是自己修改自己
		logger.Logger.Warn("用户非本人限或token错误", 0, ctx)
		ctx.Error(errors.GetError(errors.VALID_ERROR, "用户非本人限或token错误"))
		return
	}

	//获取当前时间，为密码加盐

	user.Keyword = service.GetPassWord(user.Keyword, &user)
	if err := service.UpdateUserService(uid, &user); err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, response.SuccessResponse("ok", nil))

}

// 验证用户图形验证码
func checkPictureAuth(ctx *gin.Context) *errors.MyError {
	//用户输入的验证码
	authCode := ctx.PostForm("auth_code")
	//用户验证码id
	authId := ctx.PostForm("auth_id")
	//验证图片码
	if !auth.VerifyCaptcha(authId, authCode) {
		logger.Logger.Info("用户图片验证码输入错误", 0,
			fmt.Sprintf("auth_id:%v,auth_code:%v", authId, authCode))
		return errors.GetError(errors.VALID_ERROR, "验证码错误")
	}
	return nil
}
