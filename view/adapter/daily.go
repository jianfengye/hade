package adapter

import (
	"hade/model"
	"hade/view/swagger/models"
)

// 将model.daily转换为输出
func ToDaily(daily *model.Daily, content []model.Content) models.Daily {
	ret := models.Daily{
		Author:    &daily.Author,
		CreatedAt: daily.CreatedAt.String(),
		Day:       &daily.Day,
		ID:        &daily.ID,
		Title:     &daily.Title,
	}

	ret.Content = make([]*models.DailyContentItems, len(content))
	for i, item := range content {
		ret.Content[i] = &models.DailyContentItems{
			Comment: item.Comment,
			Link:    item.Link,
			Title:   item.Title,
		}
	}
	return ret
}