package controller

import (
	"fmt"
	"github.com/IkezawaYuki/pictweet-go/domain/dto"
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
	return &pictweetController{
		Interactor: usecase.NewPictweetInteractor(
			adapter.NewTweetRepository(h),
			adapter.NewCommentAdapter(h),
			adapter.NewUserAdapter(h),
		)}
}

func (p *pictweetController) FetchTweets() echo.HandlerFunc {
	return func(c echo.Context) error {
		tweets, err := p.Interactor.Index()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}
		fmt.Println(tweets)
		return c.JSON(http.StatusOK, tweets)
	}
}

func (p *pictweetController) PostTweet() echo.HandlerFunc {
	return func(c echo.Context) error {
		title := c.FormValue("title")
		image := c.FormValue("url")
		text := c.FormValue("comment")
		userID := "1" // todo

		tweet, err := dto.NewTweetDto(userID, image, title, text)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		if err := p.Interactor.CreateTweet(tweet); err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}

		//return c.JSON(http.StatusCreated, "OK!")
		return c.JSON(http.StatusCreated, fmt.Errorf("error"))
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
