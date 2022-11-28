package main

import (
	"fmt"
	_ "github.com/FengZhg/go_tools/gin_logrus"
	_ "user/config"
	"user/dao"
	_ "user/dao"
	"user/users/controller"
)

// 主入口函数
func main() {

}

func init() {
	//redisClient, err := dao.NewRedisClient()
	//if err != nil {
	//	panic(err)
	//}
	dbClient, err := dao.NewDBClient()
	if err != nil {
		panic(err)
	}
	//dao.RedisClient = redisClient
	dao.DBClient = dbClient

	a, b := controller.CreatUser("test", "123456", "123")
	//a, b := api.GetAllDataOfOneUser("a")
	fmt.Printf("%v,%v", a, b)

}
