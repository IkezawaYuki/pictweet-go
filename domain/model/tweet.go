package model

import "time"

type Tweet struct {
	ID        uint `gorm:"primary_key"`
	UserID    uint
	Image     string
	Title     string
	Text      string
	CreatedAt time.Time  `json:"-"`
	UpdateAt  time.Time  `json:"-"`
	DeleteAt  *time.Time `sql:"index" json:"-"`
}

type Tweets []Tweet
