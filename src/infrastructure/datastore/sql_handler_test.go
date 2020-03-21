package datastore

import (
	"fmt"
	"github.com/IkezawaYuki/pictweet-go/src/domain/model"
	"testing"
	"time"
)

func TestSqlHandler_FindTweetByIDWithComment(t *testing.T) {
	handler := NewSqlHandler()
	tweet, err := handler.FindTweetByIDWithComment(uint(1))
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(tweet)
	fmt.Println(tweet.Comment)
	fmt.Println(tweet.User)
	fmt.Println(tweet.Comment[0].User)
	fmt.Println(tweet.Comment[1].User)
	fmt.Println(tweet.Comment[2].User)
}

func TestSqlHandler_FindTweetAll(t *testing.T) {
	handler := NewSqlHandler()
	tweets, err := handler.FindTweetAll()
	if err != nil {
		t.Error(err)
	}
	for _, t := range *tweets {
		fmt.Println(t)
	}
}

func TestSqlHandler_FindUserByID(t *testing.T) {
	handler := NewSqlHandler()
	user, err := handler.FindUserByID(uint(2))
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(user)
}

func TestSqlHandler_AddTweet(t *testing.T) {
	handler := NewSqlHandler()
	tweet := &model.Tweet{
		ID:        5,
		UserID:    2,
		Image:     "https://diamond.jp/quarterly/mwimgs/a/1/350/img_a12a2d5b77d4861feb24aa14187b6c39670979.jpg",
		Title:     "neko",
		Text:      "あいうえお",
		CreatedAt: time.Now(),
	}
	err := handler.AddTweet(tweet)
	if err != nil {
		t.Error(err)
	}
}
