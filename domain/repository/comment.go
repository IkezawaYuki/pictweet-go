package repository

import "github.com/IkezawaYuki/pictweet-go/domain/dto"

type CommentRepository interface {
	FindByID(uint) (*dto.CommentDto, error)
	FindByTweetID(uint) (*dto.CommentsDto, error)
	FindAll() (*dto.CommentsDto, error)
	Create(*dto.CommentDto) error
	Update(*dto.CommentDto) error
	Delete(*dto.CommentDto) error
}
