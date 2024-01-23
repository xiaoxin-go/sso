package model

import (
	"fmt"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"sso/database"
)

type TPlatform struct {
	BaseModel
	Name        string `gorm:"column:name" json:"name" binding:"-"`
	NameCn      string `gorm:"column:name_cn" json:"name_cn" binding:"-"`
	Description string `gorm:"column:description" json:"description" binding:"-"`
	Url         string `gorm:"column:url" json:"url" binding:"-"`
	IndexUrl    string `gorm:"column:index_url" json:"index_url" binding:"-"`
	KindId      int    `gorm:"column:kind_id" json:"kind_id"`
	KindName    string `gorm:"-" json:"kind_name"`
	Type        int    `gorm:"column:type" json:"type" binding:"-"`
	LoginFunc   string `gorm:"column:login_func" json:"login_func" binding:"-"`
	Enabled     int    `gorm:"column:enabled" json:"enabled" binding:"-"`
}

func (t TPlatform) TableName() string {
	return "t_platform"
}

func (t *TPlatform) FirstById(id int) error {
	return firstById(t, id)
}
func (t *TPlatform) redisKeyById() string {
	return fmt.Sprintf("sso_platform_info_by_id_%d", t.Id)
}

// QueryById 带有缓存的platform信息查询
func (t *TPlatform) QueryById(id int) error {
	t.Id = id
	k := t.redisKeyById()
	fields := []string{
		"name",
		"name_cn",
		"description",
		"url",
		"index_url",
		"type",
		"login_func",
		"enabled",
	}
	return queryById(t, k, id, fields)
}

func (t *TPlatform) FilterByIdsAndName(ids []int, name string) ([]*TPlatform, error) {
	results := make([]*TPlatform, 0)
	if len(ids) == 0 {
		return results, nil
	}
	if e := database.DB.Where("id in ? and name like ?", ids, "%"+name+"%").Find(&results).Error; e != nil {
		return nil, fmt.Errorf("获取平台信息失败, ids: %v, err: %w", ids, e)
	}
	return results, nil
}

func (t *TPlatform) FindByIds(ids []int) ([]*TPlatform, error) {
	results := make([]*TPlatform, 0)
	if len(ids) == 0 {
		return results, nil
	}
	if e := database.DB.Where("id in ?", ids).Find(&results).Error; e != nil {
		return nil, fmt.Errorf("获取平台信息失败, ids: %v, err: %w", ids, e)
	}
	return results, nil
}
func (t *TPlatform) PluckNamesByIds(ids []int) ([]string, error) {
	results := make([]string, 0)
	if len(ids) == 0 {
		return results, nil
	}
	if e := database.DB.Model(t).Where("id in ?", ids).Pluck("name", &results).Error; e != nil {
		return nil, fmt.Errorf("获取平台名列表失败, ids: %v, err: %w", ids, e)
	}
	return results, nil
}
func (t *TPlatform) AfterFind(tx *gorm.DB) error {
	if t.KindId > 0 {
		kind := TPlatformKind{}
		if e := kind.QueryById(t.KindId); e == nil {
			t.KindName = kind.Name
		}
	}
	return nil
}
func (t *TPlatform) AfterUpdate(tx *gorm.DB) error {
	fmt.Println("after update------->")
	return nil
}

type TPlatformUser struct {
	BaseModel
	PlatformId int    `gorm:"column:platform_id" json:"platform_id" binding:"-"`
	Platform   string `gorm:"-" json:"platform" binding:"-"`
	Username   string `gorm:"column:username" json:"username" binding:"-"`
	Password   string `gorm:"column:password" json:"password" binding:"-"`
	IsDefault  int    `gorm:"column:is_default" json:"is_default" binding:"-"`
}

func (t TPlatformUser) TableName() string {
	return "t_platform_user"
}

func (t *TPlatformUser) FirstById(id int) error {
	return firstById(t, id)
}
func (t *TPlatformUser) redisKeyById() string {
	return fmt.Sprintf("sso_platform_user_info_by_id_%d", t.Id)
}

// QueryById 带有缓存的platform_user信息查询
func (t *TPlatformUser) QueryById(id int) error {
	t.Id = id
	k := t.redisKeyById()
	fields := []string{
		"platform_id",
		"username",
		"password",
		"is_default",
	}
	return queryById(t, k, id, fields)
}

func (t *TPlatformUser) AfterFind(tx *gorm.DB) error {
	platform := TPlatform{}
	if e := platform.QueryById(t.PlatformId); e == nil {
		t.Platform = platform.Name
	}
	return nil
}

type TPlatformKind struct {
	BaseModel
	Name        string `gorm:"column:name" json:"name" binding:"required"`
	Description string `gorm:"column:description" json:"description" binding:"required"`
}

func (t TPlatformKind) TableName() string {
	return "t_platform_kind"
}

func (t *TPlatformKind) FirstById(id int) error {
	return firstById(t, id)
}
func (t *TPlatformKind) redisKey() string {
	return fmt.Sprintf("sso_platform_kind_info_by_id_%d", t.Id)
}

// QueryById 带有缓存的platform_kind信息查询
func (t *TPlatformKind) QueryById(id int) error {
	t.Id = id
	k := t.redisKey()
	fields := []string{
		"name",
		"description",
	}
	return queryById(t, k, id, fields)
}

func (t *TPlatformKind) AfterDelete(tx *gorm.DB) (err error) {
	database.R.Del(t.redisKey())
	return
}

type TPlatformRecord struct {
	BaseModel
	UserId     int `gorm:"column:user_id" json:"user_id"`
	PlatformId int `gorm:"column:platform_id" json:"platform_id"`
}

func (t TPlatformRecord) TableName() string {
	return "t_platform_record"
}
func (t *TPlatformRecord) Create() {
	if e := database.DB.Create(t).Error; e != nil {
		zap.L().Warn("添加平台跳转记录失败", zap.Any("record", t), zap.Error(e))
	}
}

// GetTop10PlatformIdsByUserId 获取前10个经常访问的平台信息
func (t *TPlatformRecord) GetTop10PlatformIdsByUserId(userId int) ([]int, error) {
	results := make([]int, 0)
	if e := database.DB.Model(t).Where("user_id = ?", userId).Group("platform_id").Order("").Pluck("platform_id", &results).Error; e != nil {
		return nil, fmt.Errorf("获取用户前10个常用平台信息失败, user_id: %d, err: %w", userId, e)
	}
	return results, nil
}
