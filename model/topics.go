package model

import (
	"hade/bootstrap/connection"
	"github.com/jianfengye/collection"
	"github.com/jinzhu/gorm"
	"time"
)

const (
	TOPIC_STATUS_REVIEWED   = "reviewed"
	TOPIC_STATUS_UNREVIEWED = "unreviewed"
	TOPIC_STATUS_DELETED    = "deleted"
)

type Topic struct {
	Category  string    `gorm:"column:category" json:"category"`
	Content   string    `gorm:"column:content" json:"content"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	ID        int64     `gorm:"column:id;primary_key" json:"id;primary_key"`
	LikeCount int64     `gorm:"column:like_count" json:"like_count"`
	CommentCount int64     `gorm:"column:comment_count" json:"comment_count"`
	Link      string    `gorm:"column:link" json:"link"`
	Score     int64     `gorm:"column:score" json:"score"`
	Source    string    `gorm:"column:source" json:"source"`
	Status    string    `gorm:"column:status" json:"status"`
	Title     string    `gorm:"column:title" json:"title"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
	UserID    int64     `gorm:"column:user_id" json:"user_id"`
	User      User   `gorm:"ForeignKey:user_id"`
}

// TableName sets the insert table name for this struct type
func (t *Topic) TableName() string {
	return "topics"
}

func GetTopicsByTagPager(tagId, offset int64, size int64) ([]Topic, int64, error) {
	var sum int64
	var topics []Topic
	var tag Tag
	if err := connection.Default.Model(&Tag{}).First(&tag, tagId).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, 0, nil
		}
		return nil, 0, err
	}

	var taggables []Taggable
	if err := connection.Default.Model(&Taggable{}).Count(&sum).Where("taggable_type = ?", TaggableType_Topic).Where("tag_id = ?", tagId).Order("updated_at desc").Offset(offset).Limit(size).Find(&taggables).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, 0, nil
		}
		return nil, 0, err
	}

	taggableColl := collection.NewObjCollection(taggables)
	topicIds, err := taggableColl.Pluck("TaggableID").ToInt64s()
	if err != nil {
		return nil, sum, err
	}

	if err := connection.Default.Model(&Topic{}).Where("id in (?)", topicIds).Where("status = ?", TOPIC_STATUS_REVIEWED).Find(&topics).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, sum, nil
		}
		return nil, sum, err
	}
	return topics, sum, nil

}

// 分页获取所有的topic列表
func GetTopicsByPager(offset int64, size int64) ([]Topic, int64, error){
	var sum int64
	var topics []Topic
	err := connection.Default.Model(&Topic{}).Count(&sum).Where("status = ?", "reviewed").Offset(offset).Limit(size).Preload("User").Find(&topics).Error
	if err == nil {
		return topics, sum, nil
	}

	if gorm.IsRecordNotFoundError(err) {
		return nil, sum, nil
	}

	return nil, sum, nil
}


// 获取一个topic
func GetTopicWithUser(id int64) *Topic {
	var topic Topic
	err := connection.Default.Model(&Topic{}).Where("status = ?", "reviewed").Preload("User").First(&topic, id).Error
	if err == nil {
		return &topic
	}

	return nil
}

// 获取一个topic
func GetTopic(id int64) *Topic {
	var topic Topic
	err := connection.Default.Model(&Topic{}).Where("status = ?", "reviewed").First(&topic, id).Error
	if err == nil {
		return &topic
	}

	return nil
}

// 删除一个话题
func DeleteTopic(id int64) error {
	if err := connection.Default.Model(&Topic{}).Set("status", TOPIC_STATUS_DELETED).Where("id = ?", id).Error; err != nil {
		return err
	}
	return nil
}



