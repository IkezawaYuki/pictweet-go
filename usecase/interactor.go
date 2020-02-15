package usecase

import (
	"github.com/IkezawaYuki/pictweet-go/domain/model"
	"github.com/IkezawaYuki/pictweet-go/domain/repository"
	"github.com/IkezawaYuki/pictweet-go/interface/port"
)

type pictweetInteractor struct {
	TweetRepo   repository.TweetRepository
	CommentRepo repository.CommentRepository
	UserRepo    repository.UserRepository
}

func NewPictweetInteractor(tRepo repository.TweetRepository,
	cRepo repository.CommentRepository,
	uRepo repository.UserRepository) port.InputPort {
	return &pictweetInteractor{
		TweetRepo:   tRepo,
		CommentRepo: cRepo,
		UserRepo:    uRepo,
	}
}

func (i *pictweetInteractor) Index() (*model.Tweets, error) {
	tweets, err := i.TweetRepo.FindAll()
	if err != nil {
		return nil, err
	}
	return tweets, nil
}
