package controllers

import (
	time "TestTutenApi/src/Time"

	"github.com/gin-gonic/gin"
)

func TimeByUtcController() *gin.Engine {
	api := gin.Default()
	api.POST("/timeByUtc", time.PostTimeByUtc())
	return api
}
