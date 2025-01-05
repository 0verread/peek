package client

import (
	"net/http"
	"reflect"
	"testing"
)

func TestParseHeader(t *testing.T) {
	tests := []struct {
		name      string
		headerStr string
		want      http.Header
	}{
		{
			name:      "empty header",
			headerStr: "",
			want:      http.Header{},
		},
		{
			name:      "single header",
			headerStr: "Content-Type: application/json",
			want:      http.Header{"Content-Type": []string{"application/json"}},
		},
		{
			name:      "multiple headers",
			headerStr: "Content-Type: application/json; Authorization: Bearer token",
			want: http.Header{
				"Content-Type":  []string{"application/json"},
				"Authorization": []string{"Bearer token"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := parseHeader(tt.headerStr)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseHeader() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBuildPostReq(t *testing.T) {
	tests := []struct {
		name    string
		url     string
		headers http.Header
		payload []byte
		want    Request
	}{
		{
			name:    "default headers",
			url:     "https://jsonplaceholder.typicode.com/todos/1",
			headers: nil,
			payload: []byte(`{"test": "data"}`),
			want: Request{
				Url:     "https://jsonplaceholder.typicode.com/todos/1",
				Method:  Post,
				Headers: http.Header{"Content-Type": []string{"application/json"}},
				Body:    []byte(`{"test": "data"}`),
			},
		},
		{
			name:    "custom headers",
			url:     "https://jsonplaceholder.typicode.com/todos/1",
			headers: http.Header{"Custom-Header": []string{"value"}},
			payload: []byte(`{"test": "data"}`),
			want: Request{
				Url:     "https://jsonplaceholder.typicode.com/todos/1",
				Method:  Post,
				Headers: http.Header{"Custom-Header": []string{"value"}, "Content-Type": []string{"application/json"}},
				Body:    []byte(`{"test": "data"}`),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := buildPostReq(tt.url, tt.headers, tt.payload)
			if err != nil {
				t.Errorf("buildPostReq() error = %v", err)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("buildPostReq() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDo(t *testing.T) {
	tests := []struct {
		name    string
		url     string
		args    []string
		wantErr bool
	}{
		{
			name:    "GET request",
			url:     "https://jsonplaceholder.typicode.com/todos/1",
			args:    []string{"GET"},
			wantErr: false,
		},
		{
			name:    "POST request",
			url:     "https://jsonplaceholder.typicode.com/todos/1",
			args:    []string{"POST", `{"key":"value"}`, "Content-Type: application/json"},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := Do(tt.url, tt.args...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Do() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
