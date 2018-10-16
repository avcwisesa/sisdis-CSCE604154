//go:generate swagger generate spec
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

	// swagger:operation GET /api/spesifikasi.yaml spesifikasi
	//
	// Returns API specifications
	// ---
	// Produces:
	// - text/x-yaml
	//
	// Schemes: ['http']
	//
	// Responses:
	//  200:
	//	  description: successful operation
	//  500:
	//	  description: server error
	router.get(`/api/spesifikasi.yaml`, handleSpesifikasi)

	// swagger:operation GET /api/plusone/{val} plusone
	//
	// Returns an incremented value
	// ---
	// Produces:
	// - application/json
	//
	// Schemes: ['http']
	//
	// Parameters:
	// - name: "val"
	//   in: "path"
	//   description: "value that need to be incremented"
	//   required: true
	//   type: "integer"
	//   format: "int32"
	//
	// Responses:
	//  200:
	//   description: successful operation
	//  400:
	//   description: Not a number
	router.get(`/api/plusone/(?P<number>\d+)`, handlePlusOne)

	// swagger:operation POST /api/hello hello
	//
	// Hello
	// ---
	// Consumes:
	// - application/json
	//
	// Produces:
	// - application/json
	//
	// Schemes: ['http']
	//
	// Parameters:
	// - name: body
	//   in: body
	//   required: true
	//
	// Responses:
	//  200:
	//   description: successful operation
	//  400:
	//   description: request is a required property
	//  500:
	//   description: server error
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



