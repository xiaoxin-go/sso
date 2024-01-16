package tenement

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"sso/libs"
	"sso/model"
)

type Handler struct {
	libs.Controller
}

var handler *Handler

func init() {
	handler = &Handler{}
	handler.NewInstance = func() libs.Instance {
		return &model.TTenement{}
	}
	handler.NewResults = func() any {
		return &[]*model.TTenement{}
	}
}

// GetUsers 获取租户下的用户信息
func (h *Handler) GetUsers(ctx *gin.Context) {
	id, e := h.GetId(ctx)
	if e != nil {
		libs.HttpParamsError(ctx, e.Error())
		return
	}
	// 获取租户关联的用户列表
	userIds, e := new(model.TTenementUser).PluckUserIdsByTenementId(id)
	if e != nil {
		libs.HttpServerError(ctx, e.Error())
		return
	}
	// 获取用户信息
	users, e := new(model.TUser).FindByIds(userIds)
	if e != nil {
		libs.HttpServerError(ctx, e.Error())
		return
	}
	libs.HttpSuccess(ctx, users, "ok")
}

type updateUsersRequest struct {
	TenementId int   `json:"tenement_id"`
	UserIds    []int `json:"user_ids"`
}

// UpdateUsers 更新租户下用户信息
func (h *Handler) UpdateUsers(ctx *gin.Context) {
	user, e := libs.GetUser(ctx)
	if e != nil {
		libs.HttpAuthorError(ctx, e.Error())
		return
	}
	l := zap.L().With(zap.String("func", "Tenement.UpdateUsers"))
	l.Info("读取参数")
	// 读取参数
	req := updateUsersRequest{}
	if e := ctx.ShouldBindJSON(&req); e != nil {
		libs.HttpParamsError(ctx, fmt.Sprintf("读取参数失败, err: %s", e.Error()))
		return
	}
	l.Info("更新租户和用户关联")
	// 更新关联数据
	if e := new(model.TTenementUser).Update(req.TenementId, req.UserIds, user.Username); e != nil {
		l.Info("更新失败", zap.Error(e))
		libs.HttpServerError(ctx, e.Error())
		return
	}
	l.Info("更新成功")
	libs.HttpSuccess(ctx, nil, "更新成功")
}

// GetPlatforms 获取租户下的平台信息
func (h *Handler) GetPlatforms(ctx *gin.Context) {
	id, e := h.GetId(ctx)
	if e != nil {
		libs.HttpParamsError(ctx, e.Error())
		return
	}
	// 获取租户下的平台ID列表
	platformIds, e := (&model.TTenementPlatform{}).PluckPlatformIdsByTenementId(id)
	if e != nil {
		libs.HttpServerError(ctx, e.Error())
		return
	}
	// 获取平台信息
	platforms, e := (&model.TPlatform{}).FindByIds(platformIds)
	if e != nil {
		libs.HttpServerError(ctx, e.Error())
		return
	}
	libs.HttpSuccess(ctx, platforms, "ok")
}

type updatePlatformsRequest struct {
	TenementId  int   `json:"tenement_id"`
	PlatformIds []int `json:"platform_ids"`
}

// UpdatePlatforms 更新租户下平台信息
func (h *Handler) UpdatePlatforms(ctx *gin.Context) {
	user, e := libs.GetUser(ctx)
	if e != nil {
		libs.HttpAuthorError(ctx, e.Error())
		return
	}
	l := zap.L().With(zap.String("func", "Tenement.UpdatePlatforms"))
	l.Info("读取参数")
	// 读取参数
	req := updatePlatformsRequest{}
	if e := ctx.ShouldBindJSON(&req); e != nil {
		libs.HttpParamsError(ctx, fmt.Sprintf("读取参数失败, err: %s", e.Error()))
		return
	}
	l.Info("更新租户和平台关联")
	// 更新关联数据
	if e := new(model.TTenementPlatform).Update(req.TenementId, req.PlatformIds, user.Username); e != nil {
		l.Info("更新失败", zap.Error(e))
		libs.HttpServerError(ctx, e.Error())
		return
	}
	l.Info("更新成功")
	libs.HttpSuccess(ctx, nil, "更新成功")
}
