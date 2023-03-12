package mid

import (
	"github.com/XCPCBoard/common/logger"
	"github.com/XCPCBoard/user/api/handle"
	"github.com/gin-gonic/gin"
)

// SuperAdminAuthMiddleware 超管认证中间件
func SuperAdminAuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//检测超管权限
		if err := handle.CheckUserIsSuperAdmin(ctx); err != nil {
			logger.L.Error(err.Msg, err, 0, ctx.Keys)
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
