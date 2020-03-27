package usecase

import (
	"fmt"
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
	DeleteTweet(uint) error
	RegisterUser(*model.User) (int, error)
	GetFavorite(string) (*model.Tweets, error)
}

func NewPictweetUsecase(tweetRepo repository.TweetRepository) PictweetUsecase {
	return &pictweetUsecase{
		tweetRepository: tweetRepo,
	}
}

func (p *pictweetUsecase) PostTweet(tweet *model.Tweet) (int, error) {
	fmt.Println(tweet)
	return p.tweetRepository.CreateTweet(tweet)
}

func (p *pictweetUsecase) PostComment(comment *model.Comment) (int, error) {
	return p.tweetRepository.CreateComment(comment)
}

func (p *pictweetUsecase) ListTweets() (*model.Tweets, error) {
	return p.tweetRepository.FindTweetAll()
}

func (p *pictweetUsecase) ShowTweet(id uint) (*model.Tweet, error) {
	return p.tweetRepository.FindTweetByIDWithComment(id)
}

func (p *pictweetUsecase) DeleteTweet(id uint) error {
	return p.tweetRepository.DeleteTweet(id)
}

func (p *pictweetUsecase) RegisterUser(user *model.User) (int, error) {
	return p.tweetRepository.CreateUser(user)
}

func (p *pictweetUsecase) GetFavorite(email string) (*model.Tweets, error) {
	return p.tweetRepository.FindFavoriteByEmail(email)
}
