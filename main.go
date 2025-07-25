package main

import (
	"log"
	"os"
	"supplier-management/config"
	"supplier-management/router"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	r := gin.Default()
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
		panic("Error loading .env file")
	}
	config.Connect()
	path := os.Getenv("SERVER_PATH")
	router.SetupRoutes(r, path)
	port := os.Getenv("SERVICE_PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)
}
