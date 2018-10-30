package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
	"github.com/avcwisesa/sisdis/tugas1/database"

	m "github.com/avcwisesa/sisdis/tugas1/model"
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

	db := database.New(client)

	db.Migrate(&m.Customer{})
}
