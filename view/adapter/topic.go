package adapter

import (
	"backend/model"
	"backend/util"
	"backend/view/swagger/models"
	"github.com/go-openapi/strfmt"
)

// 根据用户信息获取用户
func ToTopicSummarysByTopicsWithUser(topicsWithUser []model.Topic) models.TopicSummarys  {
	topicSummaries := make([]*models.TopicSummary, len(topicsWithUser), len(topicsWithUser))
	for i, topic := range topicsWithUser{
		summary := ToTopicSummary(&topic)
		topicSummaries[i] = &summary
	}
	ret := models.TopicSummarys(topicSummaries)
	return ret
}

// 根据话题和点赞数组获取话题详情
func ToTopicByTopicWithUser(topicWithUser *model.Topic, tags []model.Tag) models.Topic {
	t := topicWithUser.CreatedAt.Format(util.TIME_ISO_NO_T)
	t2 := float32(topicWithUser.Score)
	ret := models.Topic{
		Content:   &topicWithUser.Content,
		CreatedAt: &t,
		ID:        &topicWithUser.ID,
		LikeCount: &topicWithUser.LikeCount,
		Link:      &topicWithUser.Link,
		Score:     &t2,
		Source:    &topicWithUser.Source,
		Title:     &topicWithUser.Title,
		CommentCount: &topicWithUser.CommentCount,
		User: &models.UserSummary{
			ID:         &topicWithUser.User.ID,
			Name:       &topicWithUser.User.Name,
			SmallImage: nil,
		},
	}

	outTags := make([]*models.Tag, len(tags), len(tags))
	for i, tag := range tags {
		outTags[i] = ToTag(tag)
	}

	ret.Tags = outTags
	return ret
}

// 转换为TopicSummary
func ToTopicSummary(topic *model.Topic) models.TopicSummary {
	contentSummary := topic.Content
	if len(topic.Content) > 100 {
		contentSummary = topic.Content[:100]
	}
	t1 := topic.CreatedAt.Format(util.TIME_ISO_NO_T)
	t3 := strfmt.URI(topic.Link)
	t4 := float32(topic.Score)
	summary := models.TopicSummary{
		ContentSummary: &contentSummary,
		CreatedAt:      &t1,
		ID:             &topic.ID,
		LikeCount:      &topic.LikeCount,
		CommentCount: &topic.CommentCount,
		Link:           &t3,
		Score:          &t4,
		Source:         &topic.Source,
		Title:          &topic.Title,
		User:           &models.UserSummary{
			ID:         &topic.User.ID,
			Name:       &topic.User.Name,
			SmallImage: nil,
		},
	}
	return summary
}