package models

import (
	"time"
)

// Messages [...]
type Messages struct {
	ID         uint      `gorm:"autoIncrement:true;primaryKey;column:id;type:int unsigned;not null" json:"id"`
	FromUserID uint      `gorm:"index:messages_from_user_id_index;column:from_user_id;type:int unsigned;not null;default:0" json:"fromUserId"`
	ToUserID   uint      `gorm:"index:messages_to_user_id_index;column:to_user_id;type:int unsigned;not null;default:0" json:"toUserId"`
	Content    string    `gorm:"column:content;type:text;default:null" json:"content"`
	CreateAt   time.Time `gorm:"index:messages_create_at_index;column:create_at;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"createAt"`
}

// TableName get sql table name.获取数据库表名
func (m *Messages) TableName() string {
	return "messages"
}

// User [...]
type User struct {
	UID      uint      `gorm:"autoIncrement:true;primaryKey;column:uid;type:int unsigned;not null" json:"uid"`
	Name     string    `gorm:"unique;column:name;type:varchar(20);not null;default:'';comment:用户名" json:"name"` // 用户名
	Password string    `gorm:"column:password;type:varchar(32);not null;default:''" json:"password"`
	Avatar   string    `gorm:"column:avatar;type:varchar(256);not null;default:'';comment:头像" json:"avatar"` // 头像
	CreateAt time.Time `gorm:"column:create_at;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"createAt"`
	UpdateAt time.Time `gorm:"column:update_at;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"updateAt"`
}

// TableName get sql table name.获取数据库表名
func (m *User) TableName() string {
	return "user"
}

// UserFriend [...]
type UserFriend struct {
	ID       uint `gorm:"autoIncrement:true;primaryKey;column:id;type:int unsigned;not null" json:"id"`
	UserID   uint `gorm:"uniqueIndex:uk_user_friend;column:user_id;type:int unsigned;not null;default:0" json:"userId"`
	FriendID uint `gorm:"uniqueIndex:uk_user_friend;column:friend_id;type:int unsigned;not null;default:0" json:"friendId"`
}

// TableName get sql table name.获取数据库表名
func (m *UserFriend) TableName() string {
	return "user_friend"
}
