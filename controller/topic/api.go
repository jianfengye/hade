package topic

import (
	"backend/bootstrap/connection"
	"backend/model"
	"backend/util"
	"backend/view/adapter"
	"backend/view/swagger/models"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/jianfengye/collection"
	"math"
	"strconv"
)

// Register 注册路由
func Register(router *gin.RouterGroup) {
	router.GET("/topic/list", List)
	router.GET("/topic/detail", Detail)
	router.POST("/topic/create", Create)
	router.POST("/topic/update", Update)
	router.POST("/topic/delete", Delete)
	router.POST("/topic/like", Like)
}

// List 获取话题列表
func List(c *gin.Context) {
	// 参数验证
	size, _ := strconv.ParseInt(c.DefaultQuery("size", "20"), 10, 64)
	page, _ := strconv.ParseInt(c.DefaultQuery("page", "1"), 10, 64)
	if page == 0 {
		util.AbortError(c, 400, errors.New("page不应该为0"))
		return
	}

	// 获取数据
	var sum int64
	offset := (page - 1) * size
	topics, sum, err := model.GetTopicsByPager(offset, size)
	if err != nil {
		util.AbortUnknownError(c, err)
		return
	}

	// 组织返回数据
	retTopics := adapter.ToTopicSummarysByTopicsWithUser(topics)

	totalPage := int64(math.Ceil(float64(sum) / float64(size)))
	t1 := int64(page)
	t2 := int64(size)
	pager := models.Pager{
		Page: &t1,
		Size: &t2,
		TotalPage: &totalPage,
	}

	// 返回数据
	util.ResponseSuccess(c, models.GetTopicListOKBody{Topics: retTopics, Pager: &pager})
}

// Detail 获取话题详情
func Detail(c *gin.Context) {
	// 参数验证
	topicId, _ := strconv.ParseInt(c.DefaultQuery("topic_id", "0"), 10, 64)
	if  topicId == 0 {
		util.AbortError(c, 400, errors.New("topicId不能为空"))
		return
	}

	topic := model.GetTopicWithUser(topicId)
	if topic == nil {
		util.AbortError(c, 40001, errors.New("话题不存在"))
		return
	}

	tags, err := model.GetTopicTags(topicId)
	if err != nil {
		util.AbortUnknownError(c, err)
		return
	}

	out := adapter.ToTopicByTopicWithUser(topic, tags)
	util.ResponseSuccess(c, out)
}

// Create 创建话题
func Create(c *gin.Context) {
	title, exist := c.GetPostForm("title")
	if !exist || title == "" {
		util.AbortError(c, 40001, errors.New("标题不能为空"))
		return
	}
	content, exist := c.GetPostForm("content")
	if !exist || content == "" {
		util.AbortError(c, 40001, errors.New("内容不能为空"))
		return
	}
	link, exist := c.GetPostForm("link")
	if !exist {
		util.AbortError(c, 40001, errors.New("必须传递link参数，空字符串代表置空"))
		return
	}
	tags, exist := c.GetPostFormArray("tag_ids")
	if !exist {
		util.AbortError(c, 40001, errors.New("标签必须传递，空代表置空"))
	}
	tagColl := collection.NewStrCollection(tags)
	tagIds, err := tagColl.Map(func(item interface{}, key int) interface{} {
		t := item.(string)
		i, err := strconv.ParseInt(t, 10, 64)

		if err != nil {
			tagColl.SetErr(err)
			return nil
		}
		return i
	}).ToInt64s()
	if err != nil {
		util.AbortUnknownError(c, err)
		return
	}

	// TODO: 接入登陆以后获取当前用户
	meId := int64(1)

	topic := model.Topic{}
	topic.Title = title
	topic.Content = content
	topic.Link = link
	topic.Status = model.TOPIC_STATUS_REVIEWED
	topic.UserID = meId

	if err := connection.Default.Save(&topic).Error; err != nil {
		util.AbortUnknownError(c, err)
		return
	}

	// 保存话题的topic
	if err := model.AddTagsToTopic(topic.ID, tagIds); err != nil {
		util.AbortUnknownError(c, err)
		return
	}

	out := adapter.ToTopicSummary(&topic)
	util.ResponseSuccess(c, out)
}

