package model

import (
	"fmt"
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

func (t *TPlatform) FindByIds(ids []int) ([]*TPlatform, error) {
	results := make([]*TPlatform, 0)
	if e := database.DB.Where("id in ?", ids).Find(&results).Error; e != nil {
		return nil, fmt.Errorf("获取平台信息失败, ids: %v, err: %w", ids, e)
	}
	return results, nil
}
func (t *TPlatform) PluckNamesByIds(ids []int) ([]string, error) {
	results := make([]string, 0)
	if e := database.DB.Model(t).Where("id in ?", ids).Pluck("name", &results).Error; e != nil {
		return nil, fmt.Errorf("获取平台名列表失败, ids: %v, err: %w", ids, e)
	}
	return results, nil
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
	if e := platform.FirstById(t.PlatformId); e == nil {
		t.Platform = platform.Name
	}
	return nil
}
