package dto

import "time"

type CommentDto struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	UserID    uint       `json:"user_id"`
	TweetID   uint       `json:"tweet_id"`
	Text      string     `json:"text"`
	CreatedAt time.Time  `json:"-"`
	UpdateAt  time.Time  `json:"update_at"`
	DeleteAt  *time.Time `sql:"index" json:"-"`
}

type CommentsDto []CommentDto

func (CommentDto) TableName() string {
	return "comments"
}
