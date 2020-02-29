package repository

import "github.com/IkezawaYuki/pictweet-go/domain/dto"

type TweetRepository interface {
	FindByID(uint) (*dto.TweetDto, error)
	FindAll() (*dto.TweetsDto, error)
	Create(*dto.TweetDto) (uint, error)
	Update(*dto.TweetDto) (uint, error)
	Delete(*dto.TweetDto) error
}
