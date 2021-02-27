package main

import (
	"TestTutenApi/src/Time/controllers"
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestApi(t *testing.T) {
	rc := httptest.NewRecorder()
	_, eng := gin.CreateTestContext(rc)
	controllers.TimeByUtcController(eng)
	body, _ := json.Marshal(map[string]interface{}{
		"date":     "'00:00:00'",
		"timeZone": "4",
	})
	req := httptest.NewRequest("POST", "/timeByUtc", bytes.NewReader(body))
	eng.ServeHTTP(rc, req)

	assert.Equal(t, 200, rc.Code)
	assert.Equal(t, "{\"response\":{\"time\":\"4:00:00 AM\",\"zone\":\"UTC\"}}", rc.Body.String())

}
