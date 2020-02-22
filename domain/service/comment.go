package service

import (
	"github.com/IkezawaYuki/pictweet-go/domain/dto"
	"github.com/IkezawaYuki/pictweet-go/domain/entity"
)

type CommentService struct {
}

func (s *CommentService) NewCommentByDtos(commentDtos *dto.CommentsDto, userDtos *dto.UsersDto) *entity.Comments {
	var comments entity.Comments
	for _, t := range *commentDtos {
		for _, u := range *userDtos {
			if t.UserID == u.ID {
				comment := entity.Comment{
					ID:       t.ID,
					Author:   u.Name,
					Avatar:   u.Avatar,
					Text:     t.Text,
					PostDate: t.CreatedAt.Format("2006/01/02 03:04"),
				}
				comments = append(comments, comment)
				continue
			}
		}
	}
	return &comments
}
