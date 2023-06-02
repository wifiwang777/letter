package models

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _MessagesMgr struct {
	*_BaseMgr
}

// MessagesMgr open func
func MessagesMgr(db *gorm.DB) *_MessagesMgr {
	if db == nil {
		panic(fmt.Errorf("MessagesMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_MessagesMgr{_BaseMgr: &_BaseMgr{DB: db.Table("messages"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// Debug open debug.打开debug模式查看sql语句
func (obj *_MessagesMgr) Debug() *_MessagesMgr {
	obj._BaseMgr.DB = obj._BaseMgr.DB.Debug()
	return obj
}

// GetTableName get sql table name.获取数据库名字
func (obj *_MessagesMgr) GetTableName() string {
	return "messages"
}

// Reset 重置gorm会话
func (obj *_MessagesMgr) Reset() *_MessagesMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_MessagesMgr) Get() (result Messages, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Messages{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_MessagesMgr) Gets() (results []*Messages, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Messages{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_MessagesMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(Messages{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_MessagesMgr) WithID(id uint) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithFromUserID from_user_id获取
func (obj *_MessagesMgr) WithFromUserID(fromUserID uint) Option {
	return optionFunc(func(o *options) { o.query["from_user_id"] = fromUserID })
}

// WithToUserID to_user_id获取
func (obj *_MessagesMgr) WithToUserID(toUserID uint) Option {
	return optionFunc(func(o *options) { o.query["to_user_id"] = toUserID })
}

// WithContent content获取
func (obj *_MessagesMgr) WithContent(content string) Option {
	return optionFunc(func(o *options) { o.query["content"] = content })
}

// WithCreateAt create_at获取
func (obj *_MessagesMgr) WithCreateAt(createAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_at"] = createAt })
}

// GetByOption 功能选项模式获取
func (obj *_MessagesMgr) GetByOption(opts ...Option) (result Messages, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Messages{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_MessagesMgr) GetByOptions(opts ...Option) (results []*Messages, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Messages{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_MessagesMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]Messages, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(Messages{}).Where(options.query)
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
func (obj *_MessagesMgr) GetFromID(id uint) (result Messages, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Messages{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_MessagesMgr) GetBatchFromID(ids []uint) (results []*Messages, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Messages{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromFromUserID 通过from_user_id获取内容
func (obj *_MessagesMgr) GetFromFromUserID(fromUserID uint) (results []*Messages, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Messages{}).Where("`from_user_id` = ?", fromUserID).Find(&results).Error

	return
}

// GetBatchFromFromUserID 批量查找
func (obj *_MessagesMgr) GetBatchFromFromUserID(fromUserIDs []uint) (results []*Messages, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Messages{}).Where("`from_user_id` IN (?)", fromUserIDs).Find(&results).Error

	return
}

// GetFromToUserID 通过to_user_id获取内容
func (obj *_MessagesMgr) GetFromToUserID(toUserID uint) (results []*Messages, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Messages{}).Where("`to_user_id` = ?", toUserID).Find(&results).Error

	return
}

// GetBatchFromToUserID 批量查找
func (obj *_MessagesMgr) GetBatchFromToUserID(toUserIDs []uint) (results []*Messages, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Messages{}).Where("`to_user_id` IN (?)", toUserIDs).Find(&results).Error

	return
}

// GetFromContent 通过content获取内容
func (obj *_MessagesMgr) GetFromContent(content string) (results []*Messages, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Messages{}).Where("`content` = ?", content).Find(&results).Error

	return
}

// GetBatchFromContent 批量查找
func (obj *_MessagesMgr) GetBatchFromContent(contents []string) (results []*Messages, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Messages{}).Where("`content` IN (?)", contents).Find(&results).Error

	return
}

// GetFromCreateAt 通过create_at获取内容
func (obj *_MessagesMgr) GetFromCreateAt(createAt time.Time) (results []*Messages, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Messages{}).Where("`create_at` = ?", createAt).Find(&results).Error

	return
}

// GetBatchFromCreateAt 批量查找
func (obj *_MessagesMgr) GetBatchFromCreateAt(createAts []time.Time) (results []*Messages, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Messages{}).Where("`create_at` IN (?)", createAts).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_MessagesMgr) FetchByPrimaryKey(id uint) (result Messages, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Messages{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchIndexByMessagesFromUserIDIndex  获取多个内容
func (obj *_MessagesMgr) FetchIndexByMessagesFromUserIDIndex(fromUserID uint) (results []*Messages, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Messages{}).Where("`from_user_id` = ?", fromUserID).Find(&results).Error

	return
}

// FetchIndexByMessagesToUserIDIndex  获取多个内容
func (obj *_MessagesMgr) FetchIndexByMessagesToUserIDIndex(toUserID uint) (results []*Messages, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Messages{}).Where("`to_user_id` = ?", toUserID).Find(&results).Error

	return
}

// FetchIndexByMessagesCreateAtIndex  获取多个内容
func (obj *_MessagesMgr) FetchIndexByMessagesCreateAtIndex(createAt time.Time) (results []*Messages, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Messages{}).Where("`create_at` = ?", createAt).Find(&results).Error

	return
}
