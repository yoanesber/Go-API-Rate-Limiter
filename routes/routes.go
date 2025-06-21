package routes

import (
	"time"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"

	"github.com/yoanesber/go-api-rate-limiter/internal/handler"
	"github.com/yoanesber/go-api-rate-limiter/pkg/middleware/headers"
	"github.com/yoanesber/go-api-rate-limiter/pkg/middleware/ratelimiter"
)

// SetupRouter initializes the router and sets up the routes for the application.
func SetupRouter() *gin.Engine {
	// Create a new Gin router instance
	r := gin.Default()

	// Set up middleware for the router
	r.Use(
		headers.SecurityHeaders(),
		headers.CorsHeaders(),
		headers.ContentType(),
		gzip.Gzip(gzip.DefaultCompression),
	)

	// Set up the API version 1 routes
	api := r.Group("/api")
	{
		// Apply rate limiting middleware to the API routes
		// This will limit the rate of requests to the specified endpoints
		// 1 token every 5 seconds
		// Possible 2 requests every 5 seconds, with a maximum burst of 2 requests
		// The rate limit will reset every 5 minutes
		api.GET("/ping", ratelimiter.RateLimiter(rate.Every(5*time.Second), 2, 10*time.Second), handler.Ping)
		api.GET("/time", ratelimiter.RateLimiter(rate.Every(5*time.Second), 2, 10*time.Second), handler.CurrentTime)
	}

	// NoRoute handler for undefined routes
	// This handler will be called when no other route matches the request
	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{
			"error":   "Not Found",
			"message": "The requested resource could not be found",
		})
	})

	// NoMethod handler for unsupported HTTP methods
	// This handler will be called when a request method is not allowed for the requested resource
	r.NoMethod(func(c *gin.Context) {
		c.JSON(405, gin.H{
			"error":   "Method Not Allowed",
			"message": "The requested method is not allowed for this resource",
		})
	})

	return r
}
