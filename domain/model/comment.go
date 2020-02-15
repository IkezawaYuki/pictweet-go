package model

import "time"

type Comment struct {
	ID        uint `gorm:"primary_key"`
	UserID    uint
	TweetID   uint
	Text      string
	CreatedAt time.Time  `json:"-"`
	UpdateAt  time.Time  `json:"-"`
	DeleteAt  *time.Time `sql:"index" json:"-"`

	Tweet Tweet
}

type Comments []Comment
