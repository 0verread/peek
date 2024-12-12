package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
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
	httpReq, err := http.NewRequest(req.Method, req.Url, strings.NewReader(req.Body))

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

func Do(url string, args ...string) {
	var method string = "GET" // default it makes GET request
	if len(args) > 0 {
		if args[0] == "POST" {
			fmt.Println(args[1])
			var jsonData map[string]interface{}
			_ = json.Unmarshal([]byte(args[1]), &jsonData)
			fmt.Println(jsonData)
			if args[1] == "" {
				fmt.Println(args[1])
			}
			// method = "POST"
		}
	}
	respBody, err := NewHttpClient().makeRequest(Request{Url: url, Method: method, Body: ""})
	if err != nil {
		fmt.Println("Error in processing request")
		fmt.Println(err)
	}
	fmt.Printf("Time Taken: %d ms\n", respBody.Latency)
	fmt.Println(respBody.Body)
}
