package model

import (
	"hade/bootstrap/connection"
	"encoding/json"
	"github.com/jinzhu/gorm"
	"time"
)

type Daily struct {
	CreatedAt   time.Time `gorm:"column:created_at" json:"created_at"`
	ID          int64     `gorm:"column:id;primary_key" json:"id;primary_key"`
	Title string    `gorm:"column:title" json:"title"`
	Author string    `gorm:"column:author" json:"author"`
	Day string    `gorm:"column:day" json:"day"`
	Content string    `gorm:"column:content" json:"content"`
	UpdatedAt   time.Time `gorm:"column:updated_at" json:"updated_at"`
}

// Content内的json结构
type Content struct {
	Title string `json:"title"`
	Link string `json:"link"`
	Comment string `json:"comment"`
}

// TableName sets the insert table name for this struct type
func (d *Daily) TableName() string {
	return "daily"
}

// 获取每日情况
func GetDaily(day string) (*Daily, error) {
	daily := Daily{}
	if err := connection.Default.Model(&Daily{}).First(&daily, "day=?", day).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}
	return &daily, nil
}

// 解析content中的信息
func (d *Daily) ParseContent() ([]Content, error) {
	c := []Content{}
	err := json.Unmarshal([]byte(d.Content), &c)
	if err != nil {
		return c, err
	}
	return c, nil
}