package repository

import "github.com/IkezawaYuki/pictweet-go/domain/dto"

type UserRepository interface {
	FindByID(uint) (*dto.UserDto, error)
	FindByEmail(string) (*dto.UserDto, error)
	FindAll() (*dto.UsersDto, error)
	Create(*dto.UserDto) error
	Update(*dto.UserDto) error
	Delete(*dto.UserDto) error
	FindInUserID([]uint) (*dto.UsersDto, error)
}
