package platform

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"sort"
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
	fmt.Println("用户ID--->", user.Id)
	// 2. 获取用户所属租户
	tenementIds, e := new(model.TTenementUser).PluckTenementIdsByUserId(user.Id)
	if e != nil {
		libs.HttpServerError(ctx, e.Error())
		return
	}
	fmt.Println("用户租户ID--->", tenementIds)
	// 3. 获取租户下的平台
	platformIds, e := new(model.TTenementPlatform).PluckPlatformIdsByTenementIds(tenementIds)
	if e != nil {
		libs.HttpServerError(ctx, e.Error())
		return
	}
	fmt.Println("租户关联的平台ID: ", platformIds)
	platforms, e := new(model.TPlatform).FindByIds(platformIds)
	if e != nil {
		libs.HttpServerError(ctx, e.Error())
		return
	}
	fmt.Println("平台信息---->")
	kindPlatforms := make(map[string][]*model.TPlatform)
	// 4. 根据平台类型区分平台
	for _, v := range platforms {
		fmt.Printf("%+v\n", v)
		kindPlatforms[v.KindName] = append(kindPlatforms[v.KindName], v)
	}
	// 5. 获取用户常用平台
	top10PlatformIds, e := new(model.TPlatformRecord).GetTop10PlatformIdsByUserId(user.Id)
	if e != nil {
		libs.HttpServerError(ctx, e.Error())
		return
	}
	top10Platforms, e := new(model.TPlatform).FindByIds(top10PlatformIds)
	if e != nil {
		libs.HttpServerError(ctx, e.Error())
		return
	}

	// 6. 组装数据返回

	results := make([]*kindPlatformM, 0)
	results = append(results, &kindPlatformM{
		Kind:      "常用平台",
		Platforms: top10Platforms,
	})
	for kind, v := range kindPlatforms {
		results = append(results, &kindPlatformM{
			Kind:      kind,
			Platforms: v,
		})
	}
	sort.Sort(kindPlatformMs(results))
	libs.HttpSuccess(ctx, results, "ok")
	return
}

type kindPlatformM struct {
	Kind      string             `json:"kind"`
	Platforms []*model.TPlatform `json:"platforms"`
}

type kindPlatformMs []*kindPlatformM

func (k kindPlatformMs) Len() int { return len(k) }
func (k kindPlatformMs) Swap(i, j int) {
	k[i], k[j] = k[j], k[i]
}
func (k kindPlatformMs) Less(i, j int) bool {
	return k[i].Kind < k[j].Kind
}

type redirectRequest struct {
	PlatId int `json:"plat_id"`
}

// Goto 平台跳转
func (h *Handler) Goto(ctx *gin.Context) {
	// 1. 获取参数
	params := new(redirectRequest)
	if e := ctx.ShouldBindJSON(params); e != nil {
		libs.HttpParamsError(ctx, fmt.Sprintf("参数解析失败, err: %s", e.Error()))
		return
	}
	// 2. 获取平台信息
	platform := new(model.TPlatform)
	if e := platform.QueryById(params.PlatId); e != nil {
		libs.HttpServerError(ctx, e.Error())
		return
	}
	user, e := libs.GetUser(ctx)
	if e != nil {
		libs.HttpAuthorError(ctx, e.Error())
		return
	}
	// 3. 增加平台跳转次数
	record := model.TPlatformRecord{UserId: user.Id, PlatformId: platform.Id}
	record.Create()

	var url string
	if platform.IndexUrl != "" {
		url = platform.IndexUrl
	} else {
		url = platform.Url
	}

	libs.HttpSuccess(ctx, url, "跳转成功")
}
