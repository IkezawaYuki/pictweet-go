package model

import "time"

type Comment struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	UserID    uint       `json:"user_id"`
	TweetID   uint       `json:"tweet_id"`
	Text      string     `json:"text"`
	CreatedAt time.Time  `json:"-"`
	UpdateAt  time.Time  `json:"update_at"`
	DeleteAt  *time.Time `sql:"index" json:"-"`
}

type Comments []Comment
