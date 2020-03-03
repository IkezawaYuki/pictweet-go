package port

import (
	"github.com/IkezawaYuki/pictweet-go/domain/dto"
	"github.com/IkezawaYuki/pictweet-go/domain/entity"
)

type InputPort interface {
	Index() (*entity.Tweets, error)
	FindByID(id uint) (*entity.Tweet, error)
	CreateTweet(*dto.TweetDto) (uint, error)
	ShowTweet(uint) (*entity.TweetDetail, error)
	AddComment(*dto.CommentDto) error
	CreateUser(*dto.UserDto) error
}
