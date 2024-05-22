package config

import (
	"fmt"
	"log"
	"os"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Server struct {
		Port int `toml:"port"`
	} `toml:"server"`
	Database struct {
		User     string
		Password string
		Name     string
	}
	MysqlDB struct {
		Dsn string `toml:"dsn"`
	}
	Redis struct {
		RedisAddr string `toml:"redis_addr"`
		RedisPass string `toml:"redis_pass"`
	}
}

var Conf Config

func init() {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("无法获取当前工作目录: %v", err)
	}
	fmt.Printf("当前工作目录: %s\n", cwd)

	// 读取配置文件
	if _, err := toml.DecodeFile("config/config.toml", &Conf); err != nil {
		log.Fatalf("无法读取配置文件: %v", err)
	}
}
