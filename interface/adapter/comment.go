package adapter

import (
	"github.com/IkezawaYuki/pictweet-go/domain/model"
	"github.com/IkezawaYuki/pictweet-go/domain/repository"
)

type commentAdapter struct {
	handler SQLHandler
}

func NewCommentAdapter(h SQLHandler) repository.CommentRepository {
	return &commentAdapter{handler: h}
}

func (c commentAdapter) FindByID(int) (*model.Comment, error) {
	panic("implement me")
}

func (c commentAdapter) FindByTweetID(int) (*model.Comments, error) {
	panic("implement me")
}

func (c commentAdapter) FindAll() (*model.Comments, error) {
	panic("implement me")
}

func (c commentAdapter) Create(*model.Comment) error {
	panic("implement me")
}

func (c commentAdapter) Update(*model.Comment) error {
	panic("implement me")
}

func (c commentAdapter) Delete(*model.Comment) error {
	panic("implement me")
}
