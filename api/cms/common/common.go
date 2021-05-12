package common

import (
	"gorm.io/gorm"
	"my-apps/api/cms/pkg"
)

// 全局变量

var (
	DBEngine *gorm.DB
	ServerConfig *pkg.ServerConfig
	DatabaseConfig *pkg.DatabaseConfig
)
