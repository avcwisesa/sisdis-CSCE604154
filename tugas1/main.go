package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"

	"github.com/avcwisesa/sisdis/tugas1/handler"
)


func main() {

	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file, using to container ENV only.")
	}

	r := gin.Default()
	r.Use(cors.Default())

	port := os.Getenv("WEB_PORT")

	handler := handler.New()

	r.GET("/ewallet/ping", handler.Ping)

	r.Run(":" + port)
}