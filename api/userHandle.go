package api

import (
	"github.com/XCPCBoard/api/errors"
	response "github.com/XCPCBoard/api/http"
	"github.com/XCPCBoard/user/entity"
	"github.com/XCPCBoard/user/service"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

//********************************************************************
//					     		 CRUD
//********************************************************************

//CreatUserHandle 新建账户handle函数
func CreatUserHandle(ctx *gin.Context) {

	account := ctx.PostForm("account")
	keyword := ctx.PostForm("keyword")
	email := ctx.PostForm("email")

	if err := service.CreateUserInitService(account, keyword, email); err != nil {
		ctx.Error(errors.GetError(errors.ILLEGAL_DATA, "Creat user error"))
		return
	}

	ctx.JSON(http.StatusOK, response.SuccessResponse("ok", nil))
}

//DeleteUserHandle 删除用户
func DeleteUserHandle(ctx *gin.Context) {

	id := ctx.PostForm("id")
	uid, err := strconv.Atoi(id)
	if err != nil {
		log.Errorf(err.Error())
		ctx.Error(errors.GetError(errors.ILLEGAL_DATA, "id 错误"))
		return
	}

	//用户非管理员时，直接return
	if ok, err := CheckUserIsAdmin(ctx); !ok { //非管理员
		if err != nil {
			ctx.Error(errors.NewError(http.StatusInternalServerError, "验证管理员时错误:"+err.Error()))
			return
		}
		ctx.Error(errors.NewError(http.StatusForbidden, "非管理员"))
		return
	}

	//检测被删除用户是否为管理员
	if ok, err := service.SelectUserIsAdmin(id); ok { //是管理员

		//被删除用户为管理员，则检查登录用户是否为超级管理员
		if ok, err := CheckUserIsSuperAdmin(ctx); !ok { //非超管
			if err != nil {
				ctx.Error(errors.NewError(http.StatusInternalServerError, "查询当前用户时出错"))
				return
			} else {
				ctx.Error(errors.NewError(http.StatusForbidden, "删除用户为管理员，但是登录用户不是超级管理员"))
				return
			}
		}
		//是超管就继续
	} else if err != nil { //不是管理员，但是出错
		ctx.Error(errors.NewError(http.StatusInternalServerError, "查询删除用户时出错"))
		return
	}

	//删除
	user := entity.User{
		Id: uid,
	}
	if err := service.DeleteUserService(&user); err != nil {
		ctx.Error(errors.GetError(errors.INNER_ERROR, "Delete user error"))
		return
	}

	ctx.JSON(http.StatusOK, response.SuccessResponse("ok", nil))

}

//UpdateUserSelfHandle 更新用户数据，仅限本人（不包含基础信息，只有name，imag，signature，phone，qq)，管理员有别的函数来修改
func UpdateUserSelfHandle(ctx *gin.Context) {
	uid := ctx.PostForm("id")

	if !CheckUserIDIsCorrect(ctx, uid) { //查看是否是自己修改自己
		ctx.Error(errors.NewError(http.StatusForbidden, "用户非本人限或token错误"))
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
		ctx.Error(errors.GetError(errors.INNER_ERROR, "Update user error"))
		return
	}
	ctx.JSON(http.StatusOK, response.SuccessResponse("ok", nil))
}

//SelectUserHandle 查询用户（限制只能管理员查，或自己查自己)
func SelectUserHandle(ctx *gin.Context) {

	id := ctx.Query("id")
	if !CheckUserIDIsCorrect(ctx, id) {
		if ok, err := CheckUserIsAdmin(ctx); !ok { //非管理员
			if err != nil {
				ctx.Error(errors.NewError(http.StatusInternalServerError, "验证管理员时错误:"+err.Error()))
				return
			}
			ctx.Error(errors.NewError(http.StatusForbidden, "用户无权限或token错误"))
			return
		} //是管理员时就可以继续查
	}
	user := map[string]interface{}{}
	err := service.SelectUserService(id, &user)
	if err != nil {
		ctx.Error(errors.GetError(errors.INNER_ERROR, "select user error"))
		return
	}
	ctx.JSON(http.StatusOK, response.SuccessResponse("ok", user))

}
