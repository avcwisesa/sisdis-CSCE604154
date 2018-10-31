package main

import (
	"fmt"
	"log"
	"os"
	"io/ioutil"
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"

	h "github.com/avcwisesa/sisdis/tugas1/handler"
	c "github.com/avcwisesa/sisdis/tugas1/controller"
	d "github.com/avcwisesa/sisdis/tugas1/database"
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
	dbName := os.Getenv("MYSQL_DATABASE_NAME")

	client, err := gorm.Open("mysql",
		fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
			dbUsername,
			dbPassword,
			dbName,
		),
	)
	if err != nil {
		log.Println("Error connecting to DB")
		panic(err)
	}

	jsonFile, err := os.Open("quorum.json")
	if err != nil {
		panic(err)
	}
	defer jsonFile.Close()

	buf, _ := ioutil.ReadAll(jsonFile)

	var quorum map[string]string
	json.Unmarshal(buf, &quorum)

	port := os.Getenv("WEB_PORT")
	host := os.Getenv("EWALLET_HOST")

	db := d.New(client)
	controller := c.New(host, db)

	handler := h.New(len(quorum), controller)

	r.POST("/ewallet/ping", handler.Ping)
	r.POST("/ewallet/register", handler.Register)
	r.POST("/ewallet/getSaldo", handler.GetSaldo)
	r.POST("/ewallet/getTotalSaldo", handler.GetTotalSaldo)
	r.POST("/ewallet/transfer", handler.Transfer)

	r.Run(":" + port)
}