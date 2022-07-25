package global

import (
	"gorm.io/gorm"
	"userop-srv/config"
)

var (
	DB *gorm.DB
	ServerConfig config.ServerConfig
	NacosConfig config.NacosConfig
)
