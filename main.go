package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	route "github.com/jojiejo/hibank-loan/app"
	"github.com/sirupsen/logrus"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	logger := logrus.New()
	logger.SetFormatter(&logrus.TextFormatter{})

	router := gin.Default()
	route.GenerateRoute(router)

	router.Run(fmt.Sprintf(":%s", os.Getenv("SERVER_PORT")))
}
