package dto

import (
	"strconv"
	"time"
)

type CommentDto struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	UserID    uint       `json:"user_id"`
	TweetID   uint       `json:"tweet_id"`
	Text      string     `json:"text"`
	CreatedAt time.Time  `json:"created_at"`
	UpdateAt  time.Time  `json:"updated_at"`
	DeleteAt  *time.Time `sql:"index" json:"-"`
}

type CommentsDto []CommentDto

func (CommentDto) TableName() string {
	return "comments"
}

func NewCommentDto(tweetID, userID, text string) (*CommentDto, error) {
	userId, err := strconv.ParseUint(userID, 10, 64)
	tweetId, err := strconv.ParseUint(tweetID, 10, 64)

	if err != nil {
		return nil, err
	}
	return &CommentDto{
		UserID:  uint(userId),
		TweetID: uint(tweetId),
		Text:    text,
	}, nil
}
