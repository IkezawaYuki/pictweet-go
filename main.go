package main

import (
	"github.com/IkezawaYuki/pictweet-go/infrastructure"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		logrus.Fatalf("Error loading .env")
	}
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(&logrus.JSONFormatter{})
}

func main() {
	infrastructure.StartApplication()
}
