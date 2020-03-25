package model

import "time"

// Favorite お気に入り
type Favorite struct {
	ID        uint `gorm:"primary_key"`
	UserID    uint
	TweetID   uint
	CreatedAt time.Time
	UpdateAt  time.Time
	DeleteAt  *time.Time
}

type Favorites []Favorite
