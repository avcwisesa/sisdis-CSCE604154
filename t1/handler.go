package main

import (
	"net/url"
	"strconv"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
)

type Handler struct {
	method string
	url string
	handle func(*Request) *Response
}

func handleRequest(req *Request) *Response {

	router := newRouter()

	router.get("/hello-world", handleHello)
	router.get("/style", handleStyle)
	router.get("/background", handleBackground)
	router.get("/info", handleInfo)
	router.get("/", handleRedirectHello)

	router.post("/", handleRedirectHello)
	router.post("/hello-world", handleHelloPost)

	return router.route(req)
}

func handleHello(req *Request) *Response {
	file, err := ioutil.ReadFile("hello-world.html")
	if err != nil {
		return &Response{
			status: statusInternalServerError,
			contentType: "text/plain; charset=UTF-8",
			content: "",
		}
	}

	content := strings.Replace(string(file), "__HELLO__", "World", -1)

	return &Response{
		status: statusOK,
		contentType: "text/html; charset=UTF-8",
		content: content,
	}
}

func handleStyle(req *Request) *Response {
	file, err := ioutil.ReadFile("style.css")
	if err != nil {
		return &Response{
			status: statusInternalServerError,
			contentType: "text/plain; charset=UTF-8",
			content: "",
		}
	}

	return &Response{
		status: statusOK,
		contentType: "text/css; charset=UTF-8",
		content: string(file),
	}
}

func handleBackground(req *Request) *Response {
	file, err := ioutil.ReadFile("background.jpg")
	if err != nil {
		return &Response{
			status: statusInternalServerError,
			contentType: "image/jpeg",
			content: "",
		}
	}

	return &Response{
		status: statusOK,
		contentType: "text/html; charset=UTF-8",
		content: string(file),
	}
}

func handleInfo(req *Request) *Response {
	var content string

	switch query := req.url.Query(); query.Get("type") {
	case "random":
		content = strconv.Itoa(rand.Int())
	case "time":
		t := time.Now()
		content = t.Format("2006-01-02 15:04:05")
	default:
		content = "No Data"
	}

	return &Response{
		status: statusOK,
		contentType: "text/plain; charset=UTF-8",
		content: content,
	}
}

func handleRedirectHello(req *Request) *Response {

	extraHeader := make(map[string]string)
	extraHeader["Location"] = "/hello-world"

	return &Response{
		status: statusFound,
		contentType: "text/plain",
		content: "",
		extraHeader: extraHeader,
	}
}

func handleHelloPost(req *Request) *Response {

	if req.header["Content-Type"] != "application/x-www-form-urlencoded" {
		return &Response{
			status: statusBadRequest,
			contentType: "text/plain",
			content: "",
		}
	}

	params, err := url.ParseQuery(req.body)
	if err != nil {
		return &Response{
			status: statusInternalServerError,
			contentType: "text/plain; charset=UTF-8",
			content: err.Error(),
		}
	}

	file, err := ioutil.ReadFile("hello-world.html")
	if err != nil {
		return &Response{
			status: statusInternalServerError,
			contentType: "text/plain; charset=UTF-8",
			content: err.Error(),
		}
	}

	content := strings.Replace(string(file), "__HELLO__", params.Get("name"), -1)

	return &Response{
		status: statusOK,
		contentType: "text/html; charset=UTF-8",
		content: content,
	}
}