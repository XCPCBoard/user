package dao

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"user/config"
)

var DBClient *gorm.DB

const mysqlDriver = "mysql"

//NewDBClient 初始化db连接
func NewDBClient() (*gorm.DB, error) {
	// 判断是否存在配置
	mysqlConfig := config.Conf.Storages[mysqlDriver]

	// 初始化连接
	dsn := fmt.Sprintf("%v:%v@tcp(%v)/%v?parseTime=true",
		mysqlConfig.Username, mysqlConfig.Password, mysqlConfig.Host, mysqlConfig.Base)

	//dbClient, err := sql.Open(mysqlDriver, dsn)
	dBClient, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Errorf("Open Sql Error: %v", err)
		return nil, err
	}

	return dBClient, nil
}
