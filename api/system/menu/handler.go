package menu

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
		return new(model.TMenu)
	}
	handler.NewResults = func() any {
		return &[]*model.TMenu{}
	}
}

func (h *Handler) List(ctx *gin.Context) {
	results, total, err := h.QueryListData(ctx)
	if err != nil {
		libs.HttpServerError(ctx, err.Error())
		return
	}
	// 组装api的菜单信息
	menus := results.(*[]*model.TMenu)
	apiIds := make([]int, 0)
	for _, v := range *menus {
		apiIds = append(apiIds, v.Id)
	}
	menuApiIds, err := model.QueryMenuApisByMenuIds(apiIds)
	if err != nil {
		libs.HttpServerError(ctx, err.Error())
		return
	}
	//组装api信息
	for _, v := range *menus {
		v.ApiIds = menuApiIds[v.Id]
	}

	libs.HttpListSuccess(ctx, menus, total)
}
