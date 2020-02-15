package main

import (
	"github.com/jianfengye/hade/framework/contract"
	"github.com/jianfengye/hade/framework/provider/demo"
	"github.com/jianfengye/hade/gin"
)

func main() {
	r := gin.Default()
	gin.Register(r, &demo.DemoServiceProvider{
		C: map[string]string{"foo": "bar"},
	}, true)

	r.GET("/ping", func(c *gin.Context) {
		demoService := c.Make("demo").(contract.Demo)
		val := demoService.Get("foo")
		c.JSON(200, gin.H{
			"message": val,
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
