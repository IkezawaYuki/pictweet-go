package repository

import "github.com/IkezawaYuki/pictweet-go/src/domain/model"

type TweetRepository interface {
	FindTweetByIDWithComment(uint) (*model.Tweet, error)
	FindTweetAll() (*model.Tweets, error)
	FindUserByID(uint) (*model.User, error)
	CreateTweet(*model.Tweet) (int, error)
	CreateComment(*model.Comment) (int, error)
	DeleteTweet(uint) error
	CreateUser(*model.User) (int, error)
}
