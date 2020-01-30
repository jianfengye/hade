package model

import (
	"time"
)

const (
	TOPIC_STATUS_REVIEWED   = "reviewed"
	TOPIC_STATUS_UNREVIEWED = "unreviewed"
	TOPIC_STATUS_DELETED    = "deleted"
)

type Topic struct {
	Category     string    `gorm:"column:category" json:"category"`
	Content      string    `gorm:"column:content" json:"content"`
	CreatedAt    time.Time `gorm:"column:created_at" json:"created_at"`
	ID           int64     `gorm:"column:id;primary_key" json:"id;primary_key"`
	LikeCount    int64     `gorm:"column:like_count" json:"like_count"`
	CommentCount int64     `gorm:"column:comment_count" json:"comment_count"`
	Link         string    `gorm:"column:link" json:"link"`
	Score        int64     `gorm:"column:score" json:"score"`
	Source       string    `gorm:"column:source" json:"source"`
	Status       string    `gorm:"column:status" json:"status"`
	Title        string    `gorm:"column:title" json:"title"`
	UpdatedAt    time.Time `gorm:"column:updated_at" json:"updated_at"`
	UserID       int64     `gorm:"column:user_id" json:"user_id"`
	User         User      `gorm:"ForeignKey:user_id"`
}

// TableName sets the insert table name for this struct type
func (t *Topic) TableName() string {
	return "topics"
}
