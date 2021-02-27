package main

import (
	"TestTutenApi/src/Time/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	api := gin.Default()
	controllers.TimeByUtcController(api)
	controllers.OtherController(api)
	api.Run(":4250")
}
