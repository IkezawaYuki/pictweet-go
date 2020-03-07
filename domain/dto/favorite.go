package dto

import "time"

type FavoriteDto struct {
	ID        uint
	UserID    uint
	TweetID   uint
	CreatedAt time.Time  `json:"created_at"`
	UpdateAt  time.Time  `json:"updated_at"`
	DeleteAt  *time.Time `sql:"index" json:"-"`
}

type FavoritesDto []FavoriteDto

func (FavoriteDto) TableName() string {
	return "favorites"
}
