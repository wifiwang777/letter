package models

// User [...]
type User struct {
	UID      uint   `gorm:"primaryKey;column:uid" json:"-"`
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
