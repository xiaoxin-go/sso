package routers

import (
	"github.com/gin-gonic/gin"
	"sso/api/auth"
	"sso/api/platform"
	"sso/api/platform_kind"
	"sso/api/platform_user"
	"sso/api/system/api"
	"sso/api/system/log"
	"sso/api/system/menu"
	"sso/api/system/role"
	"sso/api/system/user"
	"sso/api/tenement"
)

type Option func(engine *gin.RouterGroup)

var options = make([]Option, 0)

func Include(opts ...Option) {
	options = append(options, opts...)
}

func IncludeRouter() {
	Include(auth.Routers)
	Include(user.Routers)
	Include(menu.Routers)
	Include(api.Routers)
	Include(log.Routers)
	Include(role.Routers)
	Include(tenement.Routers)
	Include(platform.Routers)
	Include(platform_user.Routers)
	Include(platform_kind.Routers)
}

// Init 初始化
func Init(r *gin.RouterGroup) *gin.RouterGroup {
	IncludeRouter()
	for _, opt := range options {
		opt(r)
	}
	return r
}
