package tenement

import (
	"github.com/gin-gonic/gin"
)

func Routers(e *gin.RouterGroup) {
	e.GET("/tenements", handler.List)
	e.GET("/tenement", handler.Get)
	e.POST("/tenement", handler.Create)
	e.PUT("/tenement", handler.Update)
	e.DELETE("/tenement", handler.Delete)
	e.GET("/tenement/users", handler.GetUsers)
	e.PUT("/tenement/users", handler.UpdateUsers)
	e.GET("/tenement/platforms", handler.GetPlatforms)
	e.PUT("/tenement/platforms", handler.UpdatePlatforms)
}
