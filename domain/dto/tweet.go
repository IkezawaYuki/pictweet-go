package dto

import (
	"strconv"
	"time"
)

type TweetDto struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	UserID    uint       `json:"user_id"`
	Image     string     `json:"image"`
	Title     string     `json:"title"`
	Text      string     `json:"text"`
	CreatedAt time.Time  `json:"created_at"`
	UpdateAt  time.Time  `json:"updated_at"`
	DeleteAt  *time.Time `sql:"index" json:"-"`
}

type TweetsDto []TweetDto

func (TweetDto) TableName() string {
	return "tweets"
}

func NewTweetDto(userID, image, title, text string) (*TweetDto, error) {
	userId, err := strconv.ParseUint(userID, 10, 64)
	if err != nil {
		return nil, err
	}
	return &TweetDto{
		UserID: uint(userId),
		Image:  image,
		Title:  title,
		Text:   text,
	}, nil
}
