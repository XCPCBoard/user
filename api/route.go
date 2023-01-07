package api

import (
	"github.com/XCPCBoard/api/http/middleware"
	"github.com/gin-gonic/gin"
)

func BuildUserRouteEngine(engine *gin.Engine) {
	login := engine.Group("/login")
	{
		login.POST("/create", CreatUserHandle)
		login.POST("/userlogin", UserLoginHandle)
	}

	//加上token验证
	e := engine.Group("/api", middleware.AuthMiddleware())
	{
		user := e.Group("/user")
		{
			//user.POST("/create", CreatUserHandle) 创建用户不需要token
			user.POST("/update", UpdateUserSelfHandle)
			user.POST("/delete", DeleteUserHandle)
			user.GET("/select", SelectUserHandle)
		}
		post := e.Group("/post")
		{
			post.POST("/create", CreatePostHandle)
			post.POST("/delete", DeletePostHandle)
			post.POST("/update", UpdatePostHandle)
			post.GET("/select", SelectPostHandle)
		}
		account := e.Group("/webaccount")
		{
			account.POST("/update", UpdateWebAccountHandle)
			account.GET("/select", SelectWebAccountHandle)
		}
	}
}
