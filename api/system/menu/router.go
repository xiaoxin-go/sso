package menu

import (
	"github.com/gin-gonic/gin"
)

func Routers(e *gin.RouterGroup) {
	e.GET("/system/menus", handler.List)
	e.GET("/system/menu", handler.Get)
	e.POST("/system/menu", handler.Create)
	e.PUT("/system/menu", handler.Update)
	e.DELETE("/system/menu", handler.Delete)
}
