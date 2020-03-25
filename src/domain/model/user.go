package model

import "time"

// User ユーザー
type User struct {
	ID        uint `gorm:"primary_key"`
	Name      string
	Email     string
	Avatar    string
	Favorites []Favorite
	CreatedAt time.Time
	UpdateAt  time.Time
	DeleteAt  *time.Time `sql:"index"`
}

type Users []User
