package models

import (
	"time"
)

type OutMessages struct {
	ID         uint   `json:"ID"`
	FromUserID uint   `json:"fromUserId"`
	ToUserID   uint   `json:"toUserId"`
	Content    string `json:"content"`
	CreateAt   string `json:"createAt"`
}

// Messages [...]
type Messages struct {
	ID         uint      `gorm:"primaryKey;column:id" json:"-"`
	FromUserID uint      `gorm:"column:from_user_id" json:"fromUserId"`
	ToUserID   uint      `gorm:"column:to_user_id" json:"toUserId"`
	Content    string    `gorm:"column:content" json:"content"`
	CreateAt   time.Time `gorm:"column:create_at" json:"createAt"`
}

// TableName get sql table name.获取数据库表名
func (m *Messages) TableName() string {
	return "messages"
}

// MessagesColumns get sql column name.获取数据库列名
var MessagesColumns = struct {
	ID         string
	FromUserID string
	ToUserID   string
	Content    string
	CreateAt   string
}{
	ID:         "id",
	FromUserID: "from_user_id",
	ToUserID:   "to_user_id",
	Content:    "content",
	CreateAt:   "create_at",
}

// User [...]
type User struct {
	UID      uint   `gorm:"primaryKey;column:uid" json:"uid"`
	Name     string `gorm:"column:name" json:"name"`
	Password string `gorm:"column:password" json:"password"`
}

// TableName get sql table name.获取数据库表名
func (m *User) TableName() string {
	return "user"
}

// UserColumns get sql column name.获取数据库列名
var UserColumns = struct {
	UID      string
	Name     string
	Password string
}{
	UID:      "uid",
	Name:     "name",
	Password: "password",
}

// UserFriend [...]
type UserFriend struct {
	ID       uint `gorm:"primaryKey;column:id" json:"-"`
	UserID   uint `gorm:"column:user_id" json:"userId"`
	FriendID uint `gorm:"column:friend_id" json:"friendId"`
}

// TableName get sql table name.获取数据库表名
func (m *UserFriend) TableName() string {
	return "user_friend"
}

// UserFriendColumns get sql column name.获取数据库列名
var UserFriendColumns = struct {
	ID       string
	UserID   string
	FriendID string
}{
	ID:       "id",
	UserID:   "user_id",
	FriendID: "friend_id",
}
