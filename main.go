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
	}, false)

	r.GET("/ping", func(c *gin.Context) {
		//demoService := c.Make("demo").(contract.Demo)
		demoService2 := c.MakeNew("demo", []interface{}{
			map[string]string{"foo": "bar2"},
		}).(contract.Demo)
		val := demoService2.Get("foo")
		c.JSON(200, gin.H{
			"message": val,
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
