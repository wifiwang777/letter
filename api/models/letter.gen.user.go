package models

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _UserMgr struct {
	*_BaseMgr
}

// UserMgr open func
func UserMgr(db *gorm.DB) *_UserMgr {
	if db == nil {
		panic(fmt.Errorf("UserMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_UserMgr{_BaseMgr: &_BaseMgr{DB: db.Table("user"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_UserMgr) GetTableName() string {
	return "user"
}

// Reset 重置gorm会话
func (obj *_UserMgr) Reset() *_UserMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_UserMgr) Get() (result User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_UserMgr) Gets() (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_UserMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(User{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithUID uid获取
func (obj *_UserMgr) WithUID(uid uint) Option {
	return optionFunc(func(o *options) { o.query["uid"] = uid })
}

// WithName name获取
func (obj *_UserMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithPassword password获取
func (obj *_UserMgr) WithPassword(password string) Option {
	return optionFunc(func(o *options) { o.query["password"] = password })
}

// GetByOption 功能选项模式获取
func (obj *_UserMgr) GetByOption(opts ...Option) (result User, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_UserMgr) GetByOptions(opts ...Option) (results []*User, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromUID 通过uid获取内容
func (obj *_UserMgr) GetFromUID(uid uint) (result User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where("`uid` = ?", uid).First(&result).Error

	return
}

// GetBatchFromUID 批量查找
func (obj *_UserMgr) GetBatchFromUID(uids []uint) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where("`uid` IN (?)", uids).Find(&results).Error

	return
}

// GetFromName 通过name获取内容
func (obj *_UserMgr) GetFromName(name string) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where("`name` = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量查找
func (obj *_UserMgr) GetBatchFromName(names []string) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where("`name` IN (?)", names).Find(&results).Error

	return
}

// GetFromPassword 通过password获取内容
func (obj *_UserMgr) GetFromPassword(password string) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where("`password` = ?", password).Find(&results).Error

	return
}

// GetBatchFromPassword 批量查找
func (obj *_UserMgr) GetBatchFromPassword(passwords []string) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where("`password` IN (?)", passwords).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_UserMgr) FetchByPrimaryKey(uid uint) (result User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where("`uid` = ?", uid).First(&result).Error

	return
}
