package global

import (
	"gorm.io/gorm"
	"inventory-srv/config"
)

var (
	DB *gorm.DB
	ServerConfig config.ServerConfig
	NacosConfig config.NacosConfig
)
