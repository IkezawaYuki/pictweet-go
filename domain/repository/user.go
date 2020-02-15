package repository

import "github.com/IkezawaYuki/pictweet-go/domain/model"

type UserRepository interface {
	FindByID(int) (*model.User, error)
	FindAll() (*model.Users, error)
	Create(*model.User) error
	Update(*model.User) error
	Delete(*model.User) error
}
