package config

import (
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
	"os"
)

func init() {
	// 读取
	f, err := os.Open(configPath)
	defer f.Close()
	if err != nil {
		log.Errorf("Init config Error %v", err)
		return
	}
	// 解构
	err = yaml.NewDecoder(f).Decode(&Conf)
	if err != nil {
		log.Errorf("Decode Conf Error %v", err)
		panic(err)
	}
}

var (
	Conf       = Config{}
	configPath = "./config/config.yaml"
)

type Storage struct {
	Host     string `yaml:"host"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type Config struct {
	Storages map[string]Storage `yaml:"storages"`
}
