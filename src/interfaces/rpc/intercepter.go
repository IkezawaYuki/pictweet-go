package rpc

import (
	"context"
	"firebase.google.com/go"
	"firebase.google.com/go/auth"
	"google.golang.org/api/option"
	"google.golang.org/grpc"
	"log"
	"os"
)

var methods = []string{
	"/pictweet.PictweetService/ListTweets", // todo delete
	"/pictweet.PictweetService/PostTweet",
	"/pictweet.PictweetService/PostComment",
	"/pictweet.PictweetService/DeleteTweet",
	"/pictweet.PictweetService/RegisterUser",
}

func Interceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		if CheckAuthenticationRequired(info.FullMethod) {
			opt := option.WithCredentialsFile(os.Getenv("KEY_JSON_PATH"))
			config := &firebase.Config{
				ProjectID: os.Getenv("PROJECT_ID"),
			}
			app, err := firebase.NewApp(context.Background(), config, opt)
			if err != nil {
				log.Fatalf("error initialize firebase :%v", err)
			}

			auth, err := app.Auth(context.Background())
			if err != nil {
				log.Fatalf("error auth :%v", err)
			}
			ctx = context.WithValue(ctx, "firebase", auth)
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

func VerifyFirebaseIDToken(ctx context.Context, auth *auth.Client) (*auth.Token, error) {
	// todo
}
