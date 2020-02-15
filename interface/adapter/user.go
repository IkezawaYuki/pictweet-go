package adapter

import (
	"github.com/IkezawaYuki/pictweet-go/domain/model"
	"github.com/IkezawaYuki/pictweet-go/domain/repository"
)

type userAdapter struct {
	handler SQLHandler
}

func NewUserAdapter(h SQLHandler) repository.CommentRepository {
	return &userAdapter{handler: h}
}

func (u userAdapter) FindByID(int) (*model.Comment, error) {
	panic("implement me")
}

func (u userAdapter) FindByTweetID(int) (*model.Comments, error) {
	panic("implement me")
}

func (u userAdapter) FindAll() (*model.Comments, error) {
	panic("implement me")
}

func (u userAdapter) Create(*model.Comment) error {
	panic("implement me")
}

func (u userAdapter) Update(*model.Comment) error {
	panic("implement me")
}

func (u userAdapter) Delete(*model.Comment) error {
	panic("implement me")
}
