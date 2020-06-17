package routes

import (
	"time"

	"github.com/jianfengye/hade/app/http/controllers/demo"
	"github.com/jianfengye/hade/framework/gin"
	"github.com/jianfengye/hade/framework/middleware/cors"
)

// Routes put all router here
func Routes(r *gin.Engine) {
	handler := cors.New(cors.Config{
		AllowOrigins:     []string{"https://foo.com"},
		AllowMethods:     []string{"PUT", "PATCH"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://github.com"
		},
		MaxAge: 12 * time.Hour,
	})

	r.Use(handler)
	r.GET("/ping", demo.Ping)
	r.GET("/demo", demo.Demo)
	r.GET("/demo2", demo.Demo)
}
