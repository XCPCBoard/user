package server

import (
	"github.com/FengZhg/go_tools/gin_middleware"
	"github.com/gin-gonic/gin"
	"net/http"
	"os/exec"
)

func init() {
	exec.Command("")
}

func NewServer() *gin.Engine {
	// 生成默认
	engine := gin.Default()

	// 中间件
	// 超时控制中间件
	engine.Use(gin_middleware.RequestLogMiddleware())
	engine.Use(gin_middleware.ReplyMiddleware())
	engine.Use(gin_middleware.TimeoutMiddleware())
	//engine.Use(util.AuthMiddleware())

	// 不校验登录态的接口 只有ldap登陆的时候没有中间件，在model文件中进行区分


	// 校验登录态的接口


	engine.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "Server Started")
	})

	return engine
}
