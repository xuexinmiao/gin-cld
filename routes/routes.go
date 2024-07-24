package routes

import (
	"CLD/logger"

	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, "Hello World")
	})
	return r
}
