package main

import (
	"hade/bootstrap/logger"
	"hade/controller/comment"
	"hade/controller/knowledge"
	"hade/controller/tag"
	"hade/controller/topic"
	"hade/middleware"
	"flag"
	"fmt"

	"os"
	"path"
	"log"

	"hade/bootstrap/connection"
	"hade/bootstrap/config"
	"hade/util"

	"github.com/gin-gonic/gin"
)

func main() {
	// 读取配置文件
	var cf string
	flag.StringVar(&cf, "config", "", "config path")
	flag.Parse()
	if cf == "" {
		cf = path.Join(util.RootFolder(), "env.default.yaml")
		if !util.FileIsExist(cf) {
			flag.PrintDefaults()
			os.Exit(1)
		}
	}

	if err := config.Init(cf); err != nil {
		log.Fatalln(err)
	}

	// 加载日志
	if err := logger.Init(config.Default); err != nil {
		log.Fatalln(err)
	}

	// 加载数据库
	if err := connection.Init(config.Default); err != nil {
		log.Fatalln(err)
	}
	defer connection.Destory(config.Default)

	// 加载路由
	r := gin.Default()
	r.Use(middleware.CorsMiddleware())
	api := r.Group("/api")
	{
		topic.Register(api)
		comment.Register(api)
		tag.Register(api)
		knowledge.Register(api)
	}

	// 启动服务
	ip := config.Default.GetString("app.ip")
	port := config.Default.GetString("app.port")
	addr := fmt.Sprint(ip, ":", port)
	if err := r.Run(addr); err != nil {
		log.Fatalln(err)
	}
}