// Update 更新话题
func Update(c *gin.Context) {
	// 参数验证
	topicId, _ := strconv.ParseInt(c.DefaultPostForm("topic_id", "0"), 10, 64)
	if  topicId == 0 {
		util.AbortError(c, 400, errors.New("topicId不能为空"))
		return
	}

	topic := model.GetTopic(topicId)
	if topic == nil {
		util.AbortError(c, 40001, errors.New("话题不存在"))
		return
	}

	title, exist := c.GetPostForm("title")
	if !exist || title == "" {
		util.AbortError(c, 40001, errors.New("标题不能为空"))
		return
	}
	content, exist := c.GetPostForm("content")
	if !exist || content == "" {
		util.AbortError(c, 40001, errors.New("内容不能为空"))
		return
	}
	link, exist := c.GetPostForm("link")
	if !exist {
		util.AbortError(c, 40001, errors.New("必须传递link参数，空字符串代表置空"))
		return
	}
	tags, exist := c.GetPostFormArray("tag_ids")
	if !exist {
		util.AbortError(c, 40001, errors.New("标签必须传递，空代表置空"))
		return
	}
	tagColl := collection.NewStrCollection(tags)
	tagIds, err := tagColl.Map(func(item interface{}, key int) interface{} {
		t := item.(string)
		i, err := strconv.ParseInt(t, 10, 64)

		if err != nil {
			tagColl.SetErr(err)
			return nil
		}
		return i
	}).ToInt64s()
	if err != nil {
		util.AbortUnknownError(c, err)
		return
	}

	topic.Title = title
	topic.Content = content
	topic.Link = link

	if err := connection.Default.Save(&topic).Error; err != nil {
		util.AbortUnknownError(c, err)
		return
	}

	if err := model.UpdateTopicTags(topicId, tagIds); err != nil {
		util.AbortUnknownError(c, err)
		return
	}

	util.ResponseSuccess(c, nil)
}

// Delete 删除话题，软删除
func Delete(c *gin.Context) {
	// 参数验证
	topicId, _ := strconv.ParseInt(c.DefaultPostForm("topic_id", "0"), 10, 64)
	if  topicId == 0 {
		util.AbortError(c, 400, errors.New("topicId不能为空"))
		return
	}

	topic := model.GetTopic(topicId)
	if topic == nil {
		util.AbortError(c, 40001, errors.New("话题不存在"))
		return
	}

	// TODO
	meId := int64(1)
	if topic.UserID != meId {
		util.AbortError(c, 40001, errors.New("你不是话题创建者，不能删除话题"))
	}

	if err := model.DeleteTopic(topicId); err != nil {
		util.AbortUnknownError(c, err)
		return
	}

	util.ResponseSuccess(c, nil)
}

// Like 为某个话题点赞
func Like(c *gin.Context) {
	// 参数验证
	topicId, _ := strconv.ParseInt(c.DefaultPostForm("topic_id", "0"), 10, 64)
	if  topicId == 0 {
		util.AbortError(c, 400, errors.New("topicId不能为空"))
		return
	}

	topic := model.GetTopic(topicId)
	if topic == nil {
		util.AbortError(c, 40001, errors.New("话题不存在"))
		return
	}

	// TODO: 接入登陆以后获取当前用户
	meId := int64(1)

	exist, err := model.CheckTopicLiked(topic.ID, meId)
	if err != nil {
		util.AbortUnknownError(c, err)
		return
	}

	if exist {
		util.AbortError(c, 40002, errors.New("你已经点赞过了"))
		return
	}
	if err := model.DoTopicLike(topic.ID, meId); err != nil {
		util.AbortUnknownError(c, err)
		return
	}

	util.ResponseSuccess(c, nil)
}

