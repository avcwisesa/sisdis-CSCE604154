package main

import (
	"fmt"
	"encoding/json"
	"net/url"
	"os"
	"strconv"
	"io/ioutil"
	"math/rand"
	"net/http"
	"regexp"
	"strings"
	"time"
)

type Handler struct {
	method string
	url *regexp.Regexp
	handle func(*Request) *Response
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

func handlePlusOne(req *Request) *Response {

	path := strings.Split(req.url.Path, "/")
	param, err := strconv.Atoi(path[3])
	if err != nil {
		return &Response{
			status: statusBadRequest,
			contentType: "text/plain",
			content: "",
		}
	}

	content := make(map[string]interface{})

	content["apiversion"] = 1
	content["plusoneret"] = param + 1

	jsonContent, _ := json.Marshal(content)

	return &Response{
		status: statusOK,
		contentType: "application/json",
		content: string(jsonContent),
	}
}

func handleHelloApi(req *Request) *Response {

	if req.header["Content-Type"] != "application/json" {
		return &Response{
			status: statusBadRequest,
			contentType: "text/plain",
			content: "",
		}
	}

	now := time.Now()
	content := make(map[string]interface{})

	var params map[string]interface{}

	json.Unmarshal([]byte(req.body), &params)

	fmt.Println(params)
	if params["request"] == nil {
		content["title"] = "Bad Request"
		content["status"] = 400
		content["detail"] = "'request' is a required property"

		jsonContent, _ := json.Marshal(content)

		return &Response{
			status: statusBadRequest,
			contentType: "application/json",
			content: string(jsonContent),
		}
	}

	timeApi := os.Getenv("TIME_API")

	resp, err := http.Get(timeApi)
	if err != nil {
		return &Response{
			status: statusInternalServerError,
			contentType: "text/plain",
			content: err.Error(),
		}
	}

	var timeResponse map[string]interface{}

	buf, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(buf, &timeResponse)

	count = count + 1

	content["apiversion"] = 1
	content["count"] = count
	content["currentvisit"] = now
	content["response"] = fmt.Sprintf("Good %s, %s", timeResponse["state"], params["request"])

	jsonContent, _ := json.Marshal(content)

	return &Response{
		status: statusOK,
		contentType: "application/json",
		content: string(jsonContent),
	}
}
