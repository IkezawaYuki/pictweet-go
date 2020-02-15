package main

import (
	"github.com/IkezawaYuki/pictweet-go/domain/model"
	"github.com/IkezawaYuki/pictweet-go/infrastructure"
)

func main() {
	db, err := infrastructure.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.Debug().AutoMigrate(&model.User{})
	db.Debug().AutoMigrate(&model.Tweet{})
	db.Debug().AutoMigrate(&model.Comment{})

}
