package model

import (
	"hade/bootstrap/connection"
	"errors"
	"github.com/jinzhu/gorm"
	"time"
)

type Likeable struct {
	CreatedAt   time.Time `gorm:"column:created_at" json:"created_at"`
	ID          int64     `gorm:"column:id;primary_key" json:"id;primary_key"`
	LikableID   int64     `gorm:"column:likable_id" json:"likable_id"`
	LikableType string    `gorm:"column:likable_type" json:"likable_type"`
	UpdatedAt   time.Time `gorm:"column:updated_at" json:"updated_at"`
	UserID      int64     `gorm:"column:user_id" json:"user_id"`
}

// TableName sets the insert table name for this struct type
func (l *Likeable) TableName() string {
	return "likables"
}


// 话题喜欢查询返回结构
type TopicLike struct {
	TopicId int64 `gorm:"column:topic_id" json:"topic_id"`
	LikeCount int64 `gorm:"column:like_count" json:"like_count"`
}

// 获取Topic喜欢个数
func GetTopicLikeCount(topicIds []int64) (map[int64]int64, error) {
	if len(topicIds) == 0 {
		return nil, nil
	}

	var topicLikes []TopicLike
	if err := connection.Default.Model(&Likeable{}).Select("count(*) as like_count, likable_id as topic_id").Where("likable_type = ?", "topic").Where("likable_id in (?)", topicIds).Group("likable_id").Scan(&topicLikes).Error; !gorm.IsRecordNotFoundError(err){
		return nil, err
	}
	out := make(map[int64]int64, len(topicLikes))
	for _, tl := range topicLikes {
		out[tl.TopicId] = tl.LikeCount
	}
	return out, nil
}

// 某人是否已经喜欢过某个话题
func CheckTopicLiked(topicId int64, userId int64) (bool, error) {
	if userId == 0 || topicId == 0 {
		return false, errors.New("参数错误")
	}

	var count int
	if err := connection.Default.Model(&Likeable{}).Where("user_id = ?", userId).Where("likable_id = ?", topicId).Where("likable_type = ?", "topic").Count(count).Error; err != nil {
		return false ,err
	}

	if count > 0 {
		return true, nil
	}

	return false, nil
}

// 某人点赞某个话题
func DoTopicLike(topicId, userId int64) error {
	if userId == 0 || topicId == 0 {
		return errors.New("参数错误")
	}

	topicLike := Likeable{
		LikableID:   topicId,
		LikableType: "topic",
		UserID:      userId,
	}
	if err := connection.Default.Create(&topicLike).Error; err != nil {
		return err
	}
	return nil
}