package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/yoanesber/go-api-rate-limiter/routes"
)

func TestRateLimiter_Ping_ReturnsOK(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Create a new router instance
	r := routes.SetupRouter()

	// Create a test server
	ts := httptest.NewServer(r)
	defer ts.Close()

	client := &http.Client{}

	// Send a single request to /ping, should succeed
	resp, err := client.Get(ts.URL + "/api/ping")
	if err != nil {
		t.Fatalf("Failed to make request: %v", err)
	}
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode, "Expected 200 OK for /ping")
}

func TestRateLimiter_Time_Returns429AfterBurst(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Create a new router instance
	r := routes.SetupRouter()

	// Create a test server
	ts := httptest.NewServer(r)
	defer ts.Close()

	client := &http.Client{}

	// Test the rate limiter with too many requests
	var tooManyCount int
	for i := 0; i < 5; i++ {
		resp, err := client.Get(ts.URL + "/api/time")
		if err != nil {
			t.Fatalf("Failed to make request: %v", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode == http.StatusTooManyRequests {
			tooManyCount++
		}
	}

	assert.GreaterOrEqual(t, tooManyCount, 1, "Expected at least one 429 Too Many Requests response")
}
