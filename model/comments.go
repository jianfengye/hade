package model

import (
	"hade/bootstrap/connection"
	"github.com/jinzhu/gorm"
	"time"
)

const (
	COMMENT_STATUS_REVIEWED   = "reviewed"
	COMMENT_STATUS_UNREVIEWED = "unreviewed"
	COMMNET_STATUS_DELETED    = "deleted"
)

type Comment struct {
	Content   string    `gorm:"column:content" json:"content"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	ID        int64     `gorm:"column:id;primary_key" json:"id;primary_key"`
	LeftID    int64     `gorm:"column:left_id" json:"left_id"`
	LikeCount int64     `gorm:"column:like_count" json:"like_count"`
	ParentID  int64     `gorm:"column:parent_id" json:"parent_id"`
	RightID   int64     `gorm:"column:right_id" json:"right_id"`
	Status    string    `gorm:"column:status" json:"status"`
	TopicID   int64     `gorm:"column:topic_id" json:"topic_id"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
	UserID    int64     `gorm:"column:user_id" json:"user_id"`
	User      User   `gorm:"ForeignKey:user_id"`
}

// TableName sets the insert table name for this struct type
func (c *Comment) TableName() string {
	return "comments"
}

func GetComment(id int64) (*Comment, error) {
	comment := Comment{}
	if err := connection.Default.Model(&Comment{}).First(&comment, id).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}
	return &comment, nil
}

// 删除一个话题
func DeleteComment(id int64) error {
	if err := connection.Default.Model(&Comment{}).Set("status", TOPIC_STATUS_DELETED).Where("id = ?", id).Error; err != nil {
		return err
	}
	return nil
}


// 获取话题的标签
func GetTopicComments(topicId int64) ([]Comment, error) {
	var comments []Comment
	if err := connection.Default.Model(&Comment{}).Where("topic_id = ?", topicId).Where("status = ?", COMMENT_STATUS_REVIEWED).Preload("User").Find(&comments).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}
	return comments, nil
}