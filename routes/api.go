package routes

import (
	"time"

	"github.com/jianfengye/hade/app/http/controllers/demo"
	"github.com/jianfengye/hade/framework/gin"
	"github.com/jianfengye/hade/framework/middleware/cors"
	"github.com/jianfengye/hade/framework/middleware/gzip"
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
	r.Use(gzip.Gzip(gzip.DefaultCompression, gzip.WithExcludedExtensions([]string{".pdf", ".mp4"})))
	r.GET("/ping", demo.Ping)
	r.GET("/demo", demo.Demo)
	r.GET("/demo2", demo.Demo)
}
