package main

import (
	"bufio"
	"regexp"
	"encoding/json"
	"fmt"
	"net"
)

type Router struct {
	handlers []*Handler
	defaultHandlers []*Handler
}

func newRouter() *Router {

	proto := &Router{
		handlers: make([]*Handler, 0),
		defaultHandlers: make([]*Handler, 0),
	}

	proto.defaultHandlers = append(proto.defaultHandlers, &Handler{
		method: "GET",
		handle: defaultHandleGet,
	})
	proto.defaultHandlers = append(proto.defaultHandlers, &Handler{
		method: "POST",
		handle: defaultHandlePost,
	})

	return proto
}

func (r *Router) handleConnection(conn net.Conn) {
	defer func() {
		conn.Close()
		// fmt.Println("Connection closed.")
	}()

	msg := make([]byte, 1024)
	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)

	len, err := reader.Read(msg)
	if err != nil {
		fmt.Println("Read error: " + err.Error())
	}
	// fmt.Printf("Length: %d\n", len)
	msgString := string(msg[:len])

	if len > 0 {
		request := parseRequest(msgString)

		resp := r.route(request)

		buildResponse(writer, resp)
		writer.Flush()
	}
}

func (r *Router) route (req *Request) *Response {
	fmt.Printf("%-10v|%-10v| %s\n", req.protocol, req.method, req.url)

	if req.protocol != "HTTP/1.1" && req.protocol != "HTTP/1.0" {
		return &Response{
			status: statusBadRequest,
			contentType: "text/plain; charset=UTF-8",
			content: "",
		}
	}

	for _, handler := range r.handlers {
		if handler.method == req.method && handler.url.MatchString(req.url.Path) {
			fmt.Printf("match %s | %s\n",handler.url.String(),req.url.Path)
			req.regex = handler.url
			return handler.handle(req)
		}

	}

	for _, handler := range r.defaultHandlers {
		if handler.method == req.method {
			return handler.handle(req)
		}

	}

	return &Response{
		status: statusNotImplemented,
		contentType: "text/plain; charset=UTF-8",
		content: "",
	}
}

func (r *Router) get(url string, handler func(*Request) *Response) {
	re, _ := regexp.Compile("^" + url + "(/??)$")
	// fmt.Println(re.String())

	r.handlers = append(r.handlers, &Handler{
		method: "GET",
		url: re,
		handle: handler,
	})
}

func (r *Router) post(url string, handler func(*Request) *Response) {
	re, _ := regexp.Compile("^" + url + "(/??)$")
	// fmt.Println(re.String())

	r.handlers = append(r.handlers, &Handler{
		method: "POST",
		url: re,
		handle: handler,
	})
}

func defaultHandleGet(req *Request) *Response {

	content := make(map[string]interface{})

	content["detail"] = "The requested URL was not found on the server.  If you entered the URL manually please check your spelling and try again."
	content["status"] = 404
	content["title"] = "Not Found"

	jsonContent, _ := json.Marshal(content)

	return &Response{
		status: statusNotFound,
		contentType: "application/json",
		content: string(jsonContent),
	}
}

func defaultHandlePost(req *Request) *Response {

	content := make(map[string]interface{})

	content["detail"] = "The requested URL was not found on the server.  If you entered the URL manually please check your spelling and try again."
	content["status"] = 404
	content["title"] = "Not Found"

	jsonContent, _ := json.Marshal(content)

	return &Response{
		status: statusNotFound,
		contentType: "application/json",
		content: string(jsonContent),
	}
}
