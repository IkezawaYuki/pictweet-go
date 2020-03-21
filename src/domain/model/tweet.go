package model

import "time"

// Tweet ツイート
type Tweet struct {
	ID        uint `gorm:"primary_key"`
	Comment   Comments
	UserID    uint
	User      User
	Image     string
	Title     string
	Text      string
	CreatedAt time.Time
	UpdateAt  time.Time
	DeleteAt  *time.Time `sql:"index"`
}

type Tweets []Tweet
