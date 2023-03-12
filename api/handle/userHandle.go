package handle

import (
	response "github.com/XCPCBoard/api/http"
	"github.com/XCPCBoard/user/entity"
	"github.com/XCPCBoard/user/service"
	"github.com/gin-gonic/gin"
)

//********************************************************************
//					     		 CRUD
//********************************************************************

// UpdateUserBySelfHandle 更新用户数据，仅限本人（不包含基础信息，只有name，imag，signature，phone，qq)，管理员有别的函数来修改
func UpdateUserBySelfHandle(ctx *gin.Context) {

	uid := ctx.PostForm("id")

	//自己改自己的
	errToken := CheckUserIDIsCorrect(ctx, uid)
	if errToken != nil {
		return
	}
	user := entity.Users{
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
	response.SuccessResponseAddToken(ctx, nil)
}

// SelectUserHandle 查询用户
func SelectUserHandle(ctx *gin.Context) {

	id := ctx.Query("id")
	//谁都能查
	user := make(map[string]interface{})
	err := service.SelectUserServiceWithOutPassword(id, user)
	if err != nil {
		ctx.Error(err)
		return
	}
	response.SuccessResponseAddToken(ctx, user)
}

// SelectUserByNameHandle 查询用户名
func SelectUserByNameHandle(ctx *gin.Context) {

	name := ctx.Query("name")
	//谁都能查
	user := make(map[string]interface{})
	data := new([]entity.Users)
	err := service.SelectUserServiceByName(name, data)
	if err != nil {
		ctx.Error(err)
		return
	}
	user["data"] = data
	response.SuccessResponseAddToken(ctx, user)
}

// SelectUserByAccountHandle 查询用户账户
func SelectUserByAccountHandle(ctx *gin.Context) {

	account := ctx.Query("account")
	//谁都能查
	user := make(map[string]interface{})
	data := new([]entity.Users)
	err := service.SelectUserServiceByAccount(account, data)
	if err != nil {
		ctx.Error(err)
		return
	}
	user["data"] = data
	response.SuccessResponseAddToken(ctx, user)
}
