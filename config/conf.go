package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"os"
)

const configFile = "/config/config.toml"

type Server struct {
	IP   string
	Port string
}

// MySQL Represents DSN information
type MySQL struct {
	Username  string
	Password  string
	Host      string
	Port      string
	Database  string
	Charset   string
	ParseTime bool `mapstructure:"parse_time"`
	Loc       string
}

type Config struct {
	DB     MySQL `mapstructure:"mysql"`
	Server `mapstructure:"server"`
}

var Cfg Config

func InitEnv() {
	var err error

	currentDir, err := os.Getwd()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	viper.SetConfigFile(currentDir + configFile) // 配置文件路径指定
	err = viper.ReadInConfig()                   // 找到配置文件并读取
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	if err := viper.Unmarshal(&Cfg); err != nil {
		panic(fmt.Errorf("unable to decode into struct, %v", err))
	}
	// 查看读取的配置文件信息
	log.Printf("Server: %+v\n", Cfg.Server)
	log.Printf("MySQL: %+v\n", Cfg.DB)
	//log.Println(currentDir)
}
