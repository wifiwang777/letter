package service

import (
	"errors"
	"fmt"
	"github.com/aisuosuo/letter/api/jwt"
	"github.com/aisuosuo/letter/api/models"
	"github.com/aisuosuo/letter/config/db"
	"gorm.io/gorm"
	"time"
)

var (
	UserService = new(userService)
)

type userService struct{}

func (t *userService) Register(user *models.User) error {
	userMgr := models.UserMgr(db.GetDB())
	var userCount int64
	userMgr.Where("name", user.Name).Count(&userCount)
	if userCount > 0 {
		return errors.New("user already exists")
	}

	user.Password = PasswordEncrypt(user.Password)
	return userMgr.Create(&user).Error
}

func (t *userService) Login(user *models.User) (string, error) {
	userMgr := models.UserMgr(db.GetDB())
	users, err := userMgr.GetByOptions(
		userMgr.WithName(user.Name),
		userMgr.WithPassword(PasswordEncrypt(user.Password)),
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

func (t *userService) GetFriends(uid uint) (friendList []*models.User) {
	userFriendMgr := models.UserFriendMgr(db.GetDB())
	userFriendMgr.Select("user.*").Joins("left join user on user_friend.friend_id = user.uid").Where("user_friend.user_id", uid).Scan(&friendList)
	return
}

func (t *userService) UserInfo(uid uint) (*models.User, error) {
	userMgr := models.UserMgr(db.GetDB())
	user, err := userMgr.GetByOption(
		userMgr.WithUID(uid),
	)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (t *userService) SearchUser(name string) ([]*models.User, error) {
	userMgr := models.UserMgr(db.GetDB())
	var users []*models.User
	userMgr.Where("name LIKE ?", fmt.Sprintf("%%%s%%", name)).Find(&users)
	return users, nil
}

func (t *userService) AcceptFriend(userId, friendId uint) error {
	friendShipFrom := models.UserFriend{
		UserID:   userId,
		FriendID: friendId,
	}
	friendShipTo := models.UserFriend{
		UserID:   friendId,
		FriendID: userId,
	}

	return db.GetDB().Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&friendShipFrom).Error; err != nil {
			return err
		}
		if err := tx.Create(&friendShipTo).Error; err != nil {
			return err
		}
		return nil
	})
}

func (t *userService) DeleteFriend(userId, friendId uint) error {
	userFriendMgr := models.UserFriendMgr(db.GetDB())
	return userFriendMgr.Where("user_id", userId).Where("friend_id", friendId).Delete(&models.UserFriend{}).Error
}

func (t *userService) IsFriend(userId, friendId uint) bool {
	userFriendMgr := models.UserFriendMgr(db.GetDB())
	var count int64
	userFriendMgr.Where("user_id", userId).Where("friend_id", friendId).Count(&count)
	return count > 0
}

func (t *userService) GetMessages(userId, friendId uint) (messages []*models.Messages) {
	messageMgr := models.MessagesMgr(db.GetDB())
	messageMgr.Where("from_user_id in ?", []uint{userId, friendId}).Where("to_user_id in ?", []uint{userId, friendId}).Scan(&messages)
	return
}

func (t *userService) AddMessage(userId, friendId uint, content string) error {
	messagesMgr := models.MessagesMgr(db.GetDB())
	message := models.Messages{
		FromUserID: userId,
		ToUserID:   friendId,
		Content:    content,
		CreateAt:   time.Now(),
	}
	return messagesMgr.Create(&message).Error
}

func (t *userService) UpdateAvatar(uid uint, fileName string) error {
	user, err := t.UserInfo(uid)
	if err != nil {
		return err
	}

	return models.UserMgr(db.GetDB()).Model(user).Update("avatar", fileName).Error
}
