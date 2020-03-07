package adapter

import (
	"fmt"
	"github.com/IkezawaYuki/pictweet-go/domain/dto"
	"github.com/IkezawaYuki/pictweet-go/domain/repository"
)

type favoriteRepository struct {
	handler SQLHandler
}

func NewFavoriteRepository(h SQLHandler) repository.FavoriteRepository {
	return &favoriteRepository{handler: h}
}

func (t *favoriteRepository) Toggle(userID uint, tweetID uint) (bool, error) {
	var favorite dto.FavoriteDto
	if err := t.handler.Where(&favorite, "user_id = ? and tweet_id = ?", userID, tweetID); err != nil {
		return false, err
	}
	if &favorite != nil {
		fmt.Println("nil")
		favorite.UserID = userID
		favorite.TweetID = tweetID
		if err := t.handler.Create(&favorite); err != nil {
			return false, err
		}
		return true, nil
	} else {
		if err := t.handler.Delete(&favorite); err != nil {
			return true, err
		}
		return false, nil
	}
}
