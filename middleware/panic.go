package middleware

import (
	"github.com/gin-gonic/gin"
)

// RecoveryMiddleware捕获所有panic，并且返回错误信息
func RecoveryMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// TODO: 先做一下日志记录

				// TODO: 返回
				return
			}
		}()
		c.Next()
	}
}