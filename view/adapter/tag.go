package adapter

import (
	"backend/model"
	"backend/view/swagger/models"
)

func ToTag(tag model.Tag) *models.Tag {
	out := models.Tag{
		ID:         &tag.ID,
		Name:       &tag.Name,
		TopicCount: &tag.TopicCount,
	}
	return &out
}


func ToTags(tags []model.Tag) []*models.Tag {
	out := make([]*models.Tag, len(tags))
	for i, tag := range tags {
		out[i] = ToTag(tag)
	}
	return out
}