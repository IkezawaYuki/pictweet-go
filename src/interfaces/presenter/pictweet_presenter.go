package presenter

import (
	"github.com/IkezawaYuki/pictweet-go/src/domain/model"
	"github.com/IkezawaYuki/pictweet-go/src/interfaces/rpc/pictweetpb"
)

type PictweetPresenter struct {
}

func (p *PictweetPresenter) ListTweet(tweets *model.Tweets) (*pictweetpb.ListTweetsResponse, error) {
	var res []*pictweetpb.Tweet
	for _, t := range *tweets {
		res = append(res, &pictweetpb.Tweet{
			Id:        int32(t.ID),
			Title:     t.Title,
			Content:   t.Text,
			ImageUrl:  t.Image,
			Author:    t.User.Name,
			AvatarUrl: t.User.Avatar,
			CreatedAt: t.CreatedAt.Format("2006/01/02"),
		})
	}
	return &pictweetpb.ListTweetsResponse{
		Tweet: res,
	}, nil
}

func (p *PictweetPresenter) ShowTweet(tweet *model.Tweet) (*pictweetpb.ShowTweetResponse, error) {
	var commentRes []*pictweetpb.Comment
	for _, c := range tweet.Comment {
		commentRes = append(commentRes, &pictweetpb.Comment{
			Id:        int32(c.ID),
			Text:      c.Text,
			UserId:    int32(c.UserID),
			CreatedAt: c.CreatedAt.Format("2006/01/02"),
		})
	}

	return &pictweetpb.ShowTweetResponse{
		Tweet: &pictweetpb.Tweet{
			Id:        int32(tweet.ID),
			Title:     tweet.Title,
			Content:   tweet.Text,
			ImageUrl:  tweet.Image,
			Author:    tweet.User.Name,
			AvatarUrl: tweet.User.Avatar,
			CreatedAt: tweet.CreatedAt.Format("2006/01/02"),
		},
		Comment: commentRes,
	}, nil
}
