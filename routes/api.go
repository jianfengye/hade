package routes

import (
	"github.com/jianfengye/hade/app/http/controllers/demo"
	"github.com/jianfengye/hade/framework/gin"
)

func Routes(r *gin.Engine) {
	r.GET("/ping", demo.Ping)
}
