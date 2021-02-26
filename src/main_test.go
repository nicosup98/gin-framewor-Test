package main

import (
	"TestTutenApi/src/Time/controllers"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestApi(t *testing.T) {
	sv := httptest.NewServer(controllers.TimeByUtcController())
	body := map[string]interface{}{
		"date":     "00:00:00",
		"timeZone": "UTC",
	}
	bodyBytes, _ := json.Marshal(body)
	resp, err := sv.Client().Post("/timeByUtc", "Application/json", bytes.NewReader(bodyBytes))

	if err != nil {
		panic(fmt.Sprintf("Error %v", err.Error()))
	}
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	stringResp := buf.String()
	resp.Body.Close()
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, `{"response": {"time": "4:00:00 AM","zone": "UTC"}}`, stringResp)
	sv.Close()
}
