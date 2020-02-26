package usecase

import (
	"fmt"
	"github.com/IkezawaYuki/pictweet-go/domain/dto"
	"github.com/IkezawaYuki/pictweet-go/domain/entity"
	"github.com/IkezawaYuki/pictweet-go/domain/repository"
	"github.com/IkezawaYuki/pictweet-go/domain/service"
	"github.com/IkezawaYuki/pictweet-go/interface/port"
)

type pictweetInteractor struct {
	TweetRepo      repository.TweetRepository
	CommentRepo    repository.CommentRepository
	UserRepo       repository.UserRepository
	TweetService   service.TweetService
	CommentService service.CommentService
}

func NewPictweetInteractor(tRepo repository.TweetRepository,
	cRepo repository.CommentRepository,
	uRepo repository.UserRepository) port.InputPort {
	return &pictweetInteractor{
		TweetRepo:   tRepo,
		CommentRepo: cRepo,
		UserRepo:    uRepo,
	}
}

// Index ツイート一覧表示ロジック
func (i *pictweetInteractor) Index() (*entity.Tweets, error) {
	tweetsDto, err := i.TweetRepo.FindAll()
	if err != nil {
		return nil, err
	}

	var ids []uint
	for _, tweet := range *tweetsDto {
		ids = append(ids, tweet.ID)
	}
	usersDto, err := i.UserRepo.FindInUserID(ids)
	if err != nil {
		return nil, err
	}

	return i.TweetService.NewTweetByDtos(tweetsDto, usersDto), nil
}

// CreateTweet ツイート新規投稿ロジック
func (i *pictweetInteractor) CreateTweet(dto *dto.TweetDto) error {
	if err := i.TweetRepo.Create(dto); err != nil {
		return err
	}
	return nil
}

func (i *pictweetInteractor) ShowTweet(id uint) (*entity.TweetDetail, error) {
	tweetDto, err := i.TweetRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	userID := tweetDto.UserID
	userDto, err := i.UserRepo.FindByID(userID)
	if err != nil {
		return nil, err
	}

	commentsDto, err := i.CommentRepo.FindByTweetID(id)
	if err != nil {
		return nil, err
	}

	var ids []uint
	for _, comment := range *commentsDto {
		ids = append(ids, comment.UserID)
	}
	usersDto, err := i.UserRepo.FindInUserID(ids)
	if err != nil {
		return nil, err
	}

	comments := i.CommentService.NewCommentByDtos(commentsDto, usersDto)
	fmt.Println(comments)
	tweet := &entity.TweetDetail{
		ID:       tweetDto.ID,
		UserID:   userID,
		Author:   userDto.Name,
		Avatar:   userDto.Avatar,
		Image:    tweetDto.Image,
		Title:    tweetDto.Title,
		Text:     tweetDto.Text,
		PostDate: tweetDto.CreatedAt.Format("2006/01/02 03:04"),
		Comments: *comments,
	}
	return tweet, nil
}

func (i *pictweetInteractor) AddComment(commentDto *dto.CommentDto) error {
	if err := i.CommentRepo.Create(commentDto); err != nil {
		return err
	}
	return nil
}
