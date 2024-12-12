package client

import (
	"net/http"
)

type Response struct {
	Headers http.Header
	Latency int64
	Body    string
	Status  int
}

type ResponseError struct {
	Request Request
	Err     error
}

type Request struct {
	Url     string
	Method  string
	Headers http.Header
	Body    string
}

type RequestError struct {
	Err     error
	Request *Request
}
