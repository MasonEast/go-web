package global

import (
	"myapp/config"

	"gorm.io/gorm"
)

var (
	GB_CONFIG config.Server
	GB_DB *gorm.DB
)