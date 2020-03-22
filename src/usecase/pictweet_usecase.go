package usecase

import (
	"github.com/IkezawaYuki/pictweet-go/src/domain/model"
	"github.com/IkezawaYuki/pictweet-go/src/domain/repository"
)

type pictweetUsecase struct {
	tweetRepository repository.TweetRepository
}

type PictweetUsecase interface {
	PostTweet(*model.Tweet) (int, error)
	PostComment(*model.Comment) (int, error)
	ListTweets() (*model.Tweets, error)
	ShowTweet(uint) (*model.Tweet, error)
}

func NewPictweetUsecase(tweetRepo repository.TweetRepository) PictweetUsecase {
	return &pictweetUsecase{
		tweetRepository: tweetRepo,
	}
}

func (p *pictweetUsecase) PostTweet(tweet *model.Tweet) (int, error) {
	id, err := p.tweetRepository.AddTweet(tweet)
	if err != nil {
		return -1, err
	}
	return id, nil
}

func (p *pictweetUsecase) PostComment(comment *model.Comment) (int, error) {
	id, err := p.tweetRepository.AddComment(comment)
	if err != nil {
		return -1, err
	}
	return id, nil
}

func (p *pictweetUsecase) ListTweets() (*model.Tweets, error) {
	tweets, err := p.tweetRepository.FindTweetAll()
	if err != nil {
		return nil, err
	}
	return tweets, err
}

func (p *pictweetUsecase) ShowTweet(id uint) (*model.Tweet, error) {
	tweet, err := p.tweetRepository.FindTweetByIDWithComment(id)
	if err != nil {
		return nil, err
	}
	return tweet, nil
}
