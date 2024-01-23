package platform_kind

import (
	"github.com/gin-gonic/gin"
)

func Routers(e *gin.RouterGroup) {
	e.GET("/platform_kinds", handler.List)
	e.GET("/platform_kind", handler.Get)
	e.POST("/platform_kind", handler.Create)
	e.PUT("/platform_kind", handler.Update)
	e.DELETE("/platform_kind", handler.Delete)
}
