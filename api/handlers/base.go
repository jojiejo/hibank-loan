package handlers

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
)

const logDirectory = "logs"
const logFileName = "loan-api.log"

type Server struct {
	Router   *gin.Engine
	Logger   *logrus.Logger
	Validate *validator.Validate
}

func (server *Server) Initialize() {
	gin.SetMode(gin.ReleaseMode)
	server.Router = gin.New()
	server.InitializeLogger()
	server.GenerateRoutes()
	server.Validate = validator.New(validator.WithRequiredStructEnabled())
}

func (server *Server) InitializeLogger() {
	server.Logger = logrus.New()

	logFilePath := filepath.Join(logDirectory, logFileName)
	err := os.MkdirAll(logDirectory, 0755)
	if err != nil {
		log.Fatal("Failed to create log directory:", err)
	}

	file, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Failed to open log file:", err)
	}
	server.Logger.SetOutput(file)
	server.Router.Use(gin.LoggerWithWriter(file))
}

func (server *Server) Run(addrress string) {
	log.Fatal(http.ListenAndServe(addrress, server.Router))
}
