package adapter

import (
	"github.com/IkezawaYuki/pictweet-go/domain/model"
	"github.com/IkezawaYuki/pictweet-go/domain/repository"
)

type userAdapter struct {
	handler SQLHandler
}

func NewUserAdapter(h SQLHandler) repository.UserRepository {
	return &userAdapter{handler: h}
}

func (u userAdapter) FindByID(int) (*model.User, error) {
	panic("implement me")
}

func (u userAdapter) FindAll() (*model.Users, error) {
	panic("implement me")
}

func (u userAdapter) Create(*model.User) error {
	panic("implement me")
}

func (u userAdapter) Update(*model.User) error {
	panic("implement me")
}

func (u userAdapter) Delete(*model.User) error {
	panic("implement me")
}
