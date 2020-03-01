package adapter

import (
	"github.com/IkezawaYuki/pictweet-go/domain/dto"
	"github.com/IkezawaYuki/pictweet-go/domain/repository"
)

type tweetRepository struct {
	handler SQLHandler
}

func NewTweetRepository(h SQLHandler) repository.TweetRepository {
	return &tweetRepository{handler: h}
}

func (t *tweetRepository) FindByID(id uint) (*dto.TweetDto, error) {
	var tweet dto.TweetDto
	if err := t.handler.Where(&tweet, "id = ?", id); err != nil {
		return nil, err
	}
	return &tweet, nil
}

func (t *tweetRepository) FindAll() (*dto.TweetsDto, error) {
	var tweets dto.TweetsDto
	if err := t.handler.Find(&tweets); err != nil {
		return nil, err
	}
	return &tweets, nil
}
func (t *tweetRepository) Create(tweet *dto.TweetDto) (uint, error) {
	if err := t.handler.Create(tweet); err != nil {
		return 0, err
	}
	return tweet.ID, nil
}

func (t *tweetRepository) Update(tweet *dto.TweetDto) (uint, error) {
	if err := t.handler.Save(tweet); err != nil {
		return 0, err
	}
	return tweet.ID, nil
}
func (t *tweetRepository) Delete(tweet *dto.TweetDto) error {
	if err := t.handler.Delete(tweet); err != nil {
		return err
	}
	return nil
}
