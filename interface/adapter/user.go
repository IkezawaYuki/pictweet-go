package adapter

import (
	"github.com/IkezawaYuki/pictweet-go/domain/model"
	"github.com/IkezawaYuki/pictweet-go/domain/repository"
)

type userRepository struct {
	handler SQLHandler
}

func NewUserAdapter(h SQLHandler) repository.UserRepository {
	return &userRepository{handler: h}
}

func (u *userRepository) FindByID(id int) (*model.User, error) {
	var user model.User
	if err := u.handler.Where(&user, "id = ?", id); err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *userRepository) FindAll() (*model.Users, error) {
	var users model.Users
	if err := u.handler.Find(&users); err != nil {
		return nil, err
	}
	return &users, nil
}

func (u *userRepository) Create(user *model.User) error {
	if err := u.handler.Create(&user); err != nil {
		return err
	}
	return nil
}

func (u *userRepository) Update(user *model.User) error {
	if err := u.handler.Create(&user); err != nil {
		return err
	}
	return nil
}

func (u *userRepository) Delete(user *model.User) error {
	if err := u.handler.Create(&user); err != nil {
		return err
	}
	return nil
}
