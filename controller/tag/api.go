package tag

import (
	"hade/model"
	"hade/util"
	"hade/view/adapter"
	"hade/view/swagger/models"
	"errors"
	"github.com/gin-gonic/gin"
	"math"
	"strconv"
)

func Register(router *gin.RouterGroup) {
	router.GET("/tag/list", List)
	router.GET("/tag/topics", Topics)
}

// 获取话题列表
func Topics(c *gin.Context) {
	// 参数验证
	size, _ := strconv.ParseInt(c.DefaultQuery("size", "20"), 10, 64)
	page, _ := strconv.ParseInt(c.DefaultQuery("page", "1"), 10, 64)
	if page == 0 {
		util.AbortError(c, 400, errors.New("page不应该为0"))
		return
	}
	tagId, _ := strconv.ParseInt(c.DefaultQuery("tag_id", "20"), 10, 64)

	// 获取数据
	var sum int64
	offset := (page - 1) * size
	topics, sum, err := model.GetTopicsByTagPager(tagId, offset, size)
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
	util.ResponseSuccess(c, models.GetTagTopicsOKBody{Topics: retTopics, Pager: &pager})
}

func List(c *gin.Context) {
	tags, err := model.GetAllTags()
	if err != nil {
		util.AbortUnknownError(c, err)
	}

	out := adapter.ToTags(tags)
	util.ResponseSuccess(c, out)
}