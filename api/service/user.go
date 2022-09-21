package service

import (
	"errors"
	"github.com/aisuosuo/letter/api/jwt"
	"github.com/aisuosuo/letter/api/models"
	"github.com/aisuosuo/letter/config/db"
)

var (
	UserService = new(userService)
	UserModel   = models.UserMgr(db.GetDB())
)

type userService struct{}

func (t *userService) Register(user *models.User) error {
	var userCount int64
	UserModel.Where("name", user.Name).Count(&userCount)
	if userCount > 0 {
		return errors.New("user already exists")
	}

	user.Password = PasswordEncrypt(user.Password)
	return UserModel.Create(&user).Error
}

func (t *userService) Login(user *models.User) (string, error) {
	users, err := UserModel.GetByOptions(
		UserModel.WithName(user.Name),
		UserModel.WithPassword(PasswordEncrypt(user.Password)),
	)
	if err != nil {
		return "", err
	}
	if len(users) == 0 {
		return "", errors.New("user does not exist")
	}

	claims := map[string]interface{}{
		"uid": users[0].UID,
	}
	token, err := jwt.Sign(claims)
	if err != nil {
		return "", err
	}
	return token, nil
}
