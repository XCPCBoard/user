package api

import (
	"github.com/XCPCBoard/user/service"
	response "github.com/XCPCBoard/utils/http"
	"github.com/gin-gonic/gin"
	"net/http"
)

//UpdateWebAccountHandle 更新用户网站账户
func UpdateWebAccountHandle(ctx *gin.Context) {
	account := ctx.PostFormMap("account")

	if err := service.UpdateWebAccountService(account); err != nil {
		ctx.JSON(http.StatusForbidden, response.FailResponse("update error", nil))
	} else {
		ctx.JSON(http.StatusOK, response.SuccessResponse("ok", nil))
	}
}

//SelectWebAccountHandle 查询用户网站账户
func SelectWebAccountHandle(ctx *gin.Context) {
	id := ctx.PostForm("id")

	user := map[string]interface{}{}
	err := service.SelectWebAccountService(id, &user)
	if err != nil {
		ctx.JSON(http.StatusForbidden, response.FailResponse("update error", nil))
	} else {
		ctx.JSON(http.StatusOK, response.SuccessResponse("ok", user))
	}
}
