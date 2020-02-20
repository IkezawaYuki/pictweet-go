package dto

import "time"

type TweetDto struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	UserID    uint       `json:"user_id"`
	Image     string     `json:"image"`
	Title     string     `json:"title"`
	Text      string     `json:"text"`
	CreatedAt time.Time  `json:"-"`
	UpdateAt  time.Time  `json:"update_at"`
	DeleteAt  *time.Time `sql:"index" json:"-"`
}

type TweetsDto []TweetDto

func (TweetDto) TableName() string {
	return "tweets"
}
