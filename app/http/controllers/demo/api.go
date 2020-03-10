package demo

import (
	"github.com/jianfengye/hade/framework/gin"
)

// Ping is a function for test api
func Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "ping",
	})
}
