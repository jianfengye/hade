package knowledge

import (
	"backend/model"
	"backend/util"
	"backend/view/adapter"
	"errors"
	"github.com/gin-gonic/gin"
)

func Register(router *gin.RouterGroup) {
	router.GET("knowledge/daily", Daily)
}

// 每日知识点
func Daily(c *gin.Context) {
	day := c.DefaultQuery("day", "")
	if day == "" {
		util.AbortError(c, 500, errors.New("day不能为空"))
		return
	}
	daily, err := model.GetDaily(day)
	if err != nil {
		util.AbortUnknownError(c, err)
		return
	}
	if daily == nil {
		util.AbortError(c, 500, errors.New("没有当日的信息"))
		return
	}
	cont, err := daily.ParseContent()
	if err != nil {
		util.AbortError(c, 500, errors.New("解析当日信息错误"))
		return
	}

	ret := adapter.ToDaily(daily, cont)

	util.ResponseSuccess(c, ret)
}
