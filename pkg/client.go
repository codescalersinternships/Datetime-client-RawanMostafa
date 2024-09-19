package pkg

import (
	"fmt"
	"net/http"
	"time"
)

type Client struct {
	baseUrl  string
	endpoint string
	port     string
	client   http.Client
}

func NewClient(baseUrl string, endpoint string, port string, timeout time.Duration) Client {
	return Client{
		baseUrl:  baseUrl,
		endpoint: endpoint,
		port:     port,
		client:   http.Client{Timeout: timeout},
	}
}

func (c Client) SendRequest(contentType string) (*http.Response, error) {
	req, err := http.NewRequest("GET", c.baseUrl+":"+c.port+c.endpoint, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("content-type", contentType)
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error in sending request: %v", err)
	}
	return resp, nil
}
