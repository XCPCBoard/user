package main

import (
	_ "github.com/FengZhg/go_tools/gin_logrus"
	"github.com/XCPCBoard/api/http/middleware"
	"github.com/XCPCBoard/common/config"
	"github.com/XCPCBoard/common/dao"
	_ "github.com/XCPCBoard/common/dao"
	"github.com/XCPCBoard/common/logger"
	"github.com/XCPCBoard/common/mail"
	"github.com/XCPCBoard/user/api"
	"github.com/gin-gonic/gin"
)

// 主入口函数
func main() {

	initEngine()

}

func initEngine() {
	logger.L.Partition("[XCPCBoard BEGIN]")
	logger.L.Info("[XCPCBoard BEGIN]", 0, "")
	logger.L.Partition("[XCPCBoard BEGIN]")

	engine := gin.Default()

	engine.Use(middleware.LoggerToFile())
	engine.Use(middleware.ErrorHandler())

	api.BuildUserRouteEngine(engine)
	engine.Run()

}

func init() {
	//依赖包

	config.BuildConfig("./config.yaml")

	//dao
	redisClient, err := dao.NewRedisClient()
	if err != nil {
		panic(err)
	}
	dao.RedisClient = redisClient
	dbClient, err := dao.NewDBClient()
	if err != nil {
		panic(err)
	}
	dao.DBClient = dbClient

	//log
	if err = logger.InitLogger(); err != nil {
		panic(err)
	}
	//email
	mail.InitEmail()

}
