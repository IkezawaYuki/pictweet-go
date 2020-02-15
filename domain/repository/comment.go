package repository

import "github.com/IkezawaYuki/pictweet-go/domain/model"

type CommentRepository interface {
	FindByID(int) (*model.Comment, error)
	FindByTweetID(int) (*model.Comments, error)
	FindAll() (*model.Comments, error)
	Create(*model.Comment) error
	Update(*model.Comment) error
	Delete(*model.Comment) error
}
