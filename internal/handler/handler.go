package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "pong"})
}

func CurrentTime(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"server_time": time.Now().Format(time.RFC3339)})
}
