package adapter

import (
	"github.com/IkezawaYuki/pictweet-go/domain/model"
	"github.com/IkezawaYuki/pictweet-go/domain/repository"
)

type tweetRepository struct {
	handler SQLHandler
}

func NewTweetRepository(h SQLHandler) repository.TweetRepository {
	return &tweetRepository{handler: h}
}

func (t *tweetRepository) FindByID(id int) (*model.Tweet, error) {
	var tweet model.Tweet
	if err := t.handler.Where(&tweet, "id = ?", id); err != nil {
		return nil, err
	}
	return &tweet, nil
}

func (t *tweetRepository) FindAll() (*model.Tweets, error) {
	var tweets model.Tweets
	if err := t.handler.Find(&tweets); err != nil {
		return nil, err
	}
	return &tweets, nil
}
func (t *tweetRepository) Create(tweet *model.Tweet) error {
	if err := t.handler.Create(tweet); err != nil {
		return err
	}
	return nil
}

func (t *tweetRepository) Update(tweet *model.Tweet) error {
	if err := t.handler.Save(tweet); err != nil {
		return err
	}
	return nil
}
func (t *tweetRepository) Delete(tweet *model.Tweet) error {
	if err := t.handler.Delete(tweet); err != nil {
		return err
	}
	return nil
}
