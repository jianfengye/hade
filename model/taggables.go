package model

import "time"

type Taggable struct {
	CreatedAt    time.Time `gorm:"column:created_at" json:"created_at"`
	ID           int64     `gorm:"column:id;primary_key" json:"id;primary_key"`
	TagID        int64     `gorm:"column:tag_id" json:"tag_id"`
	TaggableID   int64     `gorm:"column:taggable_id" json:"taggable_id"`
	TaggableType string    `gorm:"column:taggable_type" json:"taggable_type"`
	UpdatedAt    time.Time `gorm:"column:updated_at" json:"updated_at"`
}

// TableName sets the insert table name for this struct type
func (t *Taggable) TableName() string {
	return "taggables"
}