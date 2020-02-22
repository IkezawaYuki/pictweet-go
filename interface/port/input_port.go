package port

import (
	"github.com/IkezawaYuki/pictweet-go/domain/dto"
	"github.com/IkezawaYuki/pictweet-go/domain/entity"
)

type InputPort interface {
	Index() (*entity.Tweets, error)
	CreateTweet(*dto.TweetDto) error
	ShowTweet(uint) (*entity.TweetDetail, error)
	AddComment(*dto.CommentDto) error
}
