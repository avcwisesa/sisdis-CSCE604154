package main

import (
	"bufio"
	"net"
	"fmt"
)

func main() {
	var port = ":8080"

	ln, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Listening on port %s\n", port)

	defer func() {
		ln.Close()
		fmt.Println("Server stoped listening.")
	}()

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Connection error: " + err.Error())
		}
		fmt.Println("Connection established!")

		handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer func() {
		conn.Close()
		fmt.Println("Connection closed.")
	}()

	msg := make([]byte, 1024)
	r := bufio.NewReader(conn)
	w := bufio.NewWriter(conn)

	len, err := r.Read(msg)
	if err != nil {
		fmt.Println("Read error: " + err.Error())
	}
	fmt.Printf("Length: %d\n", len)
	msgString := string(msg[:len])

	if len > 0 {
		request := parseRequest(msgString)

		resp := handleRequest(request)

		buildResponse(w, resp)
		w.Flush()
	}
}


