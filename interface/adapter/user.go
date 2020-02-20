package adapter

import (
	"github.com/IkezawaYuki/pictweet-go/domain/dto"
	"github.com/IkezawaYuki/pictweet-go/domain/repository"
)

type userRepository struct {
	handler SQLHandler
}

func NewUserAdapter(h SQLHandler) repository.UserRepository {
	return &userRepository{handler: h}
}

func (u *userRepository) FindByID(id uint) (*dto.UserDto, error) {
	var user dto.UserDto
	if err := u.handler.Where(&user, "id = ?", id); err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *userRepository) FindAll() (*dto.UsersDto, error) {
	var users dto.UsersDto
	if err := u.handler.Find(&users); err != nil {
		return nil, err
	}
	return &users, nil
}

func (u *userRepository) Create(user *dto.UserDto) error {
	if err := u.handler.Create(&user); err != nil {
		return err
	}
	return nil
}

func (u *userRepository) Update(user *dto.UserDto) error {
	if err := u.handler.Create(&user); err != nil {
		return err
	}
	return nil
}

func (u *userRepository) Delete(user *dto.UserDto) error {
	if err := u.handler.Create(&user); err != nil {
		return err
	}
	return nil
}

func (u *userRepository) FindInUserID(userIDs []uint) (*dto.UsersDto, error) {
	var users dto.UsersDto
	if err := u.handler.In(&users, "id in (?)", userIDs); err != nil {
		return nil, err
	}
	return &users, nil
}
