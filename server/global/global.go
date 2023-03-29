package global

import (
	"github.com/zhime/blog/server/config"
	"gorm.io/gorm"
)

var (
	Config *config.Config
	DB     *gorm.DB
)
