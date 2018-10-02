package main

import (
	"net/url"
	"strings"
	"encoding/binary"
	"strconv"
	"fmt"
	"bufio"
)

var (
	statusOK = "200 OK\n"
	statusFound = "302 Found\n"
	statusBadRequest = "400 Bad Request\n"
	statusNotFound = "404 Not Found\n"
	statusInternalServerError = "500 Internal Server Error\n"
	statusNotImplemented = "501 Not Implemented\n"
)

type Request struct {
	method string
	url *url.URL
	protocol string
	header map[string]string
	body string
}

type Response struct {
	status, contentType, content string
	extraHeader map[string]string
}

func parseRequest(msgString string) *Request {

	msgSlice := strings.Split(msgString, string([]byte{13, 10, 13, 10}))
	msgHeader := strings.Split(msgSlice[0], "\n")
	msgBody := msgSlice[1]

	var header map[string]string
	var reqUrl *url.URL
	var method string
	var protocol string

	header = make(map[string]string)
	for i, line := range msgHeader {
		if i > 0 {
			lineSlice := strings.Split(line, ": ")
			if len(lineSlice) == 2 {
				header[lineSlice[0]] = strings.TrimSpace(lineSlice[1])
			}
		} else {
			lineSlice := strings.Split(line, " ")
			method = lineSlice[0]
			reqUrl, _ = url.Parse(lineSlice[1])
			protocol = strings.TrimSpace(lineSlice[2])
		}
	}

	return &Request{
		method: method,
		url: reqUrl,
		protocol: protocol,
		header: header,
		body: msgBody,
	}
}

func buildResponse(w *bufio.Writer, resp *Response) error {

	responseStatus := fmt.Sprintf("HTTP/1.1 %s", resp.status)
	w.WriteString(responseStatus)

	w.WriteString("Content-Type: " + resp.contentType + "\n")

	contentLength := binary.Size([]byte(resp.content))
	w.WriteString("Content-Length: " + strconv.Itoa(contentLength) + "\n")

	if len(resp.extraHeader) > 0 {
		for key, value := range resp.extraHeader {
			w.WriteString(key + ": " + value + "\n")
		}
	}

	w.WriteString("Connection: close\n\n")
	w.WriteString(resp.content)

	return nil
}