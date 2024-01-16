package api

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
		return new(model.TApi)
	}
	handler.NewResults = func() any {
		return &[]*model.TApi{}
	}
}

func (h *Handler) List(ctx *gin.Context) {
	results, total, err := h.QueryListData(ctx)
	if err != nil {
		libs.HttpServerError(ctx, err.Error())
		return
	}
	// 组装api的菜单信息
	apis := results.(*[]*model.TApi)
	apiIds := make([]int, 0)
	for _, v := range *apis {
		apiIds = append(apiIds, v.Id)
	}
	menuIds, err := model.QueryApiMenusByApiIds(apiIds)
	if err != nil {
		libs.HttpServerError(ctx, err.Error())
		return
	}
	//组装菜单信息
	for _, v := range *apis {
		v.MenuIds = menuIds[v.Id]
		for _, mi := range v.MenuIds {
			menu := model.TMenu{}
			if err := menu.QueryById(mi); err != nil {
				libs.HttpServerError(ctx, err.Error())
				return
			} else {
				v.Menus = append(v.Menus, menu.Name)
			}
		}
	}
	libs.HttpListSuccess(ctx, apis, total)
}
