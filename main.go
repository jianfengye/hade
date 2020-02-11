package main

import (
	_ "github.com/jianfengye/hade/framework"
	serviceProvider "github.com/jianfengye/hade/framework/provider"
	"github.com/jianfengye/hade/gin"
)

func main() {
	r := gin.Default()

	r.UseService(map[string]Provider{
		CONFIG_SERVICE: serviceProvider.DefaultViperProvider,
	})

	r.GET("/ping", func(c *gin.Context) {
		val := c.ServiceConfig().GETStr("key")
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
