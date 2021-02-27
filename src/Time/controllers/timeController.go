package controllers

import (
	time "TestTutenApi/src/Time"

	"github.com/gin-gonic/gin"
)

func TimeByUtcController(api *gin.Engine) {
	api.POST("/timeByUtc", time.PostTimeByUtc())
}

func OtherController(api *gin.Engine) {
	api.GET("/", func(ctx *gin.Context) {
		ctx.String(200, "%v", "hola")
	})
}
