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

	dbUsername := os.Getenv("MYSQL_USERNAME")
	dbPassword := os.Getenv("MYSQL_PASSWORD")
	dbHost := os.Getenv("MYSQL_HOST")
	dbName := os.Getenv("MYSQL_DATABASE_NAME")

	client, err := gorm.Open("mysql",
		fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=True&loc=Local",
			dbUsername,
			dbPassword,
			dbHost,
			dbName,
		),
	)
	if err != nil {
		log.Println("Error connecting to DB")
		panic(err)
	}

	port := os.Getenv("WEB_PORT")

	db := d.New(client)
	controller := c.New(db)

	handler := handler.New()

	r.POST("/ewallet/ping", handler.Ping)
	r.POST("/ewallet/register", handler.Register)
	r.POST("/ewallet/getSaldo", handler.GetSaldo)
	r.POST("/ewallet/getTotalSaldo", handler.GetTotalSaldo)
	r.POST("/ewallet/transfer", handler.Transfer)

	r.Run(":" + port)
}