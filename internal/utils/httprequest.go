package utils

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	fetchTimeout       = 1 * time.Second
	httpRequestTimeout = 10 * time.Second
)

var client = http.Client{
	Timeout: httpRequestTimeout,
}

type HttpRequest interface {
	MakePostRequest(rURL string, reqH map[string]string, body url.Values) ([]byte, error)
	MakePostJSONRequest(rURL string, reqH map[string]string, body string) ([]byte, error)
	MakeGetRequest(rURL string, reqH map[string]string) ([]byte, error)
}

type httpRequest struct {
}

func NewHttpRequest() HttpRequest {
	return &httpRequest{}
}

func (c httpRequest) MakePostRequest(rURL string, reqH map[string]string, body url.Values) ([]byte, error) {

	req, err := http.NewRequest("POST", rURL, strings.NewReader(body.Encode()))
	if err != nil {
		return nil, err
	}

	for kH, vH := range reqH {
		req.Header.Set(kH, vH)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func (c httpRequest) MakePostJSONRequest(rURL string, reqH map[string]string, body string) ([]byte, error) {

	var JSONStr = []byte(body)
	var bodyBuffer = bytes.NewBuffer(JSONStr)

	req, err := http.NewRequest("POST", rURL, bodyBuffer)
	if err != nil {
		return nil, err
	}

	for kH, vH := range reqH {
		req.Header.Set(kH, vH)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	fmt.Println(resp)

	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func (c httpRequest) MakeGetRequest(rURL string, reqH map[string]string) ([]byte, error) {

	req, err := http.NewRequest("GET", rURL, nil)
	if err != nil {
		return nil, err
	}

	for kH, vH := range reqH {
		req.Header.Set(kH, vH)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}
