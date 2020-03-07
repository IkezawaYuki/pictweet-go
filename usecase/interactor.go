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
	FavoriteRepo   repository.FavoriteRepository
	TweetService   service.TweetService
	CommentService service.CommentService
	UserService    service.UserService
}

func NewPictweetInteractor(tRepo repository.TweetRepository,
	cRepo repository.CommentRepository,
	uRepo repository.UserRepository,
	fRepo repository.FavoriteRepository,
	uService service.UserService) port.InputPort {
	return &pictweetInteractor{
		TweetRepo:    tRepo,
		CommentRepo:  cRepo,
		UserRepo:     uRepo,
		FavoriteRepo: fRepo,
		UserService:  uService,
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

func (i *pictweetInteractor) FindByID(id uint) (*entity.Tweet, error) {
	tweetDto, err := i.TweetRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	userDto, err := i.UserRepo.FindByID(tweetDto.UserID)
	if err != nil {
		return nil, err
	}

	return i.TweetService.NewTweetByDto(tweetDto, userDto), nil
}

// CreateTweet ツイート新規投稿ロジック
func (i *pictweetInteractor) CreateTweet(dto *dto.TweetDto) (uint, error) {
	id, err := i.TweetRepo.Create(dto)
	if err != nil {
		return id, err
	}
	return id, nil
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

func (i *pictweetInteractor) CreateUser(userDto *dto.UserDto) error {
	if !i.UserService.Exist(userDto.Email) {
		return fmt.Errorf("user is exist. email: %s", userDto.Email)
	}
	if err := i.UserRepo.Create(userDto); err != nil {
		return err
	}
	return nil
}

func (i *pictweetInteractor) ToggleFavorite(email string, tweetID uint) (bool, error) {
	user, err := i.UserRepo.FindByEmail(email)
	if err != nil {
		return false, err
	}
	if &user == nil {
		return false, fmt.Errorf("user is not found")
	}
	return i.FavoriteRepo.Toggle(user.ID, tweetID)
}
