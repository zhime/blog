package initialize

import (
	"github.com/zhime/blog/server/global"
	"github.com/zhime/blog/server/model"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"os"
)

func RegisterTables(db *gorm.DB) {

	_ = db.SetupJoinTable(&model.UserModel{}, "CollectModel", &model.UserCollectModel{})
	_ = db.SetupJoinTable(&model.MenuModel{}, "Banners", &model.MenuBannerModel{})
	err := db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(
		&model.BannerModel{},
		&model.TagModel{},
		&model.MessageModel{},
		&model.AdvertModel{},
		&model.UserModel{},
		&model.CommentModel{},
		&model.ArticleModel{},
		&model.MenuModel{},
		&model.MenuBannerModel{},
		&model.FadeBackModel{},
		&model.LoginModel{},
	)
	if err != nil {
		global.Logger.Error("register table failed", zap.Error(err))
		os.Exit(0)
	}
	global.Logger.Info("register table success")
}
