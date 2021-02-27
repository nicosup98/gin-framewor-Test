package time

import (
	timemodels "TestTutenApi/src/Time/models"
	"fmt"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"

	"net/http"

	"time"
)

func PostTimeByUtc() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		body := timemodels.TimeByUtcDTO{}
		ctx.ShouldBindJSON(&body)
		fmt.Printf("date %v time %v", body.Date, body.TimeZone)
		horaRaw := strings.Split(body.Date, ":")
		hora, _ := strconv.ParseInt(horaRaw[0], 10, 0)
		minuto, _ := strconv.ParseInt(horaRaw[1], 10, 0)
		segundo, _ := strconv.ParseInt(horaRaw[2], 10, 0)
		now := time.Now()
		timeConverted := time.Date(now.Year(), now.Month(), now.Day(), int(hora), int(minuto), int(segundo), 0, time.UTC)
		timeZone, _ := strconv.ParseInt(body.TimeZone, 10, 0)
		timeToset := 3600000000000 * timeZone
		resp := map[string]map[string]string{
			"response": {
				"time": timeConverted.Add(time.Duration(timeToset)).Format("3:04:05 PM"),
				"zone": "UTC",
			},
		}
		ctx.JSON(http.StatusOK, resp)
	}
}
