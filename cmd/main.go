package main

import (
	"log"
	"os"

	http "github.com/gabmenezesdev/go-tech-challenge/internal/infra/http"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var (
	APP_PORT = "APP_PORT"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router := gin.Default()
	http.InitRoutes(router)

	router.Run(os.Getenv(APP_PORT))
}
