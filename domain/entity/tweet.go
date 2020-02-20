package entity

import (
	"github.com/IkezawaYuki/pictweet-go/domain/dto"
	"time"
)

type Tweet struct {
	ID       uint      `gorm:"primary_key" json:"id"`
	Author   string    `json:"author"`
	Avator   string    `json:"avator"`
	Image    string    `json:"image"`
	Title    string    `json:"title"`
	Text     string    `json:"text"`
	PostDate time.Time `json:"post_date"`
}

type Tweets []Tweet

func NewTweetFactoryByDtos(tweetDtos *dto.TweetsDto, userDtos *dto.UsersDto) *Tweets {
	var tweets Tweets
	for _, t := range *tweetDtos {
		for _, u := range *userDtos {
			if t.UserID == u.ID {
				tweet := Tweet{
					ID:       t.ID,
					Author:   u.Name,
					Avator:   u.Avatar,
					Image:    t.Image,
					Title:    t.Title,
					Text:     t.Text,
					PostDate: time.Time{},
				}
				tweets = append(tweets, tweet)
				continue
			}
		}
	}
	return &tweets
}
