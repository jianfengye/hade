package main

import (
	"fmt"

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
		demoService2, err := c.MakeNew("demo", []interface{}{
			map[string]string{"foo": "bar2"},
		})
		if err != nil {
			infos := fmt.Sprintf("%+v", err)
			c.JSON(200, gin.H{
				"message": infos,
			})
			return
		}
		val := demoService2.(contract.Demo).Get("foo")
		c.JSON(200, gin.H{
			"message": val,
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
