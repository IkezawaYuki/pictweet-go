package rpc

import (
	"context"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"fmt"
	"google.golang.org/api/option"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"log"
	"os"
	"strings"
)

const TokenKey = "token"

var (
	authenticationRequired = []string{
		"/pictweet.PictweetService/PostTweet",
		"/pictweet.PictweetService/PostComment",
		"/pictweet.PictweetService/DeleteTweet",
		"/pictweet.PictweetService/RegisterUser",
		"/pictweet.PictweetService/FavoriteTweets",
	}
)

func AuthorizationUnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		if CheckAuthenticationRequired(info.FullMethod) {
			fmt.Println("AuthFunc")
			authClient := ctx.Value("firebase")
			fmt.Println(authClient)
			if authClient == nil {
				return ctx, nil
			}
			auth := authClient.(*auth.Client)
			_, err := VerifyFirebaseIDToken(ctx, auth)
			if err != nil {
				return nil, status.Errorf(
					codes.PermissionDenied,
					"jwt_token is null",
				)
			}
			resp, err := handler(ctx, req)
			if err != nil {
				return nil, err
			}
			return resp, err
		}
		resp, err := handler(ctx, req)
		if err != nil {
			return nil, err
		}
		return resp, err
	}
}

func CheckAuthenticationRequired(method string) bool {
	for _, m := range authenticationRequired {
		if m == method {
			return true
		}
	}
	return false
}

func VerifyFirebaseIDToken(ctx context.Context, auth *auth.Client) (*auth.Token, error) {
	fmt.Println("VerifyFirebase ID Token")
	headerAuth, ok := metadata.FromIncomingContext(ctx)
	fmt.Println(headerAuth)
	if !ok {
		return nil, fmt.Errorf("permission denied")
	}
	authString := headerAuth["authorization"]
	fmt.Println(authString)
	if len(authString) == 0 {
		return nil, fmt.Errorf("permission denied")
	}
	token := strings.Replace(authString[0], "Bearer ", "", 1)
	fmt.Println(token)
	jwtToken, err := auth.VerifyIDToken(context.Background(), token)
	return jwtToken, err
}

func AuthFunc(ctx context.Context) (context.Context, error) {
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
	return ctx, nil
}
