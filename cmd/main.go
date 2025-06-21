package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/yoanesber/go-api-rate-limiter/routes"
)

func main() {
	// Get environment variables
	env := os.Getenv("ENV")
	port := os.Getenv("PORT")
	isSSL := os.Getenv("IS_SSL")
	apiVersion := os.Getenv("API_VERSION")

	if env == "" || port == "" || isSSL == "" || apiVersion == "" {
		fmt.Println("Environment variables ENV, PORT, IS_SSL, and API_VERSION must be set.")
		fmt.Println("Example: ENV=DEVELOPMENT PORT=8080 IS_SSL=FALSE API_VERSION=v1")
		return
	}

	// Set Gin mode
	gin.SetMode(gin.DebugMode)
	if env == "PRODUCTION" {
		gin.SetMode(gin.ReleaseMode)
	}

	// Setup router
	r := routes.SetupRouter()
	r.SetTrustedProxies(nil) // Set trusted proxies to nil to avoid issues with forwarded headers

	// Start the server
	if err := r.Run(":" + port); err != nil {
		fmt.Printf("Failed to start server: %v\n", err)
		return
	}
}
