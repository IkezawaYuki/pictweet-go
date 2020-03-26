package rpc

import (
	"context"
	"fmt"
	"github.com/IkezawaYuki/pictweet-go/src/domain/model"
	"github.com/IkezawaYuki/pictweet-go/src/interfaces/presenter"
	"github.com/IkezawaYuki/pictweet-go/src/interfaces/rpc/pictweetpb"
	"github.com/IkezawaYuki/pictweet-go/src/usecase"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

type pictweetService struct {
	pictweetUsecase   usecase.PictweetUsecase
	pictweetPresenter presenter.PictweetPresenter
}

func NewPictweetService(u usecase.PictweetUsecase) pictweetpb.PictweetServiceServer {
	return &pictweetService{
		pictweetUsecase:   u,
		pictweetPresenter: presenter.PictweetPresenter{},
	}
}

func Interceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		fmt.Println("interceptor")
		resp, err := handler(ctx, req)
		if err != nil {
			return nil, err
		}
		return resp, err
	}
}

func (p *pictweetService) ListTweets(ctx context.Context, req *pictweetpb.ListTweetsRequest) (*pictweetpb.ListTweetsResponse, error) {
	tweet, err := p.pictweetUsecase.ListTweets()
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal error: %v", err),
		)
	}
	return p.pictweetPresenter.ListTweet(tweet)
}

func (p *pictweetService) PostTweet(ctx context.Context, req *pictweetpb.PostTweetRequest) (*pictweetpb.PostTweetResponse, error) {
	id, err := p.pictweetUsecase.PostTweet(&model.Tweet{
		UserID:    uint(req.GetUserId()),
		Image:     req.GetImageUrl(),
		Title:     req.GetTitle(),
		Text:      req.GetContent(),
		CreatedAt: time.Now(),
	})
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("internal error: %v", err),
		)
	}
	return &pictweetpb.PostTweetResponse{
		Result: fmt.Sprintf("success tweet_id: %d", id),
	}, nil
}

func (p *pictweetService) ShowTweet(ctx context.Context, req *pictweetpb.ShowTweetRequest) (*pictweetpb.ShowTweetResponse, error) {
	tweets, err := p.pictweetUsecase.ShowTweet(uint(req.GetTweetId()))
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("internal error: %v", err),
		)
	}
	return p.pictweetPresenter.ShowTweet(tweets)
}

func (p *pictweetService) PostComment(ctx context.Context, req *pictweetpb.PostCommentRequest) (*pictweetpb.PostCommentResponse, error) {
	id, err := p.pictweetUsecase.PostComment(&model.Comment{
		UserID:    uint(req.GetUserId()),
		TweetID:   uint(req.GetTweetId()),
		Text:      req.GetText(),
		CreatedAt: time.Now(),
	})
	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("invalide argument: %v", err),
		)
	}
	return &pictweetpb.PostCommentResponse{
		Result: fmt.Sprintf("success comment_id: %d", id),
	}, nil
}

func (p *pictweetService) DeleteTweet(ctx context.Context, req *pictweetpb.DeleteTweetRequest) (*pictweetpb.DeleteTweetResponse, error) {
	err := p.pictweetUsecase.DeleteTweet(uint(req.GetTweetId()))
	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("invalide argument: %v", err),
		)
	}
	return &pictweetpb.DeleteTweetResponse{
		Result: fmt.Sprintf("success delete tweet_id: %d", req.GetTweetId()),
	}, nil
}

func (p *pictweetService) RegisterUser(ctx context.Context, req *pictweetpb.RegisterUserRequest) (*pictweetpb.RegsiterUserResponse, error) {
	id, err := p.pictweetUsecase.RegisterUser(&model.User{
		Name:      req.GetName(),
		Email:     req.GetEmail(),
		Avatar:    req.GetAvatar(),
		CreatedAt: time.Now(),
	})
	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("invalide argument: %v", err),
		)
	}
	return &pictweetpb.RegsiterUserResponse{
		Result: fmt.Sprintf("success user_id: %d", id),
	}, nil
}

func (p *pictweetService) FetchTweets(ctx context.Context, req *pictweetpb.FetchTweetsRequest) (*pictweetpb.FetchTweetsResponse, error) {
	fmt.Println("fetch tweets")
	fmt.Println(req.GetEmail())
	tweets, err := p.pictweetUsecase.GetFavorite(req.GetEmail())
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal error: %v", err),
		)
	}
	return p.pictweetPresenter.ListFavorite(tweets)
}
