package api

import (
	"github.com/XCPCBoard/user/entity"
	"github.com/XCPCBoard/user/service"
	response "github.com/XCPCBoard/utils/http"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

//CreatUserHandle 新建账户handle函数
func CreatUserHandle(ctx *gin.Context) {

	account := ctx.PostForm("account")
	keyword := ctx.PostForm("keyword")
	email := ctx.PostForm("email")

	if err := service.CreateUserInitService(account, keyword, email); err != nil {
		ctx.JSON(http.StatusForbidden, response.FailResponse("create user error", nil))
	} else {
		ctx.JSON(http.StatusOK, response.SuccessResponse("ok", nil))
	}
}

//DeleteUserHandle 删除用户
func DeleteUserHandle(ctx *gin.Context) {
	account := ctx.PostForm("account")
	id := ctx.PostForm("id")
	uid, err := strconv.Atoi(id)
	if err != nil {
		log.Errorf(err.Error())
		ctx.JSON(http.StatusForbidden, response.FailResponse("user id error", nil))
	}

	user := entity.User{
		Id:      uid,
		Account: account,
	}
	if err := service.DeleteUserService(&user); err != nil {
		ctx.JSON(http.StatusForbidden, response.FailResponse("delete user error", nil))
	} else {
		ctx.JSON(http.StatusOK, response.SuccessResponse("ok", nil))
	}

}

//UpdateUserHandle 更新用户
func UpdateUserHandle(ctx *gin.Context) {
	user := ctx.PostFormMap("user")
	if err := service.UpdateUserService(user); err != nil {
		ctx.JSON(http.StatusForbidden, response.FailResponse("Update user error", nil))
	} else {
		ctx.JSON(http.StatusOK, response.SuccessResponse("ok", nil))
	}
}

//SelectUserHandle 查询用户
func SelectUserHandle(ctx *gin.Context) {

	id := ctx.PostForm("id")
	user := map[string]interface{}{}
	err := service.SelectUserService(id, &user)
	if err != nil {
		ctx.JSON(http.StatusForbidden, response.FailResponse("Select user error", nil))
	} else {
		ctx.JSON(http.StatusOK, response.SuccessResponse("ok", user))
	}
}
