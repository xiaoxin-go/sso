package model

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"sso/database"
	"time"
)

type TTenement struct {
	BaseModel
	Name        string   `gorm:"column:name" json:"name" binding:"-"`
	Description string   `gorm:"column:description" json:"description" binding:"-"`
	Users       []string `gorm:"-" json:"users"`
	UserIds     []int    `gorm:"-" json:"user_ids"`
	Platforms   []string `gorm:"-" json:"platforms"`
	PlatformIds []int    `gorm:"-" json:"platform_ids"`
}

func (t TTenement) TableName() string {
	return "t_tenement"
}

// AfterFind 获取租户关联的平台和用户信息
func (t *TTenement) AfterFind(tx *gorm.DB) (err error) {
	// 获取租户关联的用户信息
	userIds, err := new(TTenementUser).PluckUserIdsByTenementId(t.Id)
	if err != nil {
		return
	}
	userNames, err := new(TUser).PluckNamesByIds(userIds)
	if err != nil {
		return
	}
	t.Users = userNames
	// 获取租户关联的平台信息
	platformIds, err := new(TTenementPlatform).PluckPlatformIdsByTenementId(t.Id)
	if err != nil {
		return
	}
	platformNames, err := new(TPlatform).PluckNamesByIds(platformIds)
	if err != nil {
		return
	}
	t.Platforms = platformNames
	return
}

func (t *TTenement) FirstById(id int) error {
	if e := database.DB.First(t, id).Error; errors.Is(e, gorm.ErrRecordNotFound) {
		return fmt.Errorf("租户信息不存在, id: %d", id)
	} else if e != nil {
		return fmt.Errorf("获取租户信息失败, id: %d, err: %w", id, e)
	}
	return nil
}
func (t *TTenement) redisKeyById() string {
	return fmt.Sprintf("sso_tenement_info_by_id_%d", t.Id)
}

// QueryById 带有缓存的tenement信息查询
func (t *TTenement) QueryById(id int) error {
	t.Id = id
	k := t.redisKeyById()
	fields := []string{
		"name",
		"description",
	}
	results, err := database.R.HMGet(k, fields...).Result()
	if err != nil || results[0] == nil {
		if e1 := t.FirstById(id); e1 != nil {
			return e1
		}
		database.R.HMSet(k, toMap(t, fields))
		database.R.Expire(k, 2*time.Hour)
	} else {
		setData(t, fields, results)
	}
	return nil
}

type TTenementPlatform struct {
	BaseModel
	TenementId int `gorm:"column:tenement_id" json:"tenement_id" binding:"-"`
	PlatformId int `gorm:"column:platform_id" json:"platform_id" binding:"-"`
}

func (t TTenementPlatform) TableName() string {
	return "t_tenement_platform"
}

func (t *TTenementPlatform) PluckPlatformIdsByTenementId(tenementId int) ([]int, error) {
	result := make([]int, 0)
	if e := database.DB.Model(t).Where("tenement_id = ?", tenementId).Pluck("platform_id", &result).Error; e != nil {
		return nil, fmt.Errorf("获取租户关联的平台ID列表失败, tenement_id: %d, err: %w", tenementId, e)
	}
	return result, nil
}
func (t *TTenementPlatform) PluckPlatformIdsByTenementIds(tenementIds []int) ([]int, error) {
	result := make([]int, 0)
	if e := database.DB.Model(t).Where("tenement_id in ?", tenementIds).Pluck("platform_id", &result).Error; e != nil {
		return nil, fmt.Errorf("获取租户关联的平台ID列表失败, tenement_ids: %d, err: %w", tenementIds, e)
	}
	return result, nil
}

func (t *TTenementPlatform) Update(tenementId int, platformIds []int, operator string) error {
	oldPlatformIds, e := t.PluckPlatformIdsByTenementId(tenementId)
	if e != nil {
		return e
	}
	tx := database.DB.Begin()
	if e := tx.Delete(t, "tenement_id = ?", tenementId).Error; e != nil {
		tx.Rollback()
		return fmt.Errorf("清除租户关联的平台信息失败, err: %w", e)
	}
	bulks := make([]*TTenementPlatform, 0)
	for _, v := range platformIds {
		bulks = append(bulks, &TTenementPlatform{
			TenementId: tenementId,
			PlatformId: v,
		})
	}
	if len(bulks) > 0 {
		if e := tx.Create(bulks).Error; e != nil {
			tx.Rollback()
			return fmt.Errorf("关联租户和平台失败, err: %w", e)
		}
	}
	tx.Commit()
	AddLog(operator, "关联租户与平台, tenementId: %d, oldPlatformIds: %v, newPlatformIds: %v", tenementId, oldPlatformIds, platformIds)
	return nil
}

type TTenementUser struct {
	BaseModel
	TenementId int `gorm:"column:tenement_id" json:"tenement_id" binding:"-"`
	UserId     int `gorm:"column:user_id" json:"user_id" binding:"-"`
}

func (t TTenementUser) TableName() string {
	return "t_tenement_user"
}

func (t *TTenementUser) PluckUserIdsByTenementId(tenementId int) ([]int, error) {
	result := make([]int, 0)
	if e := database.DB.Model(t).Where("tenement_id = ?", tenementId).Pluck("user_id", &result).Error; e != nil {
		return nil, fmt.Errorf("获取租户关联的用户ID列表失败, tenement_id: %d, err: %w", tenementId, e)
	}
	return result, nil
}
func (t *TTenementUser) PluckTenementIdsByUserId(userId int) ([]int, error) {
	result := make([]int, 0)
	if e := database.DB.Model(t).Where("user_id = ?", userId).Pluck("tenement_id", &result).Error; e != nil {
		return nil, fmt.Errorf("获取用户关联的租户ID列表失败, user_id: %d, err: %w", userId, e)
	}
	return result, nil
}
func (t *TTenementUser) Update(tenementId int, userIds []int, operator string) error {
	// 1. 获取旧关联数据
	oldUserIds, e := t.PluckUserIdsByTenementId(tenementId)
	if e != nil {
		return e
	}
	tx := database.DB.Begin()
	if e := tx.Delete(t, "tenement_id = ?", tenementId).Error; e != nil {
		tx.Rollback()
		return fmt.Errorf("清除租户关联的用户信息失败, err: %w", e)
	}
	bulks := make([]*TTenementUser, 0)
	for _, v := range userIds {
		bulks = append(bulks, &TTenementUser{
			TenementId: tenementId,
			UserId:     v,
			BaseModel:  BaseModel{CreatedBy: operator},
		})
	}
	if len(bulks) > 0 {
		if e := tx.Create(bulks).Error; e != nil {
			tx.Rollback()
			return fmt.Errorf("关联租户和用户失败, err: %w", e)
		}
	}
	tx.Commit()
	AddLog(operator, "关联租户与用户, tenementId: %d, oldUserIds: %v, newUserIds: %v", tenementId, oldUserIds, userIds)
	return nil
}
