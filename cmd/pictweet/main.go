package main

import (
	"github.com/IkezawaYuki/pictweet-go/src/interfaces/rpc"
	"github.com/IkezawaYuki/pictweet-go/src/interfaces/rpc/pictweetpb"
	"github.com/IkezawaYuki/pictweet-go/src/registry"
	"github.com/IkezawaYuki/pictweet-go/src/usecase"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
	"os/signal"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	err := godotenv.Load()
	if err != nil {
		logrus.Fatalf("Error loading .env")
	}

	port := os.Getenv("PORT")
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("failed to listend: %v", err)
	}
	ctn, err := registry.NewContainer()
	if err != nil {
		log.Fatalf("failed to build container: %v", err)
	}

	s := grpc.NewServer()

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

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("stopping grpc server...")
	s.GracefulStop()
	ctn.Clean()
}
