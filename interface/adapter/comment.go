package adapter

import (
	"github.com/IkezawaYuki/pictweet-go/domain/model"
	"github.com/IkezawaYuki/pictweet-go/domain/repository"
)

type commentRepository struct {
	handler SQLHandler
}

func NewCommentAdapter(h SQLHandler) repository.CommentRepository {
	return &commentRepository{handler: h}
}

func (c *commentRepository) FindByID(id int) (*model.Comment, error) {
	var comment model.Comment
	if err := c.handler.Find(&comment, id); err != nil {
		return nil, err
	}
	return &comment, nil
}

func (c *commentRepository) FindByTweetID(tweetID int) (*model.Comments, error) {
	var comments model.Comments
	if err := c.handler.Where(&comments, "tweet_id = ?", tweetID); err != nil {
		return nil, err
	}
	return &comments, nil
}

func (c *commentRepository) FindAll() (*model.Comments, error) {
	var comments model.Comments
	if err := c.handler.Find(comments); err != nil {
		return nil, err
	}
	return &comments, nil
}

func (c *commentRepository) Create(comment *model.Comment) error {
	if err := c.handler.Create(&comment); err != nil {
		return err
	}
	return nil
}

func (c *commentRepository) Update(comment *model.Comment) error {
	if err := c.handler.Save(&comment); err != nil {
		return err
	}
	return nil
}

func (c *commentRepository) Delete(comment *model.Comment) error {
	if err := c.handler.Delete(&comment); err != nil {
		return err
	}
	return nil
}
