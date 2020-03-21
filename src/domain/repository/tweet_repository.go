package repository

import "github.com/IkezawaYuki/pictweet-go/src/domain/model"

type TweetRepository interface {
	FindTweetByIDWithComment(uint) (*model.Tweet, error)
	FindTweetAll() (*model.Tweets, error)
	FindUserByID(uint) (*model.User, error)
	AddTweet(*model.Tweet) error
	AddComment(*model.Comment) error
}
