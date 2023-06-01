package models

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _UserFriendMgr struct {
	*_BaseMgr
}

// UserFriendMgr open func
func UserFriendMgr(db *gorm.DB) *_UserFriendMgr {
	if db == nil {
		panic(fmt.Errorf("UserFriendMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_UserFriendMgr{_BaseMgr: &_BaseMgr{DB: db.Table("user_friend"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_UserFriendMgr) GetTableName() string {
	return "user_friend"
}

// Reset 重置gorm会话
func (obj *_UserFriendMgr) Reset() *_UserFriendMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_UserFriendMgr) Get() (result UserFriend, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UserFriend{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_UserFriendMgr) Gets() (results []*UserFriend, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UserFriend{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_UserFriendMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(UserFriend{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_UserFriendMgr) WithID(id uint) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithUserID user_id获取
func (obj *_UserFriendMgr) WithUserID(userID uint) Option {
	return optionFunc(func(o *options) { o.query["user_id"] = userID })
}

// WithFriendID friend_id获取
func (obj *_UserFriendMgr) WithFriendID(friendID uint) Option {
	return optionFunc(func(o *options) { o.query["friend_id"] = friendID })
}

// GetByOption 功能选项模式获取
func (obj *_UserFriendMgr) GetByOption(opts ...Option) (result UserFriend, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(UserFriend{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_UserFriendMgr) GetByOptions(opts ...Option) (results []*UserFriend, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(UserFriend{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_UserFriendMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]UserFriend, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(UserFriend{}).Where(options.query)
	query.Count(&count)
	resultPage.SetTotal(count)
	if len(page.GetOrederItemsString()) > 0 {
		query = query.Order(page.GetOrederItemsString())
	}
	err = query.Limit(int(page.GetSize())).Offset(int(page.Offset())).Find(&results).Error

	resultPage.SetRecords(results)
	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容
func (obj *_UserFriendMgr) GetFromID(id uint) (result UserFriend, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UserFriend{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_UserFriendMgr) GetBatchFromID(ids []uint) (results []*UserFriend, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UserFriend{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromUserID 通过user_id获取内容
func (obj *_UserFriendMgr) GetFromUserID(userID uint) (results []*UserFriend, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UserFriend{}).Where("`user_id` = ?", userID).Find(&results).Error

	return
}

// GetBatchFromUserID 批量查找
func (obj *_UserFriendMgr) GetBatchFromUserID(userIDs []uint) (results []*UserFriend, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UserFriend{}).Where("`user_id` IN (?)", userIDs).Find(&results).Error

	return
}

// GetFromFriendID 通过friend_id获取内容
func (obj *_UserFriendMgr) GetFromFriendID(friendID uint) (results []*UserFriend, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UserFriend{}).Where("`friend_id` = ?", friendID).Find(&results).Error

	return
}

// GetBatchFromFriendID 批量查找
func (obj *_UserFriendMgr) GetBatchFromFriendID(friendIDs []uint) (results []*UserFriend, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UserFriend{}).Where("`friend_id` IN (?)", friendIDs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_UserFriendMgr) FetchByPrimaryKey(id uint) (result UserFriend, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UserFriend{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchIndexByIndexUserID  获取多个内容
func (obj *_UserFriendMgr) FetchIndexByIndexUserID(userID uint) (results []*UserFriend, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UserFriend{}).Where("`user_id` = ?", userID).Find(&results).Error

	return
}
