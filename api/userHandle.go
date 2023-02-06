package api

import (
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

//********************************************************************
//					     		 CRUD
//********************************************************************

////CreatUserHandle 新建账户handle函数
//func CreatUserHandle(ctx *gin.Context) {
//
//	account := ctx.PostForm("account")
//	keyword := ctx.PostForm("keyword")
//	email := ctx.PostForm("email")
//
//	if err := service.CreateUserInitService(account, keyword, email); err != nil {
//		ctx.Error(err)
//		return
//	}
//
//	ctx.JSON(http.StatusOK, response.SuccessResponse("ok", nil))
//}

// DeleteUserHandle 删除用户
func DeleteUserHandle(ctx *gin.Context) {

	id := ctx.PostForm("id")
	_logTheCtx := fmt.Sprintf("%#v , %#v", ctx.Params, ctx.Keys)

	uid, err := strconv.Atoi(id)
	if err != nil {
		logger.Logger.Warn("用户id不为int"+err.Error(), 0, id)
		ctx.Error(errors.GetError(errors.VALID_ERROR, "id 错误"))
		return
	}

	//用户非管理员时，直接return
	if ok, err := CheckUserIsAdmin(ctx); !ok { //非管理员
		if err != nil {
			logger.Logger.Error("验证管理员时错误", err, 0, _logTheCtx)
			ctx.Error(errors.GetError(errors.INNER_ERROR, "验证管理员时错误"))
			return
		}
		logger.Logger.Warn("用户非管理员，无法删除用户", 0, _logTheCtx)
		ctx.Error(errors.GetError(errors.VALID_ERROR, "用户非管理员，无法删除用户"))

		return
	}

	//检测被删除用户是否为管理员
	if ok, err := service.SelectUserIsAdmin(id); ok { //是管理员

		//被删除用户为管理员，则检查登录用户是否为超级管理员
		if ok, err2 := CheckUserIsSuperAdmin(ctx); !ok { //非超管
			if err2 != nil {
				logger.Logger.Error("查询用户是否为超管时，错误", err2, 0, ctx)
				ctx.Error(errors.GetError(errors.INNER_ERROR, "查询用户是否为超管时，错误"))
				return
			} else {
				logger.Logger.Warn("用户非超管，无法删除管理员", 0, ctx)
				ctx.Error(errors.GetError(errors.VALID_ERROR, "用户非超管，无法删除管理员"))
				return
			}
		}
		//是超管就继续
	} else if err != nil { //不是管理员，但是出错
		logger.Logger.Error("查询用户被删是否为管理员时，错误", err, 0, ctx)
		ctx.Error(errors.GetError(errors.INNER_ERROR, "查询用户被删是否为管理员时"))
		return
	}

	//删除
	user := entity.User{
		Id: uid,
	}
	if err := service.DeleteUserService(&user); err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, response.SuccessResponse("ok", nil))

}

// UpdateUserBySelfHandle 更新用户数据，仅限本人（不包含基础信息，只有name，imag，signature，phone，qq)，管理员有别的函数来修改
func UpdateUserBySelfHandle(ctx *gin.Context) {

	uid := ctx.PostForm("id")

	if !CheckUserIDIsCorrect(ctx, uid) { //查看是否是自己修改自己
		logger.Logger.Warn("用户非本人限或token错误", 0, ctx)
		ctx.Error(errors.GetError(errors.VALID_ERROR, "用户非本人限或token错误"))
		return
	}

	user := entity.User{
		Name:        ctx.PostForm("name"),
		ImagePath:   ctx.PostForm("image"),
		Signature:   ctx.PostForm("signature"),
		PhoneNumber: ctx.PostForm("phone"),
		QQNumber:    ctx.PostForm("qq"),
	}

	if err := service.UpdateUserService(uid, &user); err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, response.SuccessResponse("ok", nil))
}

// UpdateUserByAdminHandle 管理员更新用户数据
func UpdateUserByAdminHandle(ctx *gin.Context) {
	uid := ctx.PostForm("id")
	_logTheCtx := fmt.Sprintf("%#v , %#v", ctx.Params, ctx.Keys)

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

	//查询是否是超管
	ok, err := CheckUserIsSuperAdmin(ctx)
	if err != nil {
		logger.Logger.Error("查询用户是否为超管时，错误", err, 0, ctx)
		ctx.Error(errors.GetError(errors.INNER_ERROR, "查询用户是否为超管时，错误"))
		return
	}
	if !ok { //非超管，则查看被修改的是不是管理员
		if ok2, err := service.SelectUserIsAdmin(uid); err != nil {
			logger.Logger.Error("查询被修改用户是否为管理员时，错误", err, 0, ctx)
			ctx.Error(errors.GetError(errors.INNER_ERROR, "查询被修改用户是否为管理员时，错误"))
			return
		} else if ok2 { //修改的是管理员，则报错
			logger.Logger.Warn("管理员不能修改管理员", 0, ctx)
			ctx.Error(errors.GetError(errors.VALID_ERROR, "管理员不能修改管理员"))
			return
		}

	}

	user := entity.User{
		Name:        ctx.PostForm("name"),
		ImagePath:   ctx.PostForm("image"),
		Signature:   ctx.PostForm("signature"),
		PhoneNumber: ctx.PostForm("phone"),
		QQNumber:    ctx.PostForm("qq"),
	}

	if err := service.UpdateUserService(uid, &user); err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, response.SuccessResponse("ok", nil))
}

// SelectUserHandle 查询用户
func SelectUserHandle(ctx *gin.Context) {

	id := ctx.Query("id")
	//谁都能查
	user := make(map[string]interface{})
	err := service.SelectUserServiceWithOutPassword(id, &user)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, response.SuccessResponse("ok", user))
}

func UserAccountIsExist() {

}
