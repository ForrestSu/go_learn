package main

type Request struct {
	name string
}

func NewRequest() *Request {
	return &Request{
		name: "hello req,",
	}
}

type Response struct {
	name string
}

func NewResponse() *Response {
	return &Response{
		name: "world rsp.",
	}
}
