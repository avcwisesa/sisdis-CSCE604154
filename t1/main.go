package main

import (
	"github.com/joho/godotenv"

	"log"
	"net"
	"fmt"
	"os"
)

var count = 0

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var port = ":" + os.Getenv("PORT")

	ln, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Listening on port %s\n", port)

	defer func() {
		ln.Close()
		fmt.Println("Server stoped listening.")
	}()

	router := newRouter()

	router.get(`/`, handleRedirectHello)
	router.get(`/hello-world`, handleHello)
	router.get(`/style`, handleStyle)
	router.get(`/background`, handleBackground)
	router.get(`/info`, handleInfo)

	router.post(`/`, handleRedirectHello)
	router.post(`/hello-world`, handleHelloPost)

	router.get(`/api/(?P<number>\d+)`, handlePlusOne)
	router.post(`/api/hello`, handleHelloApi)

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Connection error: " + err.Error())
		}
		// fmt.Println("Connection established!")

		router.handleConnection(conn)
	}
}



