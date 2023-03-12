package service

import (
	"context"
	"fmt"
	response "github.com/XCPCBoard/api/http"
	"github.com/XCPCBoard/common/dao"
	"github.com/XCPCBoard/common/errors"
	"github.com/XCPCBoard/common/logger"
	"github.com/XCPCBoard/common/mail"
	"github.com/XCPCBoard/common/restriction"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"math/rand"
	"time"
)

var (
	//验证码存储
	EmailKey = "XCPCBoard_email_verification_code_"

	//检测验证码，错误次数不能超过5次
	EmailCheck = "XCPCBoard_email_code_check_"

	RegisterOption       = "register_option"
	ChangePasswordOption = "change_password_option"
)

// sentMsgToEmail 给用户的email发送验证码
func sentMsgToEmail(ctx *gin.Context, email string, optionSalt string, option string,
	optionFunc func(string, string, string) error) {

	_logStr := fmt.Sprintf("option:%v , email:%v", optionSalt, email)
	backGround := context.Background()

	//判断邮箱是否发送频繁
	//限制用户访问次数 59秒内发1次
	if res, err := restriction.LimitAccess(EmailCheck+email+optionSalt, time.Second*59, 1); err != nil {
		logger.L.Error("查看资源限制情况出错", err, 0, _logStr)
		ctx.Error(errors.GetError(errors.INNER_ERROR, "查看邮件发送限制情况出错"))
		return
	} else if !res {
		logger.L.Warn("用户访问资源：发送邮件，频率过高", 0, _logStr)
		ctx.Error(errors.GetError(errors.ERROR, "用户发送邮件频率过高"))
		return
	}

	//生成六位随机验证码
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	vCode := fmt.Sprintf("%06v", rnd.Int31n(1000000))

	//设置验证码到redis，10分钟存活
	if err := dao.RedisClient.Set(backGround, EmailKey+email+optionSalt, vCode, time.Minute*10).Err(); err != nil {
		logger.L.Error("redis 插入email验证码错误", err, 0, _logStr)
		ctx.Error(errors.GetError(errors.INNER_ERROR, "插入email验证码错误"))
		return
	}

	//发送验证码
	if err := optionFunc(option, email, vCode); err != nil {
		logger.L.Error("发送验证码失败", err, 0, _logStr)
		ctx.Error(errors.GetError(errors.INNER_ERROR, "发送email验证码失败"))
		return
	}
	response.SuccessResponse(ctx, nil)

}

// VerifyEmailCode 验证邮件
func VerifyEmailCode(optionSalt string, email string, code string) *errors.MyError {

	_logStr := fmt.Sprintf("option:%v , email:%v,code:%v", optionSalt, email, code)
	backGround := context.Background()

	//限制用户访问次数
	if res, err := restriction.LimitAccess(EmailCheck+email+optionSalt, time.Second*20, 10); err != nil {
		logger.L.Error("查看资源限制情况出错", err, 0, _logStr)
		return errors.GetError(errors.INNER_ERROR, "查看资源限制情况出错")
	} else if !res {
		logger.L.Warn("用户访问资源频率过高", 0, _logStr)
		return errors.GetError(errors.ERROR, "用户访问资源频率过高")
	}

	//获取验证码，这里进一步验证了邮箱
	res, err := dao.RedisClient.Get(backGround, EmailKey+email+optionSalt).Result()
	if err == redis.Nil {
		logger.L.Warn("redis email发送超时，或不存在验证码", 0, _logStr)
		return errors.GetError(errors.VALID_ERROR, "email发送超时，或不存在验证码,或email错误")
	}
	//异常错误
	if err != nil {
		logger.L.Error("redis 获取email验证码错误", err, 0, _logStr)
		return errors.GetError(errors.INNER_ERROR, "检测email验证码是否合法时错误")
	}
	//验证码不对
	if code != res {
		logger.L.Info("用户验证码输入错误", 0, _logStr)
		return errors.GetError(errors.VALID_ERROR, "用户验证码输入错误")
	}

	//ctx.JSON(http.StatusOK, response.SuccessResponse("ok", nil))
	return nil
}

// SentRegisterEmailHandle 发送注册密码（get)
func SentRegisterEmailHandle(ctx *gin.Context) {
	email := ctx.PostForm("email")
	sentMsgToEmail(ctx, email, RegisterOption, "注册账号", mail.Email.NewVerificationCode)
}

// SentResetPasswordEmailHandle 重置密码(get)
func SentResetPasswordEmailHandle(ctx *gin.Context) {
	email := ctx.PostForm("email")
	sentMsgToEmail(ctx, email, ChangePasswordOption, "重置密码", mail.Email.NewVerificationCode)
}

// SentPictureVerCode 发送图片验证码
func SentPictureVerCode(ctx *gin.Context) {

}
