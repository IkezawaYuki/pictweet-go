package service

import (
	"github.com/IkezawaYuki/pictweet-go/domain/dto"
	"github.com/IkezawaYuki/pictweet-go/domain/entity"
)

type TweetService struct {
}

func (s *TweetService) NewTweetByDtos(tweetDtos *dto.TweetsDto, userDtos *dto.UsersDto) *entity.Tweets {
	var tweets entity.Tweets
	for _, t := range *tweetDtos {
		for _, u := range *userDtos {
			if t.UserID == u.ID {
				tweet := entity.Tweet{
					ID:       t.ID,
					Author:   u.Name,
					Avatar:   u.Avatar,
					Image:    t.Image,
					Title:    t.Title,
					Text:     t.Text,
					PostDate: t.CreatedAt.Format("2006/01/02 03:04"),
				}
				tweets = append(tweets, tweet)
				continue
			}
		}
	}
	return &tweets
}

func (s *TweetService) NewTweetByDto(tweetDto *dto.TweetDto, userDto *dto.UserDto) *entity.Tweet {
	return &entity.Tweet{
		ID:       tweetDto.ID,
		Author:   userDto.Name,
		Avatar:   userDto.Avatar,
		Image:    tweetDto.Image,
		Title:    tweetDto.Title,
		Text:     tweetDto.Text,
		PostDate: tweetDto.CreatedAt.Format("2006/01/02 03:04"),
	}
}
