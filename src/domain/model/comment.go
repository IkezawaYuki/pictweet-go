package model

import "time"

// Comment コメント
type Comment struct {
	ID        uint `gorm:"primary_key"`
	UserID    uint
	User      User
	TweetID   uint
	Text      string
	CreatedAt time.Time
	UpdateAt  time.Time
	DeleteAt  *time.Time `sql:"index"`
}

type Comments []Comment
