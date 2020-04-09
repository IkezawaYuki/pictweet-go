package main

import (
	"context"
	"flag"
	"github.com/IkezawaYuki/pictweet-go/src/interfaces/rpc"
	"github.com/IkezawaYuki/pictweet-go/src/interfaces/rpc/pictweetpb"
	"github.com/IkezawaYuki/pictweet-go/src/registry"
	"github.com/IkezawaYuki/pictweet-go/src/usecase"
	"github.com/gorilla/handlers"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
)

var (
	endpoint = flag.String("endpoint", "localhost:50051", "endpoint of pictweetService")
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	err := godotenv.Load()
	if err != nil {
		logrus.Fatalf("Error loading .env")
	}

	log.Println("grpc server set up...")
	port := os.Getenv("PORT")
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("failed to listend: %v", err)
	}
	ctn, err := registry.NewContainer()
	if err != nil {
		log.Fatalf("failed to build container: %v", err)
	}
	s := grpc.NewServer(
		grpc.UnaryInterceptor(
			grpc_middleware.ChainUnaryServer(
				grpc_auth.UnaryServerInterceptor(rpc.AuthFunc),
				rpc.AuthorizationUnaryServerInterceptor(),
			),
		),
	)

	pictweetpb.RegisterPictweetServiceServer(
		s,
		rpc.NewPictweetService(ctn.Resolve("tweet-usecase").(usecase.PictweetUsecase)),
	)
	reflection.Register(s)
	go func() {
		log.Printf("start grpc server port: %s", port)
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	log.Println("reverse proxy server set up...")
	ctx, cancel := context.WithCancel(context.Background())
	mux := runtime.NewServeMux()
	newMux := handlers.CORS(
		handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS", "DELETE"}),
		handlers.AllowedOrigins([]string{"http://localhost:3000"}),
		handlers.AllowedHeaders([]string{"Content-Type", "application/json"}),
	)(mux)

	opts := []grpc.DialOption{grpc.WithInsecure()}

	err = pictweetpb.RegisterPictweetServiceHandlerFromEndpoint(ctx, mux, *endpoint, opts)
	if err != nil {
		log.Fatalf("failed to register serve: %v", err)
	}
	reversePort := os.Getenv("REVERSE_PROXY_PORT")
	go func() {
		log.Printf("start reverse proxy server port: %s", reversePort)
		if err := http.ListenAndServe(":8080", newMux); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("stopping grpc server...")
	s.GracefulStop()
	ctn.Clean()
	cancel()
}
