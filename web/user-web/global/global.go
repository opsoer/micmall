package global

import (
	ut "github.com/go-playground/universal-translator"
	"user-web/config"
	"user-web/proto"
)

var (
	Trans ut.Translator

	ServerConfig = &config.ServerConfig{}

	NacosConfig  = &config.NacosConfig{}

	UserSrvClient proto.UserClient
)
