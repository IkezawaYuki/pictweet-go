package controller

import (
	"github.com/IkezawaYuki/pictweet-go/interface/adapter"
	"github.com/IkezawaYuki/pictweet-go/interface/port"
	"github.com/IkezawaYuki/pictweet-go/usecase"
	"github.com/labstack/echo"
	"net/http"
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
		return c.JSON(http.StatusOK, tweets)
	}
}
