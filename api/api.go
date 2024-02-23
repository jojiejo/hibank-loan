package api

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/jojiejo/hibank-loan/api/handlers"
)

var server = handlers.Server{}

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func Run() {
	server.Initialize()
	apiPort := fmt.Sprintf(":%s", os.Getenv("SERVER_PORT"))
	fmt.Printf("\nListening to port %s", apiPort)
	server.Run(apiPort)
}
