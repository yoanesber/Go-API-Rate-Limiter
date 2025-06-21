package ratelimiter

import (
	"fmt"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

/**
 * RateLimiter is a middleware that implements a rate limiter for HTTP requests.
 * It allows a burst of requests and then limits the rate of requests to a specified limit.
 * It uses a map to store visitors and their last seen time, allowing for efficient rate limiting.
 * The rate limiter is based on the token bucket algorithm, where tokens are added at a specified rate,
 * and each request consumes a token. If no tokens are available, the request is rejected with a 429 Too Many Requests status.
 */

var visitors = make(map[string]*rate.Limiter)
var lastSeen = make(map[string]time.Time)

var mu sync.Mutex

// GetVisitor retrieves the visitor from the map or creates a new one if it doesn't exist.
// It updates the last seen time and returns the rate limiter for that visitor.
func GetVisitor(c *gin.Context, r rate.Limit, b int) *rate.Limiter {
	now := time.Now()

	// Set key to the visitor
	ip := c.ClientIP()
	method := c.Request.Method
	path := c.Request.URL.Path
	key := fmt.Sprintf("%s:%s:%s", ip, method, path)

	// Check if the visitor exists in the map
	// If it doesn't exist, create a new rate limiter and add it to the map
	mu.Lock()
	limiter, exists := visitors[key]
	if !exists {
		limiter = rate.NewLimiter(r, b)
		visitors[key] = limiter
	}
	lastSeen[key] = now
	mu.Unlock()

	return limiter
}

// StartVisitorCleanup starts a goroutine that cleans up expired visitors
// every minute. It checks if the last seen time of each visitor exceeds the expiration duration.
func StartVisitorCleanup(expireAfter time.Duration) {
	go func() {
		for {
			time.Sleep(time.Minute)

			// Check if the last seen time of each visitor exceeds the expiration duration
			// If it does, remove the visitor from the map
			mu.Lock()
			for key, t := range lastSeen {
				if time.Since(t) > expireAfter {
					// Remove expired visitors from the map
					delete(visitors, key)
					delete(lastSeen, key)
				}
			}
			mu.Unlock()
		}
	}()
}

// RateLimiter is a Gin middleware that implements a rate limiter for HTTP requests.
// It allows a burst of requests and then limits the rate of requests to a specified limit.
func RateLimiter(r rate.Limit, burst int, expireAfter time.Duration) gin.HandlerFunc {
	StartVisitorCleanup(expireAfter)

	return func(c *gin.Context) {
		limiter := GetVisitor(c, r, burst)

		// fmt.Printf(">>>>> Visitors values: %v\n", visitors)
		// fmt.Printf(">>>>> Last seen values: %v\n", lastSeen)
		// fmt.Printf(">>>>> Rate limit: %v\n", limiter.Limit())
		// fmt.Printf(">>>>> Burst size: %v\n", limiter.Burst())
		// fmt.Printf(">>>>> Current tokens: %v\n", limiter.Tokens())
		// fmt.Printf(">>>>> Remaining tokens: %v\n", limiter.Burst()-int(limiter.Tokens()))

		if !limiter.Allow() {
			c.AbortWithStatusJSON(429, gin.H{
				"error":   "Too Many Requests",
				"message": "You have exceeded the rate limit. Please try again later.",
			})

			return
		}

		c.Next()
	}
}
