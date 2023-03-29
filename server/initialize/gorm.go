package initialize

import (
	"github.com/zhime/blog/server/global"
	"github.com/zhime/blog/server/model"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"os"
)

func RegisterTables(db *gorm.DB) {
	err := db.AutoMigrate(
		model.User{},
	)
	if err != nil {
		global.Logger.Error("register table failed", zap.Error(err))
		os.Exit(0)
	}
	global.Logger.Info("register table success")
}
