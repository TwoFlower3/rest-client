package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

// Client ...
type Client struct {
	*http.Client
	Host string
}

// NewClient ...
func NewClient(host string) *Client {
	return &Client{
		Client: &http.Client{},
		Host:   host,
	}
}

// Do http request
func (client *Client) do(request *http.Request) (*http.Response, error) {
	response, err := client.Client.Do(request)
	if err != nil {
		return nil, fmt.Errorf("request execution error: %+v", err)
	}

	return response, nil
}

func (client *Client) unmarshal(body io.ReadCloser, model interface{}) error {

	byteBody, err := ioutil.ReadAll(body)
	if err != nil {
		return fmt.Errorf("compress body to bytes error: %+v", err)
	}

	if len(byteBody) == 0 {
		return nil
	}

	err = json.Unmarshal(byteBody, &model)
	if err != nil {
		return fmt.Errorf("response to model error: %+v", err)
	}

	return nil
}

func setQuerys(request *http.Request, querys map[string]string) {
	for key, value := range querys {
		q := request.URL.Query()
		q.Add(key, value)
		request.URL.RawQuery = q.Encode()
	}
}

func setHeaders(request *http.Request, headers map[string]string) {
	for key, value := range headers {
		request.Header.Add(key, value)
	}
}

// GET http get method
func (client *Client) GET(path string, body interface{}, headers, querys map[string]string) (int, error) {

	ctx, _ := context.WithTimeout(context.Background(), 120*time.Second)

	url := fmt.Sprintf("%s/%s", client.Host, path)

	var b []byte
	bodyReader := bytes.NewReader(b)

	request, err := http.NewRequestWithContext(ctx, "GET", url, bodyReader)
	if err != nil {
		return 0, fmt.Errorf("create request error: %+v", err)
	}

	setQuerys(request, querys)
	setHeaders(request, headers)

	response, err := client.do(request)
	if err != nil {
		return 0, fmt.Errorf("do error: %+v", err)
	}
	defer response.Body.Close()

	err = client.unmarshal(response.Body, &body)
	if err != nil {
		fmt.Printf("unmarshal error: %+v", err)
	}

	return response.StatusCode, nil
}

// POST http post method
func (client *Client) POST(path string, body interface{}, headers, querys map[string]string) (int, error) {

	ctx, _ := context.WithTimeout(context.Background(), 120*time.Second)

	url := fmt.Sprintf("%s/%s", client.Host, path)

	var b []byte
	bodyWriter := bytes.NewBuffer(b)

	err := json.NewEncoder(bodyWriter).Encode(body)
	if err != nil {
		return 0, err
	}

	request, err := http.NewRequestWithContext(ctx, "POST", url, bodyWriter)
	if err != nil {
		return 0, fmt.Errorf("create request error: %+v", err)
	}

	setQuerys(request, querys)
	setHeaders(request, headers)

	response, err := client.do(request)
	if err != nil {
		return 0, fmt.Errorf("do error: %+v", err)
	}
	defer response.Body.Close()

	err = client.unmarshal(response.Body, &body)
	if err != nil {
		fmt.Printf("unmarshal error: %+v", err)
	}

	return response.StatusCode, nil
}

// PUT http put method
func (client *Client) PUT(path string, body interface{}, headers, querys map[string]string) (int, error) {

	ctx, _ := context.WithTimeout(context.Background(), 120*time.Second)

	url := fmt.Sprintf("%s/%s", client.Host, path)

	var b []byte
	bodyWriter := bytes.NewBuffer(b)

	err := json.NewEncoder(bodyWriter).Encode(body)
	if err != nil {
		return 0, err
	}

	request, err := http.NewRequestWithContext(ctx, "PUT", url, bodyWriter)
	if err != nil {
		return 0, fmt.Errorf("create request error: %+v", err)
	}

	setQuerys(request, querys)
	setHeaders(request, headers)

	response, err := client.do(request)
	if err != nil {
		return 0, fmt.Errorf("do error: %+v", err)
	}
	defer response.Body.Close()

	err = client.unmarshal(response.Body, &body)
	if err != nil {
		fmt.Printf("unmarshal error: %+v", err)
	}

	return response.StatusCode, nil
}

// DELETE http delete method
func (client *Client) DELETE(path string, body interface{}, headers, querys map[string]string) (int, error) {

	ctx, _ := context.WithTimeout(context.Background(), 120*time.Second)

	url := fmt.Sprintf("%s/%s", client.Host, path)

	var b []byte
	bodyReader := bytes.NewReader(b)

	request, err := http.NewRequestWithContext(ctx, "DELETE", url, bodyReader)
	if err != nil {
		return 0, fmt.Errorf("create request error: %+v", err)
	}

	setQuerys(request, querys)
	setHeaders(request, headers)

	response, err := client.do(request)
	if err != nil {
		return 0, fmt.Errorf("do error: %+v", err)
	}
	defer response.Body.Close()

	err = client.unmarshal(response.Body, &body)
	if err != nil {
		fmt.Printf("unmarshal error: %+v", err)
	}

	return response.StatusCode, nil
}
