package log

import (
	"github.com/gin-gonic/gin"
)

func Routers(e *gin.RouterGroup) {
	e.POST("/system/log/list", handler.List)
}
