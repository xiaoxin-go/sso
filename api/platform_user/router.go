package platform_user

import (
	"github.com/gin-gonic/gin"
)

func Routers(e *gin.RouterGroup) {
	e.GET("/platform_users", handler.List)
	e.GET("/platform_user", handler.Get)
	e.POST("/platform_user", handler.Create)
	e.PUT("/platform_user", handler.Update)
	e.DELETE("/platform_user", handler.Delete)
}
