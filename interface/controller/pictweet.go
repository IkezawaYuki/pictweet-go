package controller

import (
	"fmt"
	"github.com/IkezawaYuki/pictweet-go/domain/dto"
	"github.com/IkezawaYuki/pictweet-go/domain/service"
	"github.com/IkezawaYuki/pictweet-go/interface/adapter"
	"github.com/IkezawaYuki/pictweet-go/interface/port"
	"github.com/IkezawaYuki/pictweet-go/usecase"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

type pictweetController struct {
	Interactor port.InputPort
}

func NewPictweetController(h adapter.SQLHandler) *pictweetController {
	userRepo := adapter.NewUserAdapter(h)
	return &pictweetController{
		Interactor: usecase.NewPictweetInteractor(
			adapter.NewTweetRepository(h),
			adapter.NewCommentAdapter(h),
			userRepo,
			adapter.NewFavoriteRepository(h),
			*service.NewUserService(userRepo),
		)}
}

func (p *pictweetController) FetchTweets() echo.HandlerFunc {
	return func(c echo.Context) error {
		tweets, err := p.Interactor.Index()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}
		return c.JSON(http.StatusOK, tweets)
	}
}

func (p *pictweetController) FetchFavorites() echo.HandlerFunc {
	return func(c echo.Context) error {
		email := c.FormValue("email")
		tweetID := c.FormValue("tweet_id")
		tweetId, err := strconv.ParseUint(tweetID, 10, 64)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		isFavorite, err := p.Interactor.ToggleFavorite(email, uint(tweetId))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}

		res := dto.ToggleFavoriteVideoResponse{
			TweetID:    tweetID,
			IsFavorite: isFavorite,
		}
		return c.JSON(http.StatusOK, res)
	}
}

func (p *pictweetController) PostTweet() echo.HandlerFunc {
	return func(c echo.Context) error {
		title := c.FormValue("title")
		image := c.FormValue("url")
		text := c.FormValue("comment")
		email := c.FormValue("email")
		fmt.Println(email)
		userID := "1"

		tweet, err := dto.NewTweetDto(userID, image, title, text)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		insertId, err := p.Interactor.CreateTweet(tweet)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}

		insertTweet, err := p.Interactor.FindByID(insertId)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}

		return c.JSON(http.StatusCreated, insertTweet)
	}
}

func (p *pictweetController) ShowTweet() echo.HandlerFunc {
	return func(c echo.Context) error {
		tweetID := c.Param("id")
		tweetId, err := strconv.ParseUint(tweetID, 10, 64)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		tweet, err := p.Interactor.ShowTweet(uint(tweetId))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}

		return c.JSON(http.StatusOK, tweet)
	}
}

func (p *pictweetController) AddComment() echo.HandlerFunc {
	return func(c echo.Context) error {
		tweetID := c.Param("tweetId")
		userID := c.FormValue("user_id")
		text := c.FormValue("comment")

		commentDto, err := dto.NewCommentDto(tweetID, userID, text)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		if err := p.Interactor.AddComment(commentDto); err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}

		return c.JSON(http.StatusCreated, commentDto)
	}
}

func (p *pictweetController) CreateUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		name := c.FormValue("name")
		email := c.FormValue("email")
		avatar := c.FormValue("avatar")
		userDto := dto.UserDto{
			Name:   name,
			Email:  email,
			Avatar: avatar,
		}
		if err := p.Interactor.CreateUser(&userDto); err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		return c.JSON(http.StatusCreated, userDto)
	}
}
