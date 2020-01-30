package topic

import (
	"errors"
	"github.com/gin-gonic/gin"
	"hade/util"
	"strconv"
)

// Register 注册路由
func Register(router *gin.RouterGroup) {
	router.GET("/topic/list", List)
}

// List 获取话题列表
func List(c *gin.Context) {
	// 参数验证
	size, _ := strconv.ParseInt(c.DefaultQuery("size", "20"), 10, 64)
	page, _ := strconv.ParseInt(c.DefaultQuery("page", "1"), 10, 64)
	if page == 0 || size == 0 {
		util.AbortError(c, 400, errors.New("page不应该为0"))
		return
	}

	// 返回数据
	util.ResponseSuccess(c, nil)
}
