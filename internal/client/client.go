package client

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

const (
	Get    = "GET"
	Post   = "POST"
	Delete = "DELETE"
	Put    = "PUT"
)

type HttpClient interface {
	makeRequest(req Request) (Response, error)
}

type DefaultClient struct {
	client *http.Client
}

func NewHttpClient() DefaultClient {
	return DefaultClient{client: &http.Client{}}
}

func (dc DefaultClient) makeRequest(req Request) (Response, error) {
	httpReq, err := http.NewRequest(req.Method, req.Url, bytes.NewBuffer(req.Body))

	if err != nil {
		fmt.Println("Something wrong with building http new request")
		return Response{}, fmt.Errorf("error")
	}

	startTime := time.Now()
	res, err := dc.client.Do(httpReq)
	timeTaken := time.Since(startTime).Milliseconds()
	if err != nil {
		fmt.Println("Somehting went wrong to call API")
		return Response{}, err
	}
	defer res.Body.Close()
	respBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return Response{}, fmt.Errorf("Error in gettting response body")
	}

	return Response{
		Latency: timeTaken,
		Body:    string(respBody),
		Status:  res.StatusCode,
		Headers: res.Header,
	}, nil
}

func buildPostReq(url string, headers http.Header, payload []byte) (Request, error) {
	if headers == nil {
		headers = make(http.Header)
	}

	if headers.Get("Content-Type") == "" {
		headers.Add("Content-Type", "application/json")
	}

	return Request{
		Url:     url,
		Method:  Post,
		Headers: headers,
		Body:    payload,
	}, nil
}

func parseHeader(headerStr string) http.Header {
	headers := make(http.Header)
	if headerStr == "" {
		return headers
	}

	// Split headers by semicolon
	headerPairs := strings.Split(headerStr, ";")
	for _, pair := range headerPairs {
		parts := strings.Split(pair, ":")
		if len(parts) == 2 {
			headers.Add(strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1]))
		}
	}
	return headers
}

func Do(url string, args ...string) (Response, error) {
	var verb string = Get // default it makes GET request
	var payload []byte
	var headers http.Header
	var req Request
	var err error

	if len(args) > 0 {
		verb = args[0]
	}

	if len(args) > 0 && (strings.EqualFold(verb, Post) || strings.EqualFold(verb, Put)) {
		verb = Post
		if len(args) > 1 {
			payload = []byte(args[1])
		}
		if len(args) > 2 {
			headers = parseHeader(args[2])
		}
		req, err = buildPostReq(url, headers, payload)
	} else {
		req, err = Request{Url: url, Method: verb}, nil
	}

	if err != nil {
		fmt.Println("Error building request for POST, error:", err)
		panic(0)
	}
	response, err := NewHttpClient().makeRequest(req)
	if err != nil {
		fmt.Println("Error in processing request, error: ", err)
	}
	return response, err
	// fmt.Printf("Status: %d  Time Taken: %d ms\n", respBody.Status, respBody.Latency)
	// cout.PrettyPrint([]byte(respBody.Body))
}
