package infrastructure

import (
	"github.com/IkezawaYuki/pictweet-go/interface/controller"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func StartApplication() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.CORS())

	//e.Use(middlewares.YoutubeService())
	//e.Use(middlewares.DatabaseService())
	//e.Use(middlewares.Firebase())
	ctr := controller.NewPictweetController(NewSqlHandler())

	g := e.Group("/api")
	{
		g.GET("/tweets", ctr.FetchTweets())
		g.POST("/tweet/", ctr.PostTweet())
		g.GET("/tweet/:id", ctr.ShowTweet())
		g.POST("/comment/:tweetId", ctr.AddComment())
	}

	//fg := g.Group("/favorite", middlewares.FirebaseGuard())
	//{
	//
	//}
	e.Logger.Fatal(e.Start(":8080"))
}
