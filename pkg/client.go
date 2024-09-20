package pkg

import (
	"fmt"
	"net/http"
	"time"

	"github.com/cenkalti/backoff/v4"
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

func (c Client) sendRequest(contentType string) (*http.Response, error) {
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

func (c Client) RetrySendRequest(contentType string) (*http.Response, error) {
	var resp *http.Response
	var err error

	expBackoff := backoff.NewExponentialBackOff()
	expBackoff.MaxElapsedTime = 10 * time.Second
	
	retryError := backoff.RetryNotify(func() error {
		resp, err = c.sendRequest(contentType)
		return err
	}, expBackoff, func(err error, d time.Duration) {
		fmt.Println("Retry Happenned")
	})

	if retryError != nil {
		return resp, fmt.Errorf("failed to make the request after retries: %v", err)
	} else {
		return resp, nil
	}
}
