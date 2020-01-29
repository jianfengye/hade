package model

import "time"

type User struct {
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	Email     string    `gorm:"column:email" json:"email"`
	ID        int64     `gorm:"column:id;primary_key" json:"id;primary_key"`
	Name      string    `gorm:"column:name" json:"name"`
	Password  string    `gorm:"column:password" json:"password"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}

// TableName sets the insert table name for this struct type
func (u *User) TableName() string {
	return "users"
}