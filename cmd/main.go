package main

import (
	"log"
	"os"

	"github.com/gabmenezesdev/go-tech-challenge/docs"
	http "github.com/gabmenezesdev/go-tech-challenge/internal/infra/http"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var (
	APP_PORT = "APP_PORT"
)

// @title GO-TECH-CHALLENGE
// @version 1.0
// @description Tech challenge API
// @host localhost:3000
// @BasePath /
// @schemes http

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router := gin.Default()
	docs.SwaggerInfo.BasePath = "/api/v1"
	http.InitRoutes(router)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	router.Run(os.Getenv(APP_PORT))
}
