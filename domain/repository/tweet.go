package repository

import "github.com/IkezawaYuki/pictweet-go/domain/model"

type TweetRepository interface {
	FindByID(int) (*model.Tweet, error)
	FindAll() (*model.Tweets, error)
	Create(*model.Tweet) error
	Update(*model.Tweet) error
	Delete(*model.Tweet) error
}
