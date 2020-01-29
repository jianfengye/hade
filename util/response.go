package util

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ErrorResponse struct {

	// 自定义错误码
	Code int64 `json:"code,omitempty"`

	// 错误信息
	Msg string `json:"msg,omitempty"`
}

// 返回未知错误
func AbortUnknownError(c *gin.Context, err error) {
	msg := "内部错误"
	if mode, existed := c.Get("error.mode"); existed && mode == "show" {
		msg = err.Error()
	}
	c.AbortWithStatusJSON(http.StatusInternalServerError, ErrorResponse{
		Code: http.StatusInternalServerError,
		Msg: msg,
	})
}

// 返回已知错误，这里的err直接显示，请确保对外显示不会泄漏
func AbortError(c *gin.Context, code int64, err error) {
	c.AbortWithStatusJSON(http.StatusInternalServerError, ErrorResponse{
		Code: code,
		Msg: err.Error(),
	})
}

func ResponseSuccess(c *gin.Context, obj interface{}) {
	c.JSON(200, obj)
}