package main

import (
	"github.com/IkezawaYuki/pictweet-go/src/domain/model"
	"github.com/IkezawaYuki/pictweet-go/src/infrastructure/datastore"
)

func main() {
	db, err := datastore.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.Debug().AutoMigrate(&model.User{})
	db.Debug().AutoMigrate(&model.Tweet{})
	db.Debug().AutoMigrate(&model.Comment{})
}
