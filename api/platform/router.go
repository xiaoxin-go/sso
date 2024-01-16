package platform

import (
	"github.com/gin-gonic/gin"
)

func Routers(e *gin.RouterGroup) {
	e.GET("/platforms", handler.List)
	e.GET("/platform", handler.Get)
	e.POST("/platform", handler.Create)
	e.PUT("/platform", handler.Update)
	e.DELETE("/platform", handler.Delete)
	e.GET("/platform/self", handler.GetSelf)
}
