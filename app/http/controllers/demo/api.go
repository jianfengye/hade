package demo

import (
	"github.com/jianfengye/hade/framework/contract"
	"github.com/jianfengye/hade/framework/gin"
)

// Ping is a function for test api
func Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "ping",
	})
}

func Demo(c *gin.Context) {
	app := c.MustMake(contract.AppKey).(contract.App)
	env := c.MustMake(contract.EnvKey).(contract.Env)
	c.JSON(200, gin.H{
		"env":          env.AppEnv(),
		"hade-version": app.Version(),
	})
}
