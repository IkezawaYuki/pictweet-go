package main

import (
	"github.com/IkezawaYuki/pictweet-go/domain/dto"
	"github.com/IkezawaYuki/pictweet-go/infrastructure"
)

func main() {
	db, err := infrastructure.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.Debug().AutoMigrate(&dto.UserDto{})
	db.Debug().AutoMigrate(&dto.TweetDto{})
	db.Debug().AutoMigrate(&dto.CommentDto{})
	db.Debug().AutoMigrate(&dto.FavoriteDto{})
}
