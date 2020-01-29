package comment

import (
	"backend/bootstrap/connection"
	"backend/model"
	"backend/util"
	"backend/view/adapter"
	"errors"
	"github.com/gin-gonic/gin"
	"strconv"
)

func Register(router *gin.RouterGroup) {
	router.POST("comment/create", Create)
	router.POST("comment/delete", Delete)
	router.GET("comment/list", List)
	router.POST("comment/append", Append)
}

/// 创建评论
func Create(c *gin.Context) {
	topicId, _ := strconv.ParseInt(c.DefaultPostForm("topic_id", "0"), 10, 64)
	if  topicId == 0 {
		util.AbortError(c, 400, errors.New("topicId不能为空"))
		return
	}

	content, exist := c.GetPostForm("content")
	if !exist || content == "" {
		util.AbortError(c, 40001, errors.New("内容不能为空"))
		return
	}

	// TODO: 接入登陆以后获取当前用户
	meId := int64(1)

	comment := model.Comment{
		Content:   content,
		ParentID:  0,
		Status:    model.COMMENT_STATUS_REVIEWED,
		TopicID:   topicId,
		UserID:    meId,
	}
	if err := connection.Default.Save(&comment).Error; err != nil {
		util.AbortUnknownError(c, err)
		return
	}

	out := adapter.ToCommentSummary(comment, nil)
	util.ResponseSuccess(c, out)
}

// 软删除评论
func Delete(c *gin.Context) {
	commentId, _ := strconv.ParseInt(c.DefaultQuery("comment_id", "0"), 10, 64)
	if commentId == 0 {
		util.AbortError(c, 40001, errors.New("评论唯一标示不能为空"))
		return
	}

	comment, err  := model.GetComment(commentId)
	if err != nil {
		util.AbortUnknownError(c, err)
		return
	}

	if comment == nil {
		util.ResponseSuccess(c, nil)
		return
	}

	if err := model.DeleteComment(commentId); err != nil {
		util.AbortUnknownError(c, err)
		return
	}
	util.ResponseSuccess(c, nil)
}

// 显示评论
func List(c *gin.Context) {
	topicId, _ := strconv.ParseInt(c.DefaultQuery("topic_id", "0"), 10, 64)
	if topicId == 0 {
		util.AbortError(c, 400, errors.New("topicId不能为空"))
		return
	}

	// 获取所有的评论列表
	comments, err := model.GetTopicComments(topicId)
	if err != nil {
		util.AbortUnknownError(c, err)
		return
	}

	out, err := adapter.ToComments(comments)
	if err != nil {
		util.AbortUnknownError(c, err)
		return
	}

	util.ResponseSuccess(c, out)
}

func Append(c *gin.Context) {
	commentId, _ := strconv.ParseInt(c.DefaultPostForm("comment_id", "0"), 10, 64)
	if  commentId == 0 {
		util.AbortError(c, 400, errors.New("commentId不能为空"))
		return
	}

	content, exist := c.GetPostForm("content")
	if !exist || content == "" {
		util.AbortError(c, 40001, errors.New("内容不能为空"))
		return
	}

	// 判断评论是否存在，且为一级评论
	comment, err := model.GetComment(commentId)
	if err != nil {
		util.AbortUnknownError(c, err)
		return
	}
	if comment.ParentID != 0 {
		util.AbortError(c, 40001, errors.New("评论不是一级评论，不能追加"))
		return
	}

	// TODO: 接入登陆以后获取当前用户
	meId := int64(1)

	commentNew := model.Comment{
		Content:   content,
		ParentID:  comment.ID,
		Status:    model.COMMENT_STATUS_REVIEWED,
		TopicID:   comment.TopicID,
		UserID:    meId,
	}
	if err := connection.Default.Save(&commentNew).Error; err != nil {
		util.AbortUnknownError(c, err)
		return
	}

	out := adapter.ToCommentSummary(commentNew, nil)
	util.ResponseSuccess(c, out)
}