package service

import (
	"errors"
	"github.com/aisuosuo/letter/api/jwt"
	"github.com/aisuosuo/letter/api/models"
	"github.com/aisuosuo/letter/config/db"
)

var (
	UserService = new(userService)
)

type userService struct{}

func (t *userService) Register(user *models.User) error {
	userMgr := models.UserMgr(db.GetDB())
	var userCount int64
	userMgr.Where(models.UserColumns.Name, user.Name).Count(&userCount)
	if userCount > 0 {
		return errors.New("user already exists")
	}

	user.Password = PasswordEncrypt(user.Password)
	return models.UserFriendMgr(db.GetDB()).Create(&user).Error
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

func (t *userService) GetUser(name string) (*models.User, error) {
	userMgr := models.UserMgr(db.GetDB())
	user, err := userMgr.GetByOption(
		userMgr.WithName(name),
	)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (t *userService) AddFriend(userId, friendId uint) error {
	userFriendMgr := models.UserFriendMgr(db.GetDB())
	friendShip := models.UserFriend{
		UserID:   userId,
		FriendID: friendId,
	}
	return userFriendMgr.Create(&friendShip).Error
}

func (t *userService) DeleteFriend(userId, friendId uint) error {
	userFriendMgr := models.UserFriendMgr(db.GetDB())
	return userFriendMgr.Where(models.UserFriendColumns.UserID, userId).Where(models.UserFriendColumns.FriendID, friendId).Delete(&models.UserFriend{}).Error
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
	}
	return messagesMgr.Create(&message).Error
}
