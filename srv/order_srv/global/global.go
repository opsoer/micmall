package global

import (
	"gorm.io/gorm"
	"order-srv/config"
	"order-srv/proto"
)

var (
	DB *gorm.DB
	ServerConfig config.ServerConfig
	NacosConfig config.NacosConfig

	GoodsSrvClient proto.GoodsClient
	InventorySrvClient proto.InventoryClient
)

