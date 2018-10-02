package main

import (
	"fmt"
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

func (r *Router) route (req *Request) *Response {
	fmt.Printf("Protocol: %s\nMethod: %s\nURL: %s\n", req.protocol, req.method, req.url)

	if req.protocol != "HTTP/1.1" && req.protocol != "HTTP/1.0" {
		return &Response{
			status: statusBadRequest,
			contentType: "text/plain; charset=UTF-8",
			content: "",
		}
	}

	for _, handler := range r.handlers {
		if handler.method == req.method && handler.url == req.url.Path {
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
	r.handlers = append(r.handlers, &Handler{
		method: "GET",
		url: url,
		handle: handler,
	})
}

func (r *Router) post(url string, handler func(*Request) *Response) {
	r.handlers = append(r.handlers, &Handler{
		method: "POST",
		url: url,
		handle: handler,
	})
}

func defaultHandleGet(req *Request) *Response {
	return &Response{
		status: statusNotFound,
		contentType: "text/plain; charset=UTF-8",
		content: "get",
	}
}

func defaultHandlePost(req *Request) *Response {
	return &Response{
		status: statusNotFound,
		contentType: "text/plain; charset=UTF-8",
		content: "post",
	}
}