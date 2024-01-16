package platform

import (
	"github.com/gin-gonic/gin"
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
		return new(model.TPlatform)
	}
	handler.NewResults = func() any {
		return &[]*model.TPlatform{}
	}
}

// GetSelf 获取用户有权限的平台
func (h *Handler) GetSelf(ctx *gin.Context) {
	// 1. 获取当前用户信息
	user, e := libs.GetUser(ctx)
	if e != nil {
		libs.HttpAuthorError(ctx, e.Error())
		return
	}
	// 2. 获取用户所属租户
	tenementIds, e := new(model.TTenementUser).PluckTenementIdsByUserId(user.Id)
	if e != nil {
		libs.HttpServerError(ctx, e.Error())
		return
	}
	// 3. 获取租户下的平台
	platformIds, e := new(model.TTenementPlatform).PluckPlatformIdsByTenementIds(tenementIds)
	if e != nil {
		libs.HttpServerError(ctx, e.Error())
		return
	}
	platforms, e := new(model.TPlatform).FindByIds(platformIds)
	if e != nil {
		libs.HttpServerError(ctx, e.Error())
		return
	}
	// 4. 根据平台类型区分平台
	for _, v := range platforms {
		v.
	}
	// 5. 获取用户常用平台
	// 6. 组装数据返回
}
