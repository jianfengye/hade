package suite1

import (
	"backend/bootstrap/connection"
	"backend/model"
)

var (
	user1 model.User
	user2 model.User

	tag1 model.Tag
	tag2 model.Tag
	tag3 model.Tag

)

func initDB() {
	user1 = model.User{
		Email:     "tester1@huoding.com",
		ID:        1,
		Name:      "tester1",
		Password:  "",
	}
	user2 = model.User{
		Email:     "tester2@huoding.com",
		ID:        2,
		Name:      "tester2",
		Password:  "",
	}
	tag1 = model.Tag{
		ID:         1,
		Name:       "Golang",
		Status:     "reviewed",
		TopicCount: 0,
	}
	tag2 = model.Tag{
		ID:         2,
		Name:       "PHP",
		Status:     "reviewed",
		TopicCount: 0,
	}
	tag3 = model.Tag{
		ID:         3,
		Name:       "JavaScript",
		Status:     "reviewed",
		TopicCount: 0,
	}
	connection.Default.Create(user1)
	connection.Default.Create(user2)
	connection.Default.Create(tag1)
	connection.Default.Create(tag2)
	connection.Default.Create(tag3)
}

