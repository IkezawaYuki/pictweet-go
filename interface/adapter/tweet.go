package adapter

import (
	"github.com/IkezawaYuki/pictweet-go/domain/model"
	"github.com/IkezawaYuki/pictweet-go/domain/repository"
)

type tweetAdapter struct {
	handler SQLHandler
}

func NewTweetRepository(h SQLHandler) repository.TweetRepository {
	return &tweetAdapter{handler: h}
}

func (t *tweetAdapter) FindByID(id int) (*model.Tweet, error) {
	var tweet model.Tweet
	if err := t.handler.Find(&tweet, id); err != nil {
		return nil, err
	}
	return &tweet, nil
}

func (t *tweetAdapter) FindAll() (*model.Tweets, error) {
	var tweets model.Tweets
	if err := t.handler.Find(&tweets); err != nil {
		return nil, err
	}
	return &tweets, nil
}
func (t *tweetAdapter) Create(tweet *model.Tweet) error {
	if err := t.handler.Create(tweet); err != nil {
		return err
	}
	return nil
}

func (t *tweetAdapter) Update(tweet *model.Tweet) error {
	if err := t.handler.Save(tweet); err != nil {
		return err
	}
	return nil
}
func (t *tweetAdapter) Delete(tweet *model.Tweet) error {
	if err := t.handler.Delete(tweet); err != nil {
		return err
	}
	return nil
}
