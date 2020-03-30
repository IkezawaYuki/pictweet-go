package rpc

import (
	"context"
	"firebase.google.com/go"
	"fmt"
	"google.golang.org/api/option"
	"google.golang.org/grpc"
	"os"
)

var methods = []string{
	"/pictweet.PictweetService/PostTweet",
	"/pictweet.PictweetService/PostComment",
	"/pictweet.PictweetService/DeleteTweet",
	"/pictweet.PictweetService/RegisterUser",
}

func Interceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		fmt.Println(info.FullMethod)

		if CheckAuthenticationRequired(info.FullMethod) {
			opt := option.WithCredentialsFile(os.Getenv("KEY_JSON_PATH"))
			config := &firebase.Config{
				ProjectID: os.Getenv("PROJECT_ID"),
			}
			app, err := firebase.NewApp(ctx, config, opt)
		}
		resp, err := handler(ctx, req)
		if err != nil {
			return nil, err
		}
		return resp, err
	}
}

func CheckAuthenticationRequired(method string) bool {
	for _, m := range methods {
		if m == method {
			return true
		}
	}
	return false
}
