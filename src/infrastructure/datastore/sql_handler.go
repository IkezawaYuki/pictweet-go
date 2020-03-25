package datastore

import (
	"github.com/IkezawaYuki/pictweet-go/src/domain/model"
	"github.com/IkezawaYuki/pictweet-go/src/domain/repository"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"os"
)

type SqlHandler struct {
	conn *gorm.DB
}

func NewSqlHandler() repository.TweetRepository {
	conn, err := Connect()
	if err != nil {
		panic(err.Error)
	}
	sqlHandler := new(SqlHandler)
	sqlHandler.conn = conn
	return sqlHandler
}

func Connect() (db *gorm.DB, err error) {
	err = godotenv.Load()
	if err != nil {
		logrus.Fatal("Error loading .env file")
	}

	db, err = gorm.Open("mysql",
		os.Getenv("DB_USERNAME")+":"+
			os.Getenv("DB_PASSWORD")+"@tcp("+
			os.Getenv("DB_HOST")+":"+
			os.Getenv("DB_PORT")+")/"+
			os.Getenv("DB_DATABASE")+"?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		logrus.Fatal(err)
	}

	return db, err
}

func (s *SqlHandler) FindTweetByIDWithComment(id uint) (*model.Tweet, error) {
	result := model.Tweet{}
	db := s.conn.Find(&result, id).
		Related(&result.Comment).
		Related(&result.User)

	if db.RecordNotFound() {
		return nil, nil
	}
	s.conn.Where("tweet_id = ?", id).Preload("User").Find(&result.Comment)
	return &result, db.Error
}

func (s *SqlHandler) FindTweetAll() (*model.Tweets, error) {
	result := model.Tweets{}
	db := s.conn.Preload("User").Find(&result)
	if db.RecordNotFound() {
		return nil, nil
	}
	return &result, db.Error
}

func (s *SqlHandler) FindUserByID(id uint) (*model.User, error) {
	result := model.User{}
	db := s.conn.Find(&result, id)
	if db.RecordNotFound() {
		return nil, nil
	}
	return &result, db.Error
}

func (s *SqlHandler) CreateTweet(tweet *model.Tweet) (int, error) {
	db := s.conn.Create(&tweet)
	if db.Error != nil {
		return -1, db.Error
	}
	return int(tweet.ID), nil
}

func (s *SqlHandler) CreateComment(comment *model.Comment) (int, error) {
	db := s.conn.Create(&comment)
	if db.Error != nil {
		return -1, db.Error
	}
	return int(comment.ID), nil
}

func (s *SqlHandler) DeleteTweet(id uint) error {
	if err := s.conn.Where("id = ?", id).Delete(&model.Tweet{}).Error; err != nil {
		return err
	}
	if err := s.conn.Where("tweet_id = ?", id).Delete(&model.Comment{}).Error; err != nil {
		return err
	}
	return nil
}

func (s *SqlHandler) CreateUser(user *model.User) (int, error) {
	db := s.conn.Create(&user)
	if db.Error != nil {
		return -1, db.Error
	}
	return int(user.ID), nil
}

func (s *SqlHandler) FindFavoriteByEmail(email string) (*model.Tweets, error) {
	return nil, nil
}
