package config

import (
	"github.com/BurntSushi/toml"
	"log"
)

type Configs struct {
	Title 	string
	Mysql 	MysqlConfig `toml:"mysql"`
	Server 	ServerConfig `toml:"server"`
}

type MysqlConfig struct {
	DataSource string `toml:"data_source"`
}

type ServerConfig struct {
	ListenAddr string `toml:"listen_addr"`
}

//数据中心配置对象
var DCConfig Configs

func init() {
	configFile := "config/config.toml"
	if _, err := toml.DecodeFile(configFile, &DCConfig); err != nil {
		log.Panic(err)
	}
}
