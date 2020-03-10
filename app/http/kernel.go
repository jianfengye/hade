package http

import (
	"github.com/jianfengye/hade/routes"

	"github.com/jianfengye/hade/framework"
	"github.com/jianfengye/hade/framework/gin"
	"github.com/jianfengye/hade/framework/provider/app"
	"github.com/jianfengye/hade/framework/provider/config"
	"github.com/jianfengye/hade/framework/provider/env"
	"github.com/jianfengye/hade/framework/provider/log"
)

// RunHttp is command
func RunHttp(container framework.Container) (*gin.Engine, error) {
	r := gin.Default()
	gin.Register(r, &app.HadeAppProvider{}, true)
	gin.Register(r, &env.HadeEnvProvider{}, true)
	gin.Register(r, &config.HadeConfigProvider{}, true)
	gin.Register(r, &log.HadeLogServiceProvider{}, true)

	routes.Routes(r)
	return r, nil
}
