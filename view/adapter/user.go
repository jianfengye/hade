package adapter

import (
	"hade/model"
	"hade/view/swagger/models"
)

// 转换成为UserSummary结构
func ToUserSummary(user model.User) *models.UserSummary {
	t := "http://aaa"
	return &models.UserSummary{
		ID:         &user.ID,
		Name:       &user.Name,
		SmallImage: &t,
	}
}