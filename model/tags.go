package model

import (
	"hade/bootstrap/connection"
	"errors"
	"github.com/jianfengye/collection"
	"github.com/jinzhu/gorm"
	"time"
)

const TaggableType_Topic = "topic"

type Tag struct {
	CreatedAt  time.Time `gorm:"column:created_at" json:"created_at"`
	ID         int64     `gorm:"column:id;primary_key" json:"id;primary_key"`
	Name       string    `gorm:"column:name" json:"name"`
	Status     string    `gorm:"column:status" json:"status"`
	TopicCount int64     `gorm:"column:topic_count" json:"topic_count"`
	UpdatedAt  time.Time `gorm:"column:updated_at" json:"updated_at"`
}

// TableName sets the insert table name for this struct type
func (t *Tag) TableName() string {
	return "tags"
}

func GetAllTags() ([]Tag, error) {
	var tags []Tag
	if err := connection.Default.Find(&tags).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}
	return tags, nil
}

// 获取话题的标签
func GetTopicTags(topicId int64) ([]Tag, error) {
	var taggables []Taggable
	if err := connection.Default.Model(&Taggable{}).Where("taggable_id = ?", topicId).Where("taggable_type = ?", TaggableType_Topic).Find(&taggables).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}

	taggableColl := collection.NewObjCollection(taggables)
	tagIds, err := taggableColl.Pluck("TagID").ToInt64s()
	if err != nil {
		return nil, err
	}

	var tags []Tag
	if err := connection.Default.Model(&Tag{}).Where("id in (?)", tagIds).Find(&tags).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}
	return tags, nil
}

// 更新一个话题的所有标签, 这里使用全删全加的方式
func UpdateTopicTags(topicId int64, tagIds []int64) error {
	// 确保所有的标签都是存在的
	var tags []Tag
	if err := connection.Default.Model(&Tag{}).Where("id in (?)", tagIds).Find(&tags).Error; err != nil {
		return err
	}

	if len(tags) != len(tagIds) {
		return errors.New("标签传递错误")
	}

	var toAddTagIds []int64

	// 获取这个话题已经有的标签
	if err := connection.Default.Model(&Taggable{}).Where("taggable_id = ?", topicId).Where("taggable_type = ?", TaggableType_Topic).Delete(&Taggable{}).Error; err != nil {
		return err
	}

	for _, toAddTagId := range toAddTagIds {
		toAdd := Taggable{
			TagID:        toAddTagId,
			TaggableID:   topicId,
			TaggableType: TaggableType_Topic,
		}
		if err := connection.Default.Model(&Taggable{}).Save(toAdd).Error; err != nil {
			return err
		}
	}

	return nil
}

// 把多个标签放到一个话题
func AddTagsToTopic(topicId int64, tagIds []int64) error {
	// 确保所有的标签都是存在的
	var tags []Tag
	if err := connection.Default.Model(&Tag{}).Where("id in (?)", tagIds).Find(&tags).Error; err != nil {
		return err
	}

	if len(tags) != len(tagIds) {
		return errors.New("标签传递错误")
	}

	var toAddTagIds []int64

	// 获取这个话题已经有的标签
	var taggables []Taggable
	if err := connection.Default.Model(&Taggable{}).Where("taggable_id = ?", topicId).Where("taggable_type = ?", TaggableType_Topic).Find(&taggables).Error; err != nil {
		if !gorm.IsRecordNotFoundError(err) {
			return err
		}
	}

	taggableColl := collection.NewObjCollection(taggables)
	tagIdColl := collection.NewInt64Collection(tagIds)

	toAddTagIds, err := tagIdColl.Diff(taggableColl.Pluck("TaggableID").Unique()).ToInt64s()
	if err != nil {
		return err
	}

	if len(toAddTagIds) == 0 {
		return nil
	}

	for _, toAddTagId := range toAddTagIds {
		// TODO 这里可以优化为批量插入
		toAdd := Taggable{
			TagID: toAddTagId,
			TaggableType: TaggableType_Topic,
			TaggableID: topicId,
		}

		if err := connection.Default.Model(&Taggable{}).Save(&toAdd).Error; err != nil {
			return err
		}
	}

	return nil
}