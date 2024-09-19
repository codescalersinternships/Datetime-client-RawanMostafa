package pkg

import (
	"fmt"
	"io"
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

func (c Client) ReadBody(resp *http.Response) (string, error) {
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error in reading request body: %v", err)

	}
	return string(body), nil
}

func (c Client) SendRequest() (*http.Response, error) {
	resp, err := c.client.Get(c.baseUrl + ":" + c.port + c.endpoint)
	if err != nil {
		return nil, fmt.Errorf("error in sending request: %v", err)
	}
	return resp, nil
}
