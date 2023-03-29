package controller

import (
	"errors"
	uuid "github.com/satori/go.uuid"
	"github.com/zhime/blog/server/global"
	"github.com/zhime/blog/server/model"
	"github.com/zhime/blog/server/utils"
	"gorm.io/gorm"
)

type UserService struct{}

func (userService *UserService) Register(u model.User) (userInter model.User, err error) {
	var modelUser model.User
	if !errors.Is(global.DB.Where("user_name = ?", u.UserName).First(&modelUser).Error, gorm.ErrRecordNotFound) { // 判断用户名是否注册
		return userInter, errors.New("用户名已注册")
	}
	u.Password = utils.BcryptHash(u.Password)
	u.UUID = uuid.NewV4()
	err = global.DB.Create(&u).Error
	return u, err
}
