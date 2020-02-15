package main

import (
	hade "github.com/jianfengye/hade/framework"
	provider "github.com/jianfengye/hade/framework/provider"
	"github.com/jianfengye/hade/gin"
)

func main() {
	r := gin.Default()

	gin.RegisterService(r, hade.CONFIG_SERVICE, provider.DefaultViperProvider, true, []interface{}{
		"path",
	})

	r.GET("/ping", func(c *gin.Context) {
		val := c.ConfigService().GetStr("key")
		c.JSON(200, gin.H{
			"message": val,
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
