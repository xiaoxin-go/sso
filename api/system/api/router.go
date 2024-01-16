package api

import (
	"github.com/gin-gonic/gin"
)

func Routers(e *gin.RouterGroup) {
	e.GET("/system/apis", handler.List)
	e.GET("/system/api", handler.Get)
	e.POST("/system/api", handler.Create)
	e.PUT("/system/api", handler.Update)
	e.DELETE("/system/api", handler.Delete)
}
